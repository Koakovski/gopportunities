package router

import (
	"github.com/Koakovski/gopportunities/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1RouterGroup := router.Group("/api/v1")

	{
		// GET ALL OPENINGS
		v1RouterGroup.GET("/opening", handler.OpeningGetAllHandler)
		// GET OPENING BY ID
		v1RouterGroup.GET("/opening/:openingId", handler.OpeningGetOneHandler)
		// CREATE NEW OPENING
		v1RouterGroup.POST("/opening", handler.OpeningCreateHandler)
		// UPDATE OPENING
		v1RouterGroup.PUT("/opening/:openingId", handler.OpeningUpdateHandler)
		// DELETE OPENING
		v1RouterGroup.DELETE("/opening/:openingId", handler.OpeningDeleteHandler)
	}
}
