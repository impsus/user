package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddUserRoutes(group *gin.RouterGroup, db *gorm.DB) {
	group.GET("/user/:id", GetUser(db))
	group.POST("/user", AddUser(db))
	group.DELETE("/user/:id", DeleteUser(db))
}

type AddUserBody struct {
	Name string `json:"name"`
}

func AddUser(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		user := AddUserBody{}
		if err := c.BindJSON(&user); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
		}
	}
}

func GetUser(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {

	}

}

func DeleteUser(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {

	}
}
