package services

import (
	"sushi/internal/core/domain"
	"sushi/internal/core/ports"
)

type UserService struct {
	repo ports.UserRepository
}

func (us *UserService) LoginUser(user domain.User) (int, error) {
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

func (us *UserService) UpdateProfile(user domain.User) error {
	return us.repo.Update(user)
}

func (us *UserService) Profile(id int) (domain.User, error) {
	return us.repo.Select("id", id)
}

func NewUserService(repository ports.UserRepository) UserService {
	return UserService{
		repo: repository,
	}
}
