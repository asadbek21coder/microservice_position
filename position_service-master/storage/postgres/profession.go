package postgres

import (
	"context"
	"fmt"
	"position_service/pkg/helper"
	"position_service/storage"

	pb "position_service/genproto/position_service"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type professionRepo struct {
	db *pgxpool.Pool
}

func NewProfessionRepo(db *pgxpool.Pool) storage.ProfessionI {
	return &professionRepo{
		db: db,
	}
}

func (r *professionRepo) Create(ctx context.Context, entity *pb.CreateProfessionRequest) (id string, err error) {
	query := `
		INSERT INTO profession (
			id,
			name
		) 
		 VALUES ($1, $2)
	`

	id = uuid.NewString()

	_, err = r.db.Exec(
		ctx,
		query,
		id,
		entity.Name,
	)

	if err != nil {
		return "", fmt.Errorf("error while inserting profession err: %w", err)
	}

	return id, nil
}

func (r *professionRepo) Update(ctx context.Context, entity *pb.Profession) (id string, err error) {

	givenId := entity.Id
	query := `UPDATE profession SET name=$1 WHERE id=$2`

	_, err = r.db.Query(ctx, query, entity.Name, givenId)

	if err != nil {
		return "", fmt.Errorf("error while updating profession err: %w", err)
	}

	return givenId, nil
}

func (r *professionRepo) GetAll(ctx context.Context, req *pb.GetAllProfessionRequest) (*pb.GetAllProfessionResponse, error) {
	var (
		resp   pb.GetAllProfessionResponse
		err    error
		filter string
		params = make(map[string]interface{})
	)

	if req.Search != "" {
		filter += " AND name ILIKE '%' || :search || '%' "
		params["search"] = req.Search
	}

	countQuery := `SELECT count(1) FROM profession WHERE true ` + filter

	q, arr := helper.ReplaceQueryParams(countQuery, params)
	err = r.db.QueryRow(ctx, q, arr...).Scan(
		&resp.Count,
	)

	if err != nil {
		return nil, fmt.Errorf("error while scanning count %w", err)
	}

	query := `SELECT
				id,
				name
			FROM profession
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
		var profession pb.Profession

		err = rows.Scan(
			&profession.Id,
			&profession.Name,
		)

		if err != nil {
			return nil, fmt.Errorf("error while scanning profession err: %w", err)
		}

		resp.Professions = append(resp.Professions, &profession)
	}

	return &resp, nil
}

func (r *professionRepo) Delete(ctx context.Context, entity *pb.Id) (id string, err error) {
	givenId := entity.Id
	query := `DELETE FROM profession WHERE id=$1`

	_, err = r.db.Exec(
		ctx,
		query,
		givenId,
	)

	if err != nil {
		return "", fmt.Errorf("error while deleting profession err: %w", err)
	}

	return givenId, nil
}
func (r *professionRepo) GetById(ctx context.Context, entity *pb.Id) (*pb.Profession, error) {
	var resp pb.Profession

	query := `SELECT * FROM profession WHERE id=$1`

	row, err := r.db.Query(
		ctx,
		query,
		entity.Id,
	)
	if err != nil {
		return nil, fmt.Errorf("error while getting profession by given id err: %w", err)
	}

	for row.Next() {
		err = row.Scan(
			&resp.Id,
			&resp.Name,
		)
		if err != nil {
			return nil, fmt.Errorf("error while scanning Profession err: %w", err)
		}

	}

	return &resp, nil
}
