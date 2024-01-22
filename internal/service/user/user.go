package user

import "api-service/internal/repository/user"

type UserService struct {
	repo user.UserRepo
}

func NewUserService(repo user.UserRepo) *UserService {
	return &UserService{repo}
}

type UserSerInterface interface {
	
}