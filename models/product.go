package models

type ProductPrimarKey struct {
	Id string `json:"id"`
}

type CreateProduct struct {
	Name       string `json:"name"`
	Price      string `json:"price"`
	CategoryId string `json:"category_id"`
}

type UpdateProduct struct {
	Id         string `json:"product_id"`
	Name       string `json:"name"`
	Price      string `json:"price"`
	CategoryId string `json:"category_id"`
}

type Product struct {
	Id         string `json:"product_id"`
	Name       string `json:"name"`
	Price      string `json:"price"`
	CategoryId string `json:"category_id"`
	CreatedAt  string `json:"created_at"`
	ApdatedAt  string `json:"updated_at"`
}

type GetListProductRequest struct {
	Limit  int32
	Offset int32
}

type GetListProductResponse struct {
	Count    int32      `json:"count"`
	Products []*Product `json:"product"`
}
