package user

import (
	"api-service/internal/repository/user"
	"api-service/internal/types"
	"api-service/pkg"
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
	DeleteUser(userId int) error
	UpdateUser(updatedUser *types.User, userId int) error
}

func (u *UserService) CreateUser(dataUser *types.User) error {
	err := pkg.Parse(dataUser)
	if err != nil {
		return err
	}
	return u.repo.CreateUser(dataUser)
}

func (u *UserService) GetAllUsers(filter types.Filter) ([]*types.User, error) { // TODO: filter
	return u.repo.GetAllUsers(filter)
}

func (u *UserService) DeleteUser(userId int) error {
	return u.repo.DeleteUser(userId)
}

func (u *UserService) UpdateUser(updatedUser *types.User, userId int) error {
	return u.repo.UpdateUser(updatedUser, userId)
}
