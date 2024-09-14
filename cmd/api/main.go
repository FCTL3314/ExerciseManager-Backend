package main

import (
	"ExerciseManager/api/router"
	"ExerciseManager/bootstrap"
	"log"
)

func main() {
	app := bootstrap.NewApplication()
	r := app.Router
	db := app.DB
	cfg := app.Cfg

	router.RegisterRoutes(r, db, cfg)

	if err := r.Run(cfg.Server.Address); err != nil {
		log.Fatal(err)
	}
}
