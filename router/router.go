package router

import (
	"github.com/ShiinaOrez/ginny/handler/login"
	"github.com/ShiinaOrez/ginny/handler/password"
	"github.com/ShiinaOrez/ginny/handler/register"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func InitRouter() {
	Router = gin.Default()
	Router.POST("/register", register.Register)
	Router.POST("/login", login.Login)
	Router.POST("/password/reset", password.ResetPassword)
	Router.POST("/password/forget", password.ForgetPassword)
	return
}
