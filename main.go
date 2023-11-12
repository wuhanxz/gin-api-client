package main

import (
	"gin-api-client/routes/system"
	"github.com/gin-gonic/gin"
)

func main() {

	ginEngine := gin.New()

	system.GroupApp.Init(ginEngine)

	ginServer := ginEngine.Run(":8000")
	if ginServer != nil {
		return
	}
}
