package bootstrap

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
)

type Application struct {
	Gin *gin.Engine
	DB  *gorm.DB
	Cfg *Config
}

func NewApplication() *Application {
	var app Application
	app.initConfig()
	app.initDB()
	app.setGinMode()
	app.initGin()
	return &app
}

func (app *Application) initConfig() {
	c, err := NewConfig()
	if err != nil {
		log.Fatal("Error during config loading. Please check if environmental files exist.")
	}
	app.Cfg = c
}

func (app *Application) initDB() {
	DBConnector := NewConnector(
		app.Cfg.DB.Name,
		app.Cfg.DB.User,
		app.Cfg.DB.Password,
		app.Cfg.DB.Host,
		app.Cfg.DB.Port,
	)
	db, err := DBConnector.Connect()
	if err != nil {
		log.Fatal("Error during database connection.")
	}
	app.DB = db
}

func (app *Application) setGinMode() {
	switch app.Cfg.Server.Mode {
	case "release":
		gin.SetMode(gin.ReleaseMode)
	case "test":
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.DebugMode)
	}
}

func (app *Application) initGin() {
	g := gin.Default()
	if err := g.SetTrustedProxies(app.Cfg.Server.TrustedProxies); err != nil {
		log.Fatal(err)
	}
	app.Gin = g
}
