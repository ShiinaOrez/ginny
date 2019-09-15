package register

import (
    "github.com/ShiinaOrez/ginny/model"
    "github.com/ShiinaOrez/ginny/handler"
    "github.com/gin-gonic/gin"
)

type RegisterPayload struct {
    Username  string  `json:"username"`
    Password  string  `json:"password"`
}

func Register(c *gin.Context) {
    var data RegisterPayload
    if err := c.BindJSON(&data); err != nil {
        handler.SendBadRequest(c)
        return
    }
    if model.CheckUserByUsername(data.Username) {
        handler.SendUnauthorized(c)
        return
    }
    model.CreateUser(data.Username, data.Password)
    handler.SendResponse(c, "Successful!")
    return
}