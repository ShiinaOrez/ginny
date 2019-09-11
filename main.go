package main

import "github.com/gin-gonic/gin"

type RegisterPayload struct {
	Username  string  `json:"username"`
	Password  string  `json:"password"`
}

func main() {
	router := gin.Default()
	router.POST("/register", func(c *gin.Context) {
        var data RegisterPayload
        if err := c.BindJSON(&data); err != nil {
        	c.JSON(400, gin.H{
                "message": "Bad Request", // 如果绑定失败，那么我们认定它为一个坏请求，按照规范，状态码应该为400
        	})
        	return
        }
        c.JSON(200, gin.H{
            "result": data.Username + data.Password,
        })
        return
	})

	router.Run()
}