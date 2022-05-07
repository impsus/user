package main

import (
	"github.com/gin-gonic/gin"
	health "github.com/impsus/user/lib/health"
	ginprometheus "github.com/zsais/go-gin-prometheus"
)

func init() {
}

func main() {
	router := gin.Default()
	ginprometheus.NewPrometheus("gin").Use(router)

	router.GET("/health", health.Health)

	router.Run(":2000")
}
