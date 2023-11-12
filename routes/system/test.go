package system

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type TestRouter struct {
}

func (s *TestRouter) InitRouter(Router *gin.RouterGroup) {
	baseRouter := Router.Group("test")
	{
		baseRouter.GET("", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"code":    200,
				"message": "this is test",
				"data":    "2222",
			})
		})
	}
}
