package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1RouterGroup := router.Group("/api/v1")

	{
		v1RouterGroup.GET("/opening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "Get Opening",
			})
		})
	}
}
