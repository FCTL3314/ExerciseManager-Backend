package main

import (
	"ExerciseManager/api/router"
	"ExerciseManager/bootstrap"
	"log"
)

func main() {
	app := bootstrap.NewApplication()

	router.RegisterRoutes(app.Router, app.DB, app.Cfg, app.LoggerGroup)

	if err := app.Router.Run(app.Cfg.Server.Address); err != nil {
		log.Fatal(err)
	}
}
