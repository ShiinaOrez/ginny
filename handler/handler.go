package handler

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func SendResponse(c *gin.Context, data interface{}) {
    c.JSON(http.StatusOK, data)
}

func SendUnauthorized(c *gin.Context) {
    c.AbortWithStatus(http.StatusUnauthorized)
}

func SendBadRequest(c *gin.Context) {
    c.AbortWithStatus(http.StatusBadRequest)
}

func SendNotFound(c *gin.Context) {
    c.AbortWithStatus(http.StatusNotFound)
}