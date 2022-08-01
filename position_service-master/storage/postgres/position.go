package postgres

import (
	"context"
	"fmt"
	pb "position_service/genproto/position_service"
	"position_service/pkg/helper"
	"position_service/storage"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type positionRepo struct {
	db *pgxpool.Pool
}

func NewPositionRepo(db *pgxpool.Pool) storage.PositionI {
	return &positionRepo{
		db: db,
	}
}

func (r *positionRepo) Create(ctx context.Context, entity *pb.CreatePositionRequest) (*pb.Position, error) {
	id := uuid.NewString()
	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})

	query := `
		INSERT INTO position (
			id,
			name,
			profession_id,
			company_id
		)
		 VALUES ($1, $2, $3, $4)
	`
	query2 := `INSERT INTO position_attributes (id, attribute_id, position_id, value) VALUES ($1,$2,$3,$4)`

	_, err = tx.Exec(ctx, query, id, entity.Name, entity.ProfessionId, entity.CompanyId)
	if err != nil {
		tx.Rollback(ctx)
		return nil, fmt.Errorf("error while inserting position err: %w", err)
	}

	for i := 0; i < len(entity.PositionAttributes); i++ {
		id2 := uuid.NewString()
		_, err := tx.Exec(
			ctx,
			query2,
			id2,
			entity.PositionAttributes[i].AttributeId,
			id,
			entity.PositionAttributes[i].Value,
		)
		if err != nil {
			tx.Rollback(ctx)
			return nil, fmt.Errorf("error while inserting position err: %w", err)
		}

	}
	err = tx.Commit(ctx)
	if err != nil {
		return nil, fmt.Errorf("error while inserting position err: %w", err)
	}

	return &pb.Position{
		Id:                 id,
		Name:               entity.Name,
		ProfessionId:       entity.ProfessionId,
		CompanyId:          entity.CompanyId,
		PositionAttributes: []*pb.GetPositionAttributes{},
	}, nil
}

func (r *positionRepo) GetAll(ctx context.Context, req *pb.GetAllPositionRequest) (*pb.GetAllPositionResponse, error) {
	var (
		resp   pb.GetAllPositionResponse
		err    error
		filter string
		params = make(map[string]interface{})
	)

	if req.Search != "" {
		filter += " AND p.name ILIKE '%' || :search || '%' "
		params["search"] = req.Search
	}

	countQuery := `SELECT count(1) FROM position as p WHERE true ` + filter

	q, arr := helper.ReplaceQueryParams(countQuery, params)
	err = r.db.QueryRow(ctx, q, arr...).Scan(
		&resp.Count,
	)

	if err != nil {
		return nil, fmt.Errorf("error while scanning count %w", err)
	}

	query := `SELECT p.*,
	jsonb_agg(jsonb_build_object(
		'id', pa.id,
		'attribute_id', pa.attribute_id,
		'position_id', pa.position_id,
		'value', pa.value,
		'attribute_name', a.name,
		'attribute_type', a.type
	)) FROM position as p
	LEFT JOIN position_attributes as pa
		on pa.position_id=p.id
	LEFT JOIN attribute as a
	on pa.attribute_id=a.id
	WHERE true ` + filter + `
	GROUP BY p.id, pa.position_id`

	query += " LIMIT :limit OFFSET :offset"
	params["limit"] = req.Limit
	params["offset"] = req.Offset

	q, arr = helper.ReplaceQueryParams(query, params)
	rows, err := r.db.Query(ctx, q, arr...)
	if err != nil {
		return nil, fmt.Errorf("error while getting rows %w", err)
	}

	for rows.Next() {
		var position pb.Position
		err = rows.Scan(
			&position.Id,
			&position.Name,
			&position.ProfessionId,
			&position.CompanyId,
			&position.PositionAttributes,
		)
		if err != nil {
			return nil, fmt.Errorf("error while scanning position err: %w", err)
		}
		resp.Positions = append(resp.Positions, &position)
	}

	return &resp, nil
}

