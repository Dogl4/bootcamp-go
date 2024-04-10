package main

import (
	"github/edrpereira_meli/go-web/cmd/server/handler"
	"github/edrpereira_meli/go-web/products"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load("../../.env")

	repository := products.NewRepository()
	service := products.NewService(repository)
	productsService := handler.NewProduct(service)

	server := gin.Default()
	productsGroup := server.Group("products")
	productsGroup.POST("/", productsService.Store())
	productsGroup.GET("/", productsService.GetAll())
	productsGroup.GET("/:id", productsService.GetById())
	productsGroup.PUT("/:id", productsService.Update())
	productsGroup.PATCH("/:id", productsService.UpdateNameAndPrice())
	productsGroup.DELETE("/:id", productsService.Delete())
	server.Run()
}
