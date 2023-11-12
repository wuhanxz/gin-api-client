package system

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseRouter struct {
}

func (s *BaseRouter) InitRouter(Router *gin.RouterGroup) {
	baseRouter := Router.Group("base")
	{
		baseRouter.GET("info", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"code":    200,
				"message": "this is info",
				"data":    "2222",
			})
		})
	}
}
