package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
    data := map[string]string{"message": "pong"}
    c.JSON(http.StatusOK, data)
}
