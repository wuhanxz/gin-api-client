package system

import (
	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	Base BaseRouter
	Test TestRouter
}

func (r *RouterGroup) Init(ginEngine *gin.Engine) {
	ginGroup := ginEngine.Group("system")
	r.Base.InitRouter(ginGroup)
	r.Test.InitRouter(ginGroup)
}

var GroupApp = new(RouterGroup)
