package password

import (
	"github.com/gin-gonic/gin"
	"github.com/ShiinaOrez/ginny/model"
	"github.com/ShiinaOrez/ginny/handler"
	"github.com/ShiinaOrez/ginny/pkg/errno"
)

type ForgetPasswordPayload struct {
	Username    string   `json:"username"`
	NewPassword string   `json:"new_password"`
}

func ForgetPassword(c *gin.Context) {
	var data ForgetPasswordPayload
	if err := c.BindJSON(&data); err != nil {
		handler.SendBadRequest(c, errno.PayloadBadRequest)
		return
	}
	model.UpdatePassword(data.Username, data.NewPassword)
	handler.SendResponse(c, struct{
		NewPassword  string  `json:"new_password"`
	}{ data.NewPassword })
	return
}