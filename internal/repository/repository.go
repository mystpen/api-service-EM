package repository

import (
	"api-service/internal/repository/user"
	"database/sql"
)

type Repository struct {
	UserRepo user.UserRepo
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		UserRepo: user.NewUserDB(db),
	}
}
