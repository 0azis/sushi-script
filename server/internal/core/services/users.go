package services

import (
	"sushi/internal/core/domain"
	"sushi/internal/core/ports"
)

type userService struct {
	repo ports.UserRepository
}

func (us *userService) LoginUser(user domain.User) (int, error) {
	var userID int
	userDB, err := us.repo.Select("phone", user.Phone)

	if userDB.ID == 0 {
		return us.repo.Insert(user)
	}

	if err != nil {
		return userID, err
	}

	return userDB.ID, nil
}

func (us *userService) UpdateProfile(user domain.User) error {
	return us.repo.Update(user)
}

func (us *userService) Profile(id int) (domain.User, error) {
	return us.repo.Select("id", id)
}

func NewUserService(repository ports.UserRepository) userService {
	return userService{
		repo: repository,
	}
}
