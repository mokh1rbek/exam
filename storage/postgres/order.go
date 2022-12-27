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

type OrderRepo struct {
	db *pgxpool.Pool
}

func NewOrderRepo(db *pgxpool.Pool) *OrderRepo {
	return &OrderRepo{
		db: db,
	}
}

func (f *OrderRepo) Create(ctx context.Context, order *models.CreateOrder) (string, error) {

	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO order(
			order_id,
			description,
			product_id,
			updated_at
		) VALUES ( $1, $2, $3, now())
	`

	_, err := f.db.Exec(ctx, query,
		id,
		order.Description,
		order.ProductId,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (f *OrderRepo) GetByPKey(ctx context.Context, pkey *models.OrderPrimarKey) (*models.Order, error) {

	var (
		id          sql.NullString
		description sql.NullString
		productId   sql.NullString
		createdAt   sql.NullString
		updatedAt   sql.NullString
	)

	query := `
		SELECT
			order_id,
			description,
			product_id,
			created_at,
			updated_at
		FROM
			order
		WHERE order_id = $1
	`

	err := f.db.QueryRow(ctx, query, pkey.Id).
		Scan(
			&id,
			&description,
			&productId,
			&createdAt,
			&updatedAt,
		)

	if err != nil {
		return nil, err
	}

	return &models.Order{
		Id:          id.String,
		Description: description.String,
		ProductId:   productId.String,
		CreatedAt:   createdAt.String,
		UpdatedAt:   updatedAt.String,
	}, nil
}

func (f *OrderRepo) GetList(ctx context.Context, req *models.GetListOrderRequest) (*models.GetListOrderResponse, error) {

	var (
		resp   = models.GetListOrderResponse{}
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
			c1.product_id
		FROM
			order AS c1
		WHERE c1.product_id IS NULL
	`

	query += offset + limit

	rows, err := f.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {

		res := &models.OrderList{}

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

		resp.Orders = append(resp.Orders, )
	}

	for ind, category := range resp.Orders {

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

func (f *OrderRepo) Update(ctx context.Context, id string, req *models.UpdateOrder) (int64, error) {

	var (
		query  = ""
		params map[string]interface{}
	)

	query = `
		UPDATE
			order
		SET
			description = :description,
			product_id = :product_id,
			updated_at = now()
		WHERE order_id = :order_id
	`

	params = map[string]interface{}{
		"description": req.Description,
		"product_id":  req.ProductId,
		"order_id": id,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	rowsAffected, err := f.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return rowsAffected.RowsAffected(), nil
}

func (f *OrderRepo) Delete(ctx context.Context, req *models.OrderPrimarKey) error {

	_, err := f.db.Exec(ctx, "DELETE FROM order WHERE order_id = $1", req.Id)
	if err != nil {
		return err
	}

	return err
}
