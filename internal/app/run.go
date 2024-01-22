package app

import (
	"api-service/internal/handlers"
	"api-service/internal/repository"
	"api-service/internal/service"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func Run() {
	// db
	// TODO: logs
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connString := fmt.Sprintf("user=%v password=%v dbname=%v sslmode=disable", dbUser, dbPassword, dbName)
	db, err := sql.Open("postgres", connString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("PostgreSQL!")
	// repo, service
	repo := repository.NewRepository(db)
	service := service.NewService(repo)

	handler := handlers.NewHandler(service)

	log.Fatal(Server(handler.Routes()))
}
