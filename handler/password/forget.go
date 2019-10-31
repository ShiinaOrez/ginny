package password

import (
	"github.com/ShiinaOrez/ginny/handler"
	"github.com/ShiinaOrez/ginny/model"
	"github.com/ShiinaOrez/ginny/pkg/errno"
	"github.com/gin-gonic/gin"
)

type ForgetPasswordPayload struct {
	Username    string `json:"username"`
	NewPassword string `json:"new_password"`
}

func ForgetPassword(c *gin.Context) {
	var data ForgetPasswordPayload
	if err := c.BindJSON(&data); err != nil {
		handler.SendBadRequest(c, errno.PayloadBadRequest)
		return
	}
	err := model.UpdatePassword(data.Username, data.NewPassword)
	if err != nil {
		handler.SendError(c, errno.UserUpdatePasswordError)
		return
	}
	handler.SendResponse(c, struct {
		NewPassword string `json:"new_password"`
	}{data.NewPassword})
	return
}
