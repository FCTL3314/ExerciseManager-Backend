package main

import (
	"ExerciseManager/internal/config"
	"ExerciseManager/internal/database"
	"ExerciseManager/internal/domain"
	"fmt"
	"gorm.io/gorm"
	"log"
)

func main() {
	cfg := initConfig()
	db := initDatabase(cfg)

	result := &domain.Exercise{}
	db.First(&result)

	fmt.Println(result)
}

func initConfig() *config.Config {
	c, err := config.NewConfig()
	if err != nil {
		log.Fatal("Error loading config. Please check if environmental files exists.")
	}
	return c
}

func initDatabase(cfg *config.Config) *gorm.DB {
	DBConnector := database.NewConnector(cfg.DB.Name, cfg.DB.User, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port)
	db, err := DBConnector.Connect()
	if err != nil {
		log.Fatal("Error connecting to database.")
	}
	return db
}
