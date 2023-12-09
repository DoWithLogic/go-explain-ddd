package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/martinyonatann/go-explain-ddd/internal/products/delivery/http"
	"github.com/martinyonatann/go-explain-ddd/internal/products/repository"
	"github.com/martinyonatann/go-explain-ddd/internal/products/usecase"
)

func main() {
	echo := echo.New()

	productRepo := repository.NewProductRepository()
	productUC := usecase.NewProductUsecase(productRepo)
	productHandlers := http.NewProductHandler(productUC)

	echo.POST("/product", productHandlers.CreateProduct)

	log.Info(echo.Start(":8888"))
}
