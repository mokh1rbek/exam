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
		offset = ""
		limit  = ""
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
			o.id,
			o.description,
			p.id,
			p.name,
			c.id,
			c.name,
			c.parent_id
		FROM
			order AS o
		JOIN product AS p ON o.product_id = p.id
		JOIN category ON p.category_id = c.id
	`

	query += offset + limit

	rows, err := f.db.Query(ctx, query)

	for rows.Next() {
		
		var (
			productCategory models.ProductCategory
			productList     models.ProductList

			orderId          sql.NullString
			orderDescription sql.NullString
			productId        sql.NullString
			productName      sql.NullString
			categoryId       sql.NullString
			categoryName     sql.NullString
			categoryParentId sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&orderId,
			&orderDescription,
			&productId,
			&productName,
			&categoryId,
			&categoryName,
			&categoryParentId,
		)

		if err != nil {
			return nil, err
		}

		productCategory.Id = categoryId.String
		productCategory.Name = categoryName.String
		productCategory.ParentID = categoryParentId.String

		productList.Id = productId.String
		productList.Name = productName.String
		productList.Category = productCategory

		resp.Orders = append(resp.Orders, models.OrderList{
			Id:          orderId.String,
			Description: orderDescription.String,
			Product:     productList,
		})

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
		"order_id":    id,
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
