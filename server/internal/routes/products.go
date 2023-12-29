package routes

import (
	"github.com/gin-gonic/gin"
	"sushi/internal/adapters/http"
	"sushi/internal/core/services"
	"sushi/internal/store"
)

func productRoutes(router *gin.RouterGroup, db store.Store, m gin.HandlerFunc) {
	product := router.Group("/product")

	service := services.NewProductService(&db)
	adapters := http.NewProductAdapters(&service)

	product.GET("/", m, adapters.GetAllProducts)
	product.GET("/:id", m, adapters.GetOneProduct)

}
