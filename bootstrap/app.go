package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
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
		app.Cfg.Database.Name,
		app.Cfg.Database.User,
		app.Cfg.Database.Password,
		app.Cfg.Database.Host,
		app.Cfg.Database.Port,
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

	v, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		log.Fatalf("Failed to retrieve the Gin validator engine. Ensure that the validator is properly configured.")
	}

	if err := RegisterCustomValidators(v, app.Cfg); err != nil {
		log.Fatal(err)
	}

	app.Router = r
}

func (app *Application) initLoggerGroup() {
	userLogger := InitUserLogger()
	workoutLogger := InitWorkoutLogger()
	exerciseLogger := InitExerciseLogger()

	loggerGroup := NewLoggerGroup(
		&userLogger,
		&workoutLogger,
		&exerciseLogger,
	)
	app.LoggerGroup = loggerGroup
}
