package user

import (
	"github.com/impsus/user/lib/uuid"

	_db "github.com/impsus/user/lib/db/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddUserRoutes(group *gin.RouterGroup, db *gorm.DB) {
	group.GET("/user/:id", GetUser(db))
	group.POST("/user", AddUser(db))
	group.DELETE("/user/:id", DeleteUser(db))
}

func AddUser(db *gorm.DB) func(c *gin.Context) {
	type AddUserResult struct {
		Username string `json:"username"`
	}

	return func(c *gin.Context) {
		username := uuid.GenUuidv4()
		db.Create(&_db.User{
			Name: username,
		})

		c.JSON(200, &AddUserResult{
			Username: username,
		})
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
