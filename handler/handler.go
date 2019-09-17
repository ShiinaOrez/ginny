package handler

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func SendResponse(c *gin.Context, data interface{}) {
    c.JSON(http.StatusOK, data)
}

func SendUnauthorized(c *gin.Context, err error) {
    c.AbortWithStatusJSON(http.StatusUnauthorized, err)
    c.Error(err)
}

func SendBadRequest(c *gin.Context, err error) {
    c.AbortWithStatusJSON(http.StatusBadRequest, err)
    c.Error(err)
}

func SendNotFound(c *gin.Context, err error) {
    c.AbortWithStatusJSON(http.StatusNotFound, err)
    c.Error(err)
}

func SendError(c *gin.Context, err error) {
    c.AbortWithStatusJSON(500, err)
    c.Error(err)
}