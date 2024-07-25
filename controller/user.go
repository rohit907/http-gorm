package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rohit907/http-gorm/config"
	"github.com/rohit907/http-gorm/models"
)

func UserController(c *gin.Context) {
	user := []models.User{}
	config.DB.Find(&user)
	c.JSON(http.StatusOK, user)
}
func CreateUser(c *gin.Context) {
	user := new(models.User)
	if err := c.ShouldBindJSON(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	config.DB.Create(&user)
	c.JSON(http.StatusOK, &user)
}
