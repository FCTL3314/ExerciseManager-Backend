package main

import (
	"ExerciseManager/api/router"
	"ExerciseManager/bootstrap"
	"log"
)

func main() {
	app := bootstrap.NewApplication()
	g := app.Gin
	db := app.DB
	cfg := app.Cfg

	router.RegisterRoutes(g, db)

	if err := g.Run(cfg.Server.Address); err != nil {
		log.Fatal(err)
	}
}
