package routes

import (
	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	System SystemRouter
	Api    ApiRouter
}

func (r *RouterGroup) InitRouter(ginEngine *gin.Engine) {
	ginGroup := ginEngine.Group("")
	//r.InitSystemRouter(ginGroup)
	//r.InitApiRouter(ginGroup)
	r.System.InitSystemRouter(ginGroup)
	r.Api.InitApiRouter(ginGroup)
}
