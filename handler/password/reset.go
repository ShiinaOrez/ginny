package password

import (
	"github.com/gin-gonic/gin"
	"github.com/ShiinaOrez/ginny/model"
	"github.com/ShiinaOrez/ginny/handler"
)

type ResetPasswordPayload struct {
	Username      string   `json:"username"`
	OldPassword   string   `json:"old_password"`
	NewPassword   string   `json:"new_password"`
}

func ResetPassword(c *gin.Context) {
	var data ResetPasswordPayload
	if err := c.BindJSON(&data); err != nil {
		handler.SendBadRequest(c)
		return
	}
	if !model.CheckPasswordValidate(data.Username, data.OldPassword) {
		handler.SendUnauthorized(c)
		return
	} else {
		model.UpdatePassword(data.Username, data.NewPassword)
		handler.SendResponse(c, struct {
			NewPassword  string  `json:"new_password"`
		}{ data.NewPassword })
	}
	return
}