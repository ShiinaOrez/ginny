package register

import (
	"github.com/ShiinaOrez/ginny/handler"
	"github.com/ShiinaOrez/ginny/model"
	"github.com/ShiinaOrez/ginny/pkg/errno"
	"github.com/ShiinaOrez/ginny/util"
	"github.com/gin-gonic/gin"
)

type RegisterPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Register(c *gin.Context) {
	var data RegisterPayload
	if err := c.BindJSON(&data); err != nil {
		handler.SendBadRequest(c, errno.PayloadBadRequest)
		return
	}
	if model.CheckUserByUsername(data.Username) {
		handler.SendError(c, errno.UserAleadyExisted)
		return
	}
	passwordHash, salt, err := util.PasswordHash(data.Password)
	if err != nil {
		handler.SendError(c, errno.PasswordHashError)
		return
	}
	model.CreateUser(data.Username, passwordHash, salt)
	handler.SendResponse(c, "Successful!")
	return
}
