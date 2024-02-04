package base

import "github.com/gin-gonic/gin"

type Config struct {
}

func (c *Config) List(context *gin.Context) {
	x := context.Query("x")
	context.JSON(200, gin.H{"msg": x})
	return
}

func (c *Config) Info(context *gin.Context) {
	x := context.Query("x")
	context.JSON(200, gin.H{"msg": x})
	return
}
