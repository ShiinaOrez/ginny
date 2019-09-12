package router

import (
    "github.com/gin-gonic/gin"
    "github.com/ShiinaOrez/ginny/handler/register"
)

var Router *gin.Engine

func InitRouter() {
    Router = gin.Default()
    Router.POST("/register", register.Register)
    return
}
