package user

import (
	"api-service/internal/repository/user"
	"api-service/internal/types"
	"fmt"
)

type UserService struct {
	repo user.UserRepo
}

func NewUserService(repo user.UserRepo) *UserService {
	return &UserService{repo}
}

type UserSerInterface interface {
	CreateUser(userData *types.User) error
	GetAllUsers(types.Filter) ([]*types.User, error)
}

func (u *UserService) CreateUser(dataUser *types.User) (error){
	//реализация обогащения
	fmt.Println("service",dataUser)
	return u.repo.CreateUser(dataUser)
}

func (u *UserService) GetAllUsers(filter types.Filter) ([]*types.User, error){ //TODO: filter
	return u.repo.GetAllUsers(filter)
}