package storage

import (
	"context"

	"exam/models"
)

type StorageI interface {
	CloseDB()
	Categories() CategoriesRepoI
	Product() ProductRepoI
	Order() OrderRepoI
}

type CategoriesRepoI interface {
	Create(ctx context.Context, req *models.CreateCategories) (string, error)
	GetByPKey(ctx context.Context, req *models.CategoriesPrimarKey) (*models.Categories, error)
	GetList(ctx context.Context, req *models.GetListCategoriesRequest) (*models.GetListCategoriesResponse, error)
	Update(ctx context.Context, id string, req *models.UpdateCategories) (int64, error)
	Delete(ctx context.Context, req *models.CategoriesPrimarKey) error
}

type ProductRepoI interface {
	Create(ctx context.Context, req *models.CreateProduct) (string, error)
	GetByPKey(ctx context.Context, req *models.ProductPrimarKey) (*models.Product, error)
	GetList(ctx context.Context, req *models.GetListProductRequest) (*models.GetListProductResponse, error)
	Update(ctx context.Context, id string, req *models.UpdateProduct) (int64, error)
	Delete(ctx context.Context, req *models.ProductPrimarKey) error
}

type OrderRepoI interface {
	Create(ctx context.Context, req *models.CreateOrder) (string, error)
	GetByPKey(ctx context.Context, req *models.OrderPrimarKey) (*models.Order, error)
	GetList(ctx context.Context, req *models.GetListOrderRequest) (*models.GetListOrderResponse, error)
	Update(ctx context.Context, id string, req *models.UpdateOrder) (int64, error)
	Delete(ctx context.Context, req *models.OrderPrimarKey) error
}
