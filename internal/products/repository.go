package products

import (
	"context"

	"github.com/martinyonatann/go-explain-ddd/internal/products/entities"
)

type ProductRepository interface {
	StoreProduct(ctx context.Context, request entities.Product) error
}
