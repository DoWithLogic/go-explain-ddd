package products

import (
	"github.com/labstack/echo/v4"
)

type ProductHandlers interface {
	CreateProduct(c echo.Context) error
}
