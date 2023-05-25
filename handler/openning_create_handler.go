package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func OpeningCreateHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "POST Opening",
	})
}
