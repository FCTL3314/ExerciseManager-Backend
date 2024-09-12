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
	c, err := NewConfig()
	if err != nil {
		log.Fatal("Error during config loading. Please check if environmental files exist.")
	}

	DBConnector := NewConnector(
		c.DB.Name,
		c.DB.User,
		c.DB.Password,
		c.DB.Host,
		c.DB.Port,
	)
	db, err := DBConnector.Connect()
	if err != nil {
		log.Fatal("Error during database connection.")
	}

	switch c.Server.Mode {
	case "release":
		gin.SetMode(gin.ReleaseMode)
	case "test":
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.DebugMode)
	}

	r := gin.Default()
	if err := r.SetTrustedProxies(c.Server.TrustedProxies); err != nil {
		panic(err)
	}

	return &Application{r, db, c}
}
