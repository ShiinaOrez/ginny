package password

import (
	"github.com/gin-gonic/gin"
	"github.com/ShiinaOrez/ginny/model"
)

type ResetPasswordPayload struct {
	Username      string   `json:"username"`
	OldPassword   string   `json:"old_password"`
	NewPassword   string   `json:"new_password"`
}

func ResetPassword(c *gin.Context) {
	var data ResetPasswordPayload
	if err := c.BindJSON(&data); err != nil {
		c.JSON(400, gin.H{
			"message": "Bad Request!",
		})
		return
	}
	if !model.CheckPasswordValidate(data.Username, data.OldPassword) {
		c.JSON(401, gin.H{
			"message": "Authentication Failed.",
		})
		return
	} else {
		model.UpdatePassword(data.Username, data.NewPassword)
		c.JSON(200, gin.H{
			"message": "Password Update Successful",
			"new_password": data.NewPassword,
		})
	}
	return
}