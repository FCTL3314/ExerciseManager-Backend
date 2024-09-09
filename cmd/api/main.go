package main

import (
	"ExerciseManager/internal/config"
	"fmt"
	"log"
)

func main() {
	cfg := initConfig()
	fmt.Println(cfg)
}

func initConfig() *config.Config {
	c, err := config.NewConfig()
	if err != nil {
		log.Fatal("Error loading environmental and configuration files.")
	}
	return c
}
