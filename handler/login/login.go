package login

import (
    "github.com/ShiinaOrez/ginny/handler"
    "github.com/ShiinaOrez/ginny/model"
    "github.com/gin-gonic/gin"
)

type LoginPayload struct {
    Username  string  `json:"username"`
    Password  string  `json:"password"`
}

func Login(c *gin.Context) {
    var data LoginPayload
    if err := c.BindJSON(&data); err != nil {
        handler.SendBadRequest(c)
        return
    }
    if !model.CheckPasswordValidate(data.Username, data.Password) {
        handler.SendUnauthorized(c)
        return
    }
    handler.SendResponse(c, "Successful!")
    return
}
