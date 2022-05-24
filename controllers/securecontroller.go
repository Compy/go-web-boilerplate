package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "pong"})
}

func Hello(context *gin.Context) {
	context.JSON(http.StatusOK, context.MustGet("claims"))
}
