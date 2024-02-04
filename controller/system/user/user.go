package user

import "github.com/gin-gonic/gin"

type User struct {
}

func (u *User) Login(context *gin.Context) {
	params := context.Param("")
	context.JSON(200, gin.H{"msg": params})
	return
}

func (u *User) Info(context *gin.Context) {

	context.JSON(200, gin.H{"msg": "user_manage"})
	return
}

func (u *User) Menu(context *gin.Context) {

}
