package products

import (
	"context"

	"github.com/martinyonatann/go-explain-ddd/internal/products/dtos"
)

type ProductUsecase interface {
	CreateProduct(context.Context, dtos.Product) error
}
