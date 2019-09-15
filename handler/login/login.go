package login

import (
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
        c.JSON(400, gin.H{
            "message": "Bad Request!",
        })
        return
    }
    if !model.CheckPasswordValidate(data.Username, data.Password) {
        c.JSON(401, gin.H{
            "message": "Authentication Failed.",
        })
        return
    } else {
        c.JSON(200, gin.H{
            "message": "Authentiaction Success.",
        })
    }
    return
}
