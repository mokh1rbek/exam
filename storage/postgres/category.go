package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"

	"exam/models"
	"exam/pkg/helper"
)

type CategoryRepo struct {
	db *pgxpool.Pool
}

func NewCategoryRepo(db *pgxpool.Pool) *CategoryRepo {
	return &CategoryRepo{
		db: db,
	}
}

func (f *CategoryRepo) Create(ctx context.Context, category *models.CreateCategory) (string, error) {

	var (
		id       = uuid.New().String()
		query    string
		nulluuid sql.NullString
	)

	query = `
		INSERT INTO category(
			category_id,
			name,
			parent_uuid,
			updated_at
		) VALUES ( $1, $2, $3, now())
	`

	if category.ParentUUID == "" {
		_, err := f.db.Exec(ctx, query,
			id,
			category.Name,
			nulluuid,
		)

		if err != nil {
			return "", err
		}
	} else {

		_, err := f.db.Exec(ctx, query,
			id,
			category.Name,
			category.ParentUUID,
		)

		if err != nil {
			return "", err
		}
	}

	return id, nil
}

func (f *CategoryRepo) GetByPKey(ctx context.Context, pkey *models.CategoryPrimarKey) (*models.Category, error) {

	var (
		id          sql.NullString
		name        sql.NullString
		parent_uuid sql.NullString
		createdAt   sql.NullString
		updatedAt   sql.NullString
	)

	query := `
		SELECT
			category_id,
			name,
			parent_uuid,
			created_at,
			updated_at
		FROM
			category
		WHERE category_id = $1
	`

	err := f.db.QueryRow(ctx, query, pkey.Id).
		Scan(
			&id,
			&name,
			&parent_uuid,
			&createdAt,
			&updatedAt,
		)

	if err != nil {
		return nil, err
	}

	return &models.Category{
		Id:         id.String,
		Name:       name.String,
		ParentUUID: parent_uuid.String,
		CreatedAt:  createdAt.String,
		UpdatedAt:  updatedAt.String,
	}, nil
}

func (f *CategoryRepo) GetList(ctx context.Context, req *models.GetListCategoryRequest) (*models.GetListCategoryResponse, error) {

	var (
		resp   = models.GetListCategoryResponse{}
		offset = " OFFSET 0"
		limit  = " LIMIT 5"
	)

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	query := `
		SELECT
			COUNT(*) OVER(),
			c1.category_id,
			c1.name,
			c1.parent_uuid
		FROM
			category AS c1
		WHERE c1.parent_uuid IS NULL
	`

	query += offset + limit

	rows, err := f.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {

		res := &models.CategoryByParent{}

		var (
			id          sql.NullString
			name        sql.NullString
			parent_uuid sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&name,
			&parent_uuid,
		)

		if err != nil {
			return nil, err
		}

		res.Id = id.String
		res.Name = name.String
		res.ParentUUID = parent_uuid.String

		resp.Categorys = append(resp.Categorys, res)
	}

	for ind, category := range resp.Categorys {

		childQuery := `
			SELECT
				c1.category_id,
				c1.name,
				c1.parent_uuid
			FROM
				category AS c1
			WHERE c1.parent_uuid =  $1
		`

		childRows, err := f.db.Query(ctx, childQuery, category.Id)
		if err != nil {
			return nil, err
		}

		for childRows.Next() {

			var (
				id          sql.NullString
				name        sql.NullString
				parent_uuid sql.NullString
			)

			err := childRows.Scan(
				&id,
				&name,
				&parent_uuid,
			)

			if err != nil {
				return nil, err
			}

			category.ChildCategory = append(category.ChildCategory, &models.ChildCategory{
				Id:       id.String,
				Name:     name.String,
				ParentId: parent_uuid.String,
			})

		}

		resp.Categorys[ind] = category
		fmt.Println(category)
	}

	return &resp, err
}

func (f *CategoryRepo) Update(ctx context.Context, id string, req *models.UpdateCategory) (int64, error) {

	var (
		query  = ""
		params map[string]interface{}
	)

	query = `
		UPDATE
			category
		SET
			name = :name,
			parent_uuid = :parent_uuid,
			updated_at = now()
		WHERE category_id = :category_id
	`

	params = map[string]interface{}{
		"name":        req.Name,
		"parent_uuid": req.ParentUUID,
		"category_id": id,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	rowsAffected, err := f.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return rowsAffected.RowsAffected(), nil
}

func (f *CategoryRepo) Delete(ctx context.Context, req *models.CategoryPrimarKey) error {

	_, err := f.db.Exec(ctx, "DELETE FROM category WHERE category_id = $1", req.Id)
	if err != nil {
		return err
	}

	return err
}