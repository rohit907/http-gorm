package route

import (
	"github.com/gin-gonic/gin"
	"github.com/rohit907/http-gorm/controller"
)

func UseRoute(c *gin.Engine) {
	c.GET("/", controller.UserController)
	c.POST("/user", controller.CreateUser)
}
