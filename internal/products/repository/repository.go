package repository

import (
	"context"

	"github.com/martinyonatann/go-explain-ddd/internal/products"
	"github.com/martinyonatann/go-explain-ddd/internal/products/entities"
)

type repository struct {
}

func NewProductRepository() products.ProductRepository {
	return &repository{}
}

func (r *repository) StoreProduct(ctx context.Context, request entities.Product) error {
	return nil
}
