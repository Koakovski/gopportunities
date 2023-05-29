package handler

import (
	"github.com/Koakovski/gopportunities/schema"
	"github.com/gin-gonic/gin"
)

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

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode int    `json:"errorCode"`
}

type CreateOpeningResponse struct {
	Data schema.OpeningResponse `json:"data"`
}

type GetOneOpeningResponse struct {
	Data schema.OpeningResponse `json:"data"`
}

type GetAllOpeningResponse struct {
	Data []schema.OpeningResponse `json:"data"`
}
