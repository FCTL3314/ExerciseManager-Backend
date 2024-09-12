package main

import (
	"ExerciseManager/api/route"
	"ExerciseManager/bootstrap"
)

func main() {
	app := bootstrap.NewApplication()
	g := app.Gin
	db := app.DB
	cfg := app.Cfg

	route.Register(g, db)

	if err := g.Run(cfg.Server.Address); err != nil {
		panic(err)
	}
}
