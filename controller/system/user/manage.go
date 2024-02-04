package user

import "github.com/gin-gonic/gin"

type UserManage struct {
}

func (m *UserManage) List(context *gin.Context) {
	context.JSON(200, gin.H{"msg": "user_manage"})
	return
}

func (m *UserManage) Info(context *gin.Context) {
	context.JSON(200, gin.H{"msg": "user_manage"})
	return
}

func (m *UserManage) Del(context *gin.Context) {
	context.JSON(200, gin.H{"msg": "user_manage"})
	return
}

func (m *UserManage) Save(context *gin.Context) {
	context.JSON(200, gin.H{"msg": "user_manage"})
	return
}
