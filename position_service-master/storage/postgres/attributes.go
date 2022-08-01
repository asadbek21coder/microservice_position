package postgres

import (
	"context"
	"fmt"
	pb "position_service/genproto/position_service"
	"position_service/pkg/helper"
	"position_service/storage"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type attributeRepo struct {
	db *pgxpool.Pool
}

func NewAttributeRepo(db *pgxpool.Pool) storage.AttributeI {
	return &attributeRepo{
		db: db,
	}
}

func (r *attributeRepo) CreateAttribute(ctx context.Context, entity *pb.CreateAttributeRequest) (id string, err error) {
	query := `
		INSERT INTO attribute (
			id,
			name,
			type
		)
		 VALUES ($1, $2,$3)
	`

	id = uuid.NewString()

	_, err = r.db.Exec(
		ctx,
		query,
		id,
		entity.Name,
		entity.AttributeType,
	)

	if err != nil {
		return "", fmt.Errorf("error while inserting attribute err: %w", err)
	}

	return id, nil
}

func (r *attributeRepo) UpdateAttribute(ctx context.Context, entity *pb.Attribute) (id string, err error) {

	givenId := entity.Id
	query := `UPDATE attribute SET name=$1, type=$2 WHERE id=$3`

	_, err = r.db.Query(ctx, query, entity.Name, entity.AttributeType, givenId)

	if err != nil {
		return "", fmt.Errorf("error while updating attribute err: %w", err)
	}

	return givenId, nil
}

func (r *attributeRepo) GetAllAttribute(ctx context.Context, req *pb.GetAllAttributeRequest) (*pb.GetAllAttributeResponse, error) {
	var (
		resp   pb.GetAllAttributeResponse
		err    error
		filter string
		params = make(map[string]interface{})
	)

	if req.Search != "" {
		filter += " AND name ILIKE '%' || :search || '%' "
		params["search"] = req.Search
	}

	countQuery := `SELECT count(1) FROM attribute WHERE true ` + filter

	q, arr := helper.ReplaceQueryParams(countQuery, params)
	err = r.db.QueryRow(ctx, q, arr...).Scan(
		&resp.Count,
	)

	if err != nil {
		return nil, fmt.Errorf("error while scanning count %w", err)
	}

	query := `SELECT
				id,
				name, 
				type
			FROM attribute
			WHERE true` + filter

	query += " LIMIT :limit OFFSET :offset"
	params["limit"] = req.Limit
	params["offset"] = req.Offset

	q, arr = helper.ReplaceQueryParams(query, params)
	rows, err := r.db.Query(ctx, q, arr...)
	if err != nil {
		return nil, fmt.Errorf("error while getting rows %w", err)
	}

	for rows.Next() {
		var attribute pb.Attribute

		err = rows.Scan(
			&attribute.Id,
			&attribute.Name,
			&attribute.AttributeType,
		)

		if err != nil {
			return nil, fmt.Errorf("error while scanning attribute err: %w", err)
		}

		resp.Attributes = append(resp.Attributes, &attribute)
	}

	return &resp, nil
}

func (r *attributeRepo) DeleteAttribute(ctx context.Context, entity *pb.AttributeId) (id string, err error) {
	givenId := entity.Id
	query := `DELETE FROM attribute WHERE id=$1`

	_, err = r.db.Exec(
		ctx,
		query,
		givenId,
	)

	if err != nil {
		return "", fmt.Errorf("error while deleting attribute err: %w", err)
	}

	return givenId, nil
}
func (r *attributeRepo) GetByIdAttribute(ctx context.Context, entity *pb.AttributeId) (*pb.Attribute, error) {
	var resp pb.Attribute
	fmt.Println(entity)

	query := `SELECT * FROM attribute WHERE id=$1`

	row, err := r.db.Query(
		ctx,
		query,
		entity.Id,
	)
	if err != nil {
		return nil, fmt.Errorf("error while getting attribute by given id err: %w", err)
	}

	for row.Next() {
		err = row.Scan(
			&resp.Id,
			&resp.Name,
			&resp.AttributeType,
		)
		if err != nil {
			return nil, fmt.Errorf("error while scanning Attribute err: %w", err)
		}

	}

	return &resp, nil
}
