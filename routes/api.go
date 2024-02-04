package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ApiRouter struct {
}

func (r *ApiRouter) InitApiRouter(Router *gin.RouterGroup) {
	apiRouter := Router.Group("api")
	{
		apiRouter.GET("info", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"code":    200,
				"message": "this is test",
				"data":    "2222",
			})
		})
	}
}
