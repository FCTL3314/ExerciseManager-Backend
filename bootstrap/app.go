package bootstrap

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
)

type Application struct {
	Router      *gin.Engine
	DB          *gorm.DB
	Cfg         *Config
	LoggerGroup *LoggerGroup
}

func NewApplication() *Application {
	var app Application
	app.initConfig()
	app.initDB()
	app.initGin()
	app.initLoggerGroup()
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
	app.setGinMode()

	r := gin.Default()
	if err := r.SetTrustedProxies(app.Cfg.Server.TrustedProxies); err != nil {
		log.Fatal(err)
	}
	app.Router = r
}

func (app *Application) initLoggerGroup() {
	userLogger := InitUserLogger()
	workoutLogger := InitWorkoutLogger()

	loggerGroup := NewLoggerGroup(
		&userLogger,
		&workoutLogger,
	)
	app.LoggerGroup = loggerGroup
}
