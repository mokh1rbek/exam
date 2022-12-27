package models

type OrderPrimarKey struct {
	Id string `json:"order_id"`
}

type CreateOrder struct {
	Description string `json:"description"`
	ProductId   string `json:"product_id"`
}

type Order struct {
	Id          string `json:"order_id"`
	Description string `json:"description"`
	ProductId   string `json:"product_id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type UpdateOrderSwagger struct {
	Description string `json:"description"`
	ProductId   string `json:"product_id"`
}

type UpdateOrder struct {
	Id          string `json:"order_id"`
	Description string `json:"description"`
	ProductId   string `json:"product_id"`
}

type GetListOrderRequest struct {
	Limit  int32
	Offset int32
}

type GetListOrderResponse struct {
	Count  int          `json:"count"`
	Orders []*OrderList `json:"orders"`
}

type OrderList struct {
	Id          string         `json:"order_id"`
	Description string         `json:"description"`
	Product     []*ProductList `json:"product"`
}
type ProductList struct {
	Id       string             `json:"order_id"`
	Name     string             `json:"name"`
	Category []*ProductCategory `json:"category"`
}
type ProductCategory struct {
	Id       string `json:"order_id"`
	Name     string `json:"name"`
	ParentID string `json:"parent_id"`
}
