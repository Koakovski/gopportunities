package handler

import "github.com/gin-gonic/gin"

func sendResponseError(ctx *gin.Context, code int, msg string) {
	ctx.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}

func sendResponse(ctx *gin.Context, code int, data any) {
	ctx.JSON(code, gin.H{
		"data": data,
	})
}
