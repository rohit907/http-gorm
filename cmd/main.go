package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rohit907/gin-gorm/config"
	"github.com/rohit907/gin-gorm/route"
)

func main() {
	config.Connect()
	router := gin.New()

	
	route.UseRoute(router)
	router.Run()
}
