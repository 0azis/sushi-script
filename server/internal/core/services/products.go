package services

import (
	"sushi/internal/core/domain"
	"sushi/internal/core/ports"
)

type productService struct {
	repo ports.ProductRepository
}

func (ps *productService) GetAllProducts() ([]domain.Product, error) {
	return ps.repo.GetAll()
}

func (ps *productService) GetOneProduct(productID int) (domain.Product, error) {
	return ps.repo.GetOne(productID)
}

func NewProductService(repository ports.ProductRepository) productService {
	return productService{
		repo: repository,
	}
}
