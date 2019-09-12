package register

import (
    "github.com/ShiinaOrez/ginny/model"
    "github.com/gin-gonic/gin"
)

type RegisterPayload struct {
    Username  string  `json:"username"`
    Password  string  `json:"password"`
}

func Register(c *gin.Context) {
    var data RegisterPayload
    if err := c.BindJSON(&data); err != nil {
        c.JSON(400, gin.H{
            "message": "Bad Request!",
        })
        return
    }
    if model.CheckUserByUsername(data.Username) {
        c.JSON(401, gin.H{
            "message": "User Already Existed!",
        })
        return
    }
    model.CreateUser(data.Username, data.Password)
    c.JSON(200, gin.H{
        "message": "Register Successful!",
    })
    return
}
