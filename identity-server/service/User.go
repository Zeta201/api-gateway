package service

import (
	"github.com/Zeta201/identity-server/model"
	"github.com/Zeta201/identity-server/repository"
)

type UserService struct {
	repository repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return UserService{
		repository: r,
	}
}

func (service UserService) Save(user model.User) error {
	return service.repository.Save(user)
}

func (service UserService) LogIn(user model.User) (string, error) {
	return service.repository.LogIn(user)
}
