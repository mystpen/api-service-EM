package service

import (
	"api-service/internal/repository"
	"api-service/internal/service/user"
)

type Service struct {
	UserService user.UserSerInterface
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		UserService: user.NewUserService(repo.UserRepo),
	}
}
