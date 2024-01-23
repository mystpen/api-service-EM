package user

import (
	"api-service/internal/repository/user"
	"api-service/internal/types"
)

type UserService struct {
	repo user.UserRepo
}

func NewUserService(repo user.UserRepo) *UserService {
	return &UserService{repo}
}

type UserSerInterface interface {
	CreateUser(userData *types.User) error
}

func (u *UserService) CreateUser(dataUser *types.User) (error){
	//реализация обогащения
	return u.repo.CreateUser(dataUser)
}