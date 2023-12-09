package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/martinyonatann/go-explain-ddd/internal/products"
	"github.com/martinyonatann/go-explain-ddd/internal/products/dtos"
)

type handlers struct {
	uc products.ProductUsecase
}

func NewProductHandler(uc products.ProductUsecase) products.ProductHandlers {
	return &handlers{uc: uc}
}

func (h *handlers) CreateProduct(c echo.Context) error {
	err := h.uc.CreateProduct(c.Request().Context(), dtos.Product{})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, nil)
}
