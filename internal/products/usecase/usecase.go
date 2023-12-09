package usecase

import (
	"context"

	"github.com/martinyonatann/go-explain-ddd/internal/products"
	"github.com/martinyonatann/go-explain-ddd/internal/products/dtos"
	"github.com/martinyonatann/go-explain-ddd/internal/products/entities"
)

type usecase struct {
	repo products.ProductRepository
}

func NewProductUsecase(r products.ProductRepository) products.ProductUsecase {
	return &usecase{repo: r}
}

func (uc *usecase) CreateProduct(ctx context.Context, request dtos.Product) error {

	return uc.repo.StoreProduct(ctx, entities.Product{})
}
