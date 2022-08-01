package postgres

import (
	"context"
	"fmt"
	pc "position_service/genproto/company_service"
	"position_service/pkg/helper"
	"position_service/storage"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type companyRepo struct {
	db *pgxpool.Pool
}

func NewCompanyRepo(db *pgxpool.Pool) storage.CompanyI {
	return &companyRepo{
		db: db,
	}
}

func (r *companyRepo) Create(ctx context.Context, entity *pc.CreateCompanyRequest) (id string, err error) {
	query := `
		INSERT INTO company (
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
		return "", fmt.Errorf("error while inserting company err: %w", err)
	}

	return id, nil
}

func (r *companyRepo) GetAll(ctx context.Context, req *pc.GetAllCompanyRequest) (*pc.GetAllCompanyResponse, error) {
	var (
		resp   pc.GetAllCompanyResponse
		err    error
		filter string
		params = make(map[string]interface{})
	)

	if req.Search != "" {
		filter += " AND name ILIKE '%' || :search || '%' "
		params["search"] = req.Search
	}

	countQuery := `SELECT count(1) FROM company WHERE true ` + filter

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
			FROM company
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
		var company pc.Company

		err = rows.Scan(
			&company.Id,
			&company.Name,
		)

		if err != nil {
			return nil, fmt.Errorf("error while scanning profession err: %w", err)
		}

		resp.Companys = append(resp.Companys, &company)
	}

	return &resp, nil
}

func (r *companyRepo) Update(ctx context.Context, entity *pc.Company) (id string, err error) {

	givenId := entity.Id
	query := `UPDATE company SET name=$1 WHERE id=$2`

	_, err = r.db.Query(ctx, query, entity.Name, givenId)

	if err != nil {
		return "", fmt.Errorf("error while updating company err: %w", err)
	}

	return givenId, nil
}

func (r *companyRepo) Delete(ctx context.Context, entity *pc.Id) (id string, err error) {
	givenId := entity.Id
	query := `DELETE FROM company WHERE id=$1`

	_, err = r.db.Exec(
		ctx,
		query,
		givenId,
	)

	if err != nil {
		return "", fmt.Errorf("error while deleting company err: %w", err)
	}

	return givenId, nil
}

func (r *companyRepo) GetById(ctx context.Context, entity *pc.Id) (*pc.Company, error) {
	var resp pc.Company
	query := `SELECT * FROM company WHERE id=$1`

	row, err := r.db.Query(
		ctx,
		query,
		entity.Id,
	)

	if err != nil {
		return nil, fmt.Errorf("error while selecting company by given id err: %w", err)
	}

	for row.Next() {
		err = row.Scan(
			&resp.Id,
			&resp.Name,
		)
	}

	return &resp, nil
}
