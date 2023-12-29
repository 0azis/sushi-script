package store

import (
	"encoding/json"
	"sushi/internal/core/domain"
)

func (store *Store) GetAll() ([]domain.Product, error) {
	allProducts := []domain.Product{}
	var nutritions []string

	err := store.DB.Select(&allProducts, `select id, name, image, category, price, weight, description from sushi_products`)

	if err != nil {
		return allProducts, err
	}

	err = store.DB.Select(&nutritions, `select json_pretty(nutrition) as nutrition from sushi_products`)

	for x := range nutritions {
		if err := json.Unmarshal([]byte(nutritions[x]), &allProducts[x].Nutrition); err != nil {
			return []domain.Product{}, err
		}
	}

	return allProducts, err

}

func (store *Store) GetOne(productID int) (domain.Product, error) {
	var oneProduct domain.Product

	err := store.DB.Get(&oneProduct, `select * from sushi_products where id = ?`, productID)

	return oneProduct, err
}
