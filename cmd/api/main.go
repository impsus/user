package main

import (
	"log"

	"github.com/gin-gonic/gin"
	db "github.com/impsus/user/lib/db"
	"github.com/impsus/user/lib/db/models"
	health "github.com/impsus/user/lib/health"
	"github.com/impsus/user/lib/user"
	ginprometheus "github.com/zsais/go-gin-prometheus"
	"gorm.io/gorm"
)

var serverConfig ServerConfig

type ServerConfig struct {
	db *gorm.DB
}

func init() {
	dbConn, err := db.GetConn()
	if err != nil {
		log.Panic(err)
	}
	serverConfig = ServerConfig{
		db: dbConn,
	}

	models.RegisterModels(serverConfig.db)

}

func main() {
	router := gin.Default()
	ginprometheus.NewPrometheus("gin").Use(router)

	router.GET("/health", health.Health)

	v1 := router.Group("v1")

	user.AddUserRoutes(v1, serverConfig.db)

	router.Run(":2000")
}
