package routes

import (
	"gin-api-client/controller/system/base"
	"gin-api-client/controller/system/user"
	"github.com/gin-gonic/gin"
)

type SystemRouter struct {
}

func (r *SystemRouter) InitSystemRouter(Router *gin.RouterGroup) {
	systemRouter := Router.Group("system")
	initUserRouter(systemRouter.Group("user"))
	initConfigRouter(systemRouter.Group("config"))
	initUserManageRouter(systemRouter.Group("user_manage"))
}

func initUserRouter(group *gin.RouterGroup) {
	systemUser := user.User{}
	group.GET("login", systemUser.Login)
	group.GET("info", systemUser.Info)
	group.GET("menu", systemUser.Menu)
}

func initConfigRouter(group *gin.RouterGroup) {
	config := base.Config{}
	group.GET("config/list", config.List)
	group.GET("config/info", config.Info)
}

func initUserManageRouter(group *gin.RouterGroup) {
	userManage := user.UserManage{}
	group.GET("user_manage/list", userManage.List)
	group.GET("user_manage/info", userManage.Info)
	group.GET("user_manage/del", userManage.Del)
	group.GET("user_manage/save", userManage.Save)
}
