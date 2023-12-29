package ports

import "sushi/internal/core/domain"

type UserRepository interface {
	Insert(user domain.User) (int, error)
	Select(key string, value any) (domain.User, error)
	Update(user domain.User) error
}

type UserService interface {
	LoginUser(user domain.User) (int, error)
	UpdateProfile(user domain.User) error
	Profile(id int) (domain.User, error)
}

type CartRepository interface {
	Add(userID, productID int) error
	Delete(userID, productID int) error
}

type CartService interface {
	AddProduct(userID, productID int) error
	DeleteProduct(userID, productID int) error
}

type ProductRepository interface {
	GetAll() ([]domain.Product, error)
	GetOne(productID int) (domain.Product, error)
}

type ProductService interface {
	GetAllProducts() ([]domain.Product, error)
	GetOneProduct(productID int) (domain.Product, error)
}
