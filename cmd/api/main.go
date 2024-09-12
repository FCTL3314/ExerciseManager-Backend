package main

import (
	"ExerciseManager/internal/config"
	"ExerciseManager/internal/database"
	"ExerciseManager/internal/domain"
	"ExerciseManager/internal/repository"
	"ExerciseManager/internal/usecase"
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"log"
)

func main() {
	cfg := initConfig()
	db := initDatabase(cfg)

	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)

	users, err := userUsecase.List(
		&domain.Params{
			OrderParams: domain.OrderParams{
				Order: "username",
			},
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	usersJSON, err := json.MarshalIndent(users, "", "  ")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(usersJSON))

}

func initConfig() *config.Config {
	c, err := config.NewConfig()
	if err != nil {
		log.Fatal("Error loading config. Please check if environmental files exists.")
	}
	return c
}

func initDatabase(cfg *config.Config) *gorm.DB {
	DBConnector := database.NewConnector(
		cfg.DB.Name,
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
	)
	db, err := DBConnector.Connect()
	if err != nil {
		log.Fatal("Error connecting to database.")
	}
	return db
}
