package storage

import (
	"context"

	"exam/models"
)

type StorageI interface {
	CloseDB()
	Category() CategoryRepoI
	Product() ProductRepoI
	Order() OrderRepoI
}

type CategoryRepoI interface {
	Create(ctx context.Context, req *models.CreateCategory) (string, error)
	GetByPKey(ctx context.Context, req *models.CategoryPrimarKey) (*models.Category, error)
	GetList(ctx context.Context, req *models.GetListCategoryRequest) (*models.GetListCategoryResponse, error)
	Update(ctx context.Context, id string, req *models.UpdateCategory) (int64, error)
	Delete(ctx context.Context, req *models.CategoryPrimarKey) error
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
