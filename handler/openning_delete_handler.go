package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func OpeningDeleteHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "DELETE Opening",
	})
}
