package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"sushi/internal/core/domain"
	"sushi/internal/core/ports"
)

type productsControllers struct {
	ProductService ports.ProductService
}

func (ps *productsControllers) GetAllProducts(c *gin.Context) {
	allProducts, err := ps.ProductService.GetAllProducts()

	if err != nil {
		fmt.Println(err)
		c.JSON(500, domain.HttpResponse{
			Error: domain.Error{
				Status:  500,
				Message: "Internal error",
			},
		})
		return
	}

	c.JSON(200, domain.HttpResponse{
		Error: domain.Error{
			Status:  200,
			Message: "Products were found",
		},
		Data: allProducts,
	})
}

func (ps *productsControllers) GetOneProduct(c *gin.Context) {
	productID, _ := strconv.Atoi(c.Param("id"))
	oneProduct, err := ps.ProductService.GetOneProduct(productID)

	if err != nil {
		c.JSON(500, domain.HttpResponse{
			Error: domain.Error{
				Status:  500,
				Message: "Internal error",
			},
		})
		return
	}

	c.JSON(200, domain.HttpResponse{
		Error: domain.Error{
			Status:  200,
			Message: "Product was found",
		},
		Data: oneProduct,
	})
}

func NewProductAdapters(service ports.ProductService) productsControllers {
	return productsControllers{
		ProductService: service,
	}
}
