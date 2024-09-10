package main

import (
	"ExerciseManager/internal/config"
	"ExerciseManager/internal/database"
	"fmt"
	"gorm.io/gorm"
	"log"
)

func main() {
	cfg := initConfig()
	db := initDatabase(cfg)

	fmt.Println(cfg)
	fmt.Println(db)
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
