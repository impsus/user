package health

import "github.com/gin-gonic/gin"

type HealthJson struct {
	Status string `json:"status"`
}

func Health(c *gin.Context) {
	c.JSON(200, &HealthJson{
		Status: "healthy",
	})
}
