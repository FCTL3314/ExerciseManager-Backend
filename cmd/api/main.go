package main

import (
	"ExerciseManager/api/router"
	"ExerciseManager/bootstrap"
	"fmt"
	"log"
)

func main() {
	app := bootstrap.NewApplication()

	router.RegisterRoutes(app.Router, app.DB, app.Cfg, app.LoggerGroup)

	fmt.Printf("Listening and serving HTTP on %s\n", app.Cfg.Server.Address)
	if err := app.Router.Run(app.Cfg.Server.Address); err != nil {
		log.Fatal(err)
	}
}