func (r *positionRepo) Update(ctx context.Context, entity *pb.UpdatePositionRequest) (*pb.PositionId, error) {

	query := `UPDATE position SET name=$1, profession_id=$2, company_id=$3 WHERE id=$4`
	query2 := `UPDATE position_attributes SET attribute_id=$1, value=$2 where id=$3`
	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})

	rows, err := tx.Query(ctx, `select id from position_attributes where position_id=$1`, entity.Id)
	if err != nil {
		tx.Rollback(ctx)
		return nil, fmt.Errorf("error while scanning position_attributes id err: %w", err)
	}
	arrayIds := make([]string, 0)
	for rows.Next() {
		var id string
		err := rows.Scan(
			&id,
		)
		if err != nil {
			tx.Rollback(ctx)
			return nil, fmt.Errorf("error while scanning position_attributes id err: %w", err)
		}
		arrayIds = append(arrayIds, id)
	}

	for i := 0; i < len(arrayIds); i++ {
		var flag bool = true
		for _, n := range entity.PositionAttributesWithId {
			if n.Id == arrayIds[i] {
				flag = false
			}
		}
		if flag {
			_, err = tx.Exec(ctx, `delete from position_attributes where id=$1`, arrayIds[i])
			if err != nil {
				tx.Rollback(ctx)
				return nil, fmt.Errorf("error while scanning position_attributes id err: %w", err)
			}

		}

	}

	_, err = tx.Exec(ctx, query, entity.Name, entity.ProfessionId, entity.CompanyId, entity.Id)

	if err != nil {
		tx.Rollback(ctx)
		return nil, fmt.Errorf("error while updating position err: %w", err)
	}
	for i := 0; i < len(entity.PositionAttributesWithId); i++ {
		if entity.PositionAttributesWithId[i].Id != "" {
			_, err := tx.Exec(ctx, query2,
				entity.PositionAttributesWithId[i].AttributeId,
				entity.PositionAttributesWithId[i].Value,
				entity.PositionAttributesWithId[i].Id)
			if err != nil {
				tx.Rollback(ctx)
				return nil, fmt.Errorf("error while updating position_attrbiutes err: %w", err)
			}
		} else {
			newId := uuid.NewString()
			_, err = tx.Exec(ctx, `INSERT INTO position_attributes (id, attribute_id, position_id, value) VALUES ($1,$2,$3,$4)`,
				newId,
				entity.PositionAttributesWithId[i].AttributeId,
				entity.Id,
				entity.PositionAttributesWithId[i].Value,
			)
			if err != nil {
				tx.Rollback(ctx)
				return nil, fmt.Errorf("error while updating position_attrbiutes err: %w", err)
			}
		}

	}
	err = tx.Commit(ctx)
	if err != nil {
		return nil, fmt.Errorf("error while committing err: %w", err)
	}

	return &pb.PositionId{
		Id: entity.Id,
	}, nil
}

func (r *positionRepo) Delete(ctx context.Context, entity *pb.PositionId) (*pb.PositionId, error) {
	givenId := entity.Id
	query1 := `DELETE FROM position_attributes WHERE position_id=$1`
	query2 := `DELETE FROM position WHERE id=$1`
	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})

	_, err = tx.Exec(ctx, query1, givenId)
	if err != nil {
		tx.Rollback(ctx)
		return nil, fmt.Errorf("error while deleting position_attributes err: %w", err)
	}

	_, err = tx.Exec(
		ctx,
		query2,
		givenId,
	)
	if err != nil {
		tx.Rollback(ctx)
		return nil, fmt.Errorf("error while deleting position err: %w", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, fmt.Errorf("error while committing err: %w", err)
	}

	return &pb.PositionId{
		Id: givenId,
	}, nil
}

func (r *positionRepo) GetById(ctx context.Context, entity *pb.PositionId) (*pb.Position, error) {
	var resp pb.Position

	query := `SELECT p.*,
	jsonb_agg(jsonb_build_object(
		'id', pa.id,
		'attribute_id', pa.attribute_id,
		'position_id', pa.position_id,
		'value', pa.value,
		'attribute_name', a.name,
		'attribute_type', a.type
	)) FROM position as p
	LEFT JOIN position_attributes as pa
		on pa.position_id=p.id
	LEFT JOIN attribute as a
	on pa.attribute_id=a.id
	WHERE p.id=$1
	GROUP BY p.id, pa.position_id`

	row, err := r.db.Query(
		ctx,
		query,
		entity.Id,
	)
	if err != nil {
		return nil, fmt.Errorf("error while getting position by given id err: %w", err)
	}

	for row.Next() {
		err = row.Scan(
			&resp.Id,
			&resp.Name,
			&resp.ProfessionId,
			&resp.CompanyId,
			&resp.PositionAttributes,
		)
	}

	return &resp, nil
}
