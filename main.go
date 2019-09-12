package main

import (
	"github.com/gin-gonic/gin"
    _ "github.com/go-sql-driver/mysql"
    "fmt"
    "database/sql"
)

type RegisterPayload struct {
	Username  string  `json:"username"`
	Password  string  `json:"password"`
}

func main() {
    db, err := GetDatabase("root", "ginny", "127.0.0.1", "3306", "go_blog")
    if err != nil {
    	panic("Link to database failed! Reason: " + err.Error())
    }
    defer db.Close()

	router := gin.Default()
	router.POST("/register", func(c *gin.Context) {
        var data RegisterPayload
        if err := c.BindJSON(&data); err != nil {
        	c.JSON(400, gin.H{
                "message": "Bad Request",
        	})
        	return
        }
        userID := 0
        if err = db.QueryRow("SELECT id FROM users WHERE username = ?", data.Username).Scan(&userID); err != nil {
    	    db.Query("INSERT INTO users (username, password) VALUES (?, ?)", data.Username, data.Password)
    	    c.JSON(200, gin.H{
                "message": data.Username + data.Password,
    	    })
            return
        } else {
            c.JSON(401, gin.H{
                "message": "User already existed.",
            })
        }
        return
	})

	router.Run()
}

func GetDatabase(username, password, host, port, dbname string) (*sql.DB, error) {
    address := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbname)
    db, err := sql.Open("mysql", address)
    if err != nil {
    	return nil, err
    }
    err = db.Ping()
    if err != nil {
    	return nil, err
    }
    return db, nil
}