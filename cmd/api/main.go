package main

import (
	"ExerciseManager/internal/config"
	"ExerciseManager/internal/database"
	"ExerciseManager/internal/repository"
	"fmt"
	"gorm.io/gorm"
	"log"
)

func main() {
	cfg := initConfig()
	db := initDatabase(cfg)

	UserRepository := repository.NewUserRepository(db)

	user, err := UserRepository.GetByUsername("TestUser")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(user)
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
