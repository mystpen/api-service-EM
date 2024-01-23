package user

import (
	"api-service/internal/types"
	"database/sql"
)

type UserDB struct {
	DB *sql.DB
}

func NewUserDB(db *sql.DB) *UserDB {
	return &UserDB{DB: db}
}

type UserRepo interface {
	CreateUser(userData *types.User) error
}

func (u *UserDB) CreateUser(dataUser *types.User)(error){
	return nil
}