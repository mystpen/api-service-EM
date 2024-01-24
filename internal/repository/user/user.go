package user

import (
	"api-service/internal/types"
	"database/sql"
	"fmt"
	"log"
)


//TODO: logs
var PageSize = 3

type UserDB struct {
	DB *sql.DB
}

func NewUserDB(db *sql.DB) *UserDB {
	return &UserDB{DB: db}
}

type UserRepo interface {
	CreateUser(userData *types.User) error
	GetAllUsers(types.Filter) ([]*types.User, error)
}

func (u *UserDB) CreateUser(dataUser *types.User) error {
	var nationalityID, genderID int
	err := u.DB.QueryRow(
		"INSERT INTO nationalities (nationality_name) VALUES ($1) ON CONFLICT (nationality_name) DO NOTHING RETURNING nationality_id",
		dataUser.Nationality,
	).Scan(&nationalityID)
	if err != nil {
		fmt.Println("create err:", err)
		return err
	}
	err = u.DB.QueryRow(
		"INSERT INTO gender (gender_name) VALUES ($1) ON CONFLICT (gender_name) DO NOTHING RETURNING gender_id", 
		dataUser.Gender).Scan(&genderID)
	if err != nil {
		fmt.Println("create err:", err)
		return err
	}
	_, err = u.DB.Exec("INSERT INTO people (name, surname, patronymic, age, gender_id, nationality_id) VALUES($1, $2, $3, $4, $5, $6)",
		dataUser.Name,
		dataUser.Surname,
		dataUser.Patronymic,
		dataUser.Age,
		nationalityID,
		genderID,
	)
	if err != nil {
		fmt.Println("create like err:", err)
		return err
	}

	return nil
}

func (u *UserDB) GetAllUsers(filter types.Filter) ([]*types.User, error) {
	offset := (filter.Page - 1) * PageSize
	query := "SELECT * FROM people INNER JOIN nationalities ON people.nationality_id=nationalities.nationality_id"
	if filter.Nationality != "" {
		query += fmt.Sprintf(" AND nationalities.nationality_name = '%s'", filter.Nationality)
	}

	query += fmt.Sprintf(" INNER JOIN genders ON people.gender_id=genders.gender_id")
	if filter.Gender != "" {
		query += fmt.Sprintf(" AND genders.gender_name = '%s'", filter.Gender)
	}

	if filter.Age != -1 {
		query += fmt.Sprintf("WHERE age=%v", filter.Age)
	}

	query += fmt.Sprintf(" LIMIT %d OFFSET %d", PageSize, offset)

	rows, err := u.DB.Query(query)
	if err != nil {
		log.Println("Failed to execute query:", err) ///
		return nil, err
	}
	defer rows.Close()

	users := []*types.User{}
	for rows.Next() {
		var user types.User
		if err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Surname,
			&user.Patronymic,
			&user.Age,
			&user.Gender,
			&user.Nationality); err != nil {
			log.Println("Failed to scan row:", err)
			return nil, err
		}
	}
	return users, nil
}

func (u *UserDB) DeleteUser(userId int) error {
	_, err := u.DB.Exec("DELETE FROM people WHERE id=$1", userId)
	if err != nil {
		log.Printf("Delete: %v\n", err)
	}
	return nil
}

func (u *UserDB) UpdateUser(updatedUser *types.User, userId int) error {
	query := fmt.Sprintf("UPDATE users SET name = $1, surname = $2, patronymic = $3, age = $4 WHERE person_id = $5")

	_, err := u.DB.Exec(query, updatedUser.Name, updatedUser.Surname, updatedUser.Patronymic, updatedUser.Age, userId)
	if err != nil {
		log.Println("Failed to execute query:", err)
		return err
	}
	return nil
}