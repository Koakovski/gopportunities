package router

import (
	"github.com/Koakovski/gopportunities/handler"
	"github.com/gin-gonic/gin"

	docs "github.com/Koakovski/gopportunities/docs"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func initializeRoutes(router *gin.Engine) {
	basePath := "/api/v1"
	// Initialize Handler
	handler.InitalizarHandler()

	docs.SwaggerInfo.BasePath = basePath

	v1RouterGroup := router.Group(basePath)

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

	// Initialize Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
