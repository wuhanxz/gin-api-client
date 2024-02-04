package main

import (
	"gin-api-client/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	ginEngine := gin.New()

	new(routes.RouterGroup).InitRouter(ginEngine)

	ginServer := ginEngine.Run(":8000")
	if ginServer != nil {
		return
	}
}
