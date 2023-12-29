package services

import (
	"sushi/internal/core/ports"
)

type CartService struct {
	repo ports.CartRepository
}

func (cs *CartService) AddProduct(userID, productID int) error {
	return cs.repo.Add(userID, productID)
}

func (cs *CartService) DeleteProduct(userID, productID int) error {
	return cs.repo.Delete(userID, productID)
}
