package main

import (
	"ExerciseManager/api/controller"
	"ExerciseManager/bootstrap"
	"ExerciseManager/internal/repository"
	"ExerciseManager/internal/usecase"
	"fmt"
	"gorm.io/gorm"
	"log"
)

func main() {
	cfg := initConfig()
	db := initDatabase(cfg)

	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)

	fmt.Println(userController)

}

func initConfig() *bootstrap.Config {
	c, err := bootstrap.NewConfig()
	if err != nil {
		log.Fatal("Error loading config. Please check if environmental files exists.")
	}
	return c
}

func initDatabase(cfg *bootstrap.Config) *gorm.DB {
	DBConnector := bootstrap.NewConnector(
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
