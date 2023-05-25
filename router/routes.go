package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1RouterGroup := router.Group("/api/v1")

	{
		// GET ALL OPENINGS
		v1RouterGroup.GET("/opening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "GET Openings",
			})
		})
		// GET OPENING BY ID
		v1RouterGroup.GET("/opening/:id", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "GET Opening",
			})
		})
		// CREATE NEW OPENING
		v1RouterGroup.POST("/opening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "POST Opening",
			})
		})
		// UPDATE OPENING
		v1RouterGroup.PUT("/opening/:id", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "PUT Opening",
			})
		})
		// DELETE OPENING
		v1RouterGroup.DELETE("/opening/:id", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "DELETE Opening",
			})
		})
	}
}
