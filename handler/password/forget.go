package password

import (
	"github.com/gin-gonic/gin"
	"github.com/ShiinaOrez/ginny/model"
)

type ForgetPasswordPayload struct {
	Username    string   `json:"username"`
	NewPassword string   `json:"new_password"`
}

func ForgetPassword(c *gin.Context) {
	var data ForgetPasswordPayload
	if err := c.BindJSON(&data); err != nil {
		c.JSON(400, gin.H{
			"message": "Bad Request",
		})
		return
	}
	model.UpdatePassword(data.Username, data.NewPassword)
	c.JSON(200, gin.H{
		"message": "Update Password Successful.",
		"new_password": data.NewPassword,
	})
	return
}