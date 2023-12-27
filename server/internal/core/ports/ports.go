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
