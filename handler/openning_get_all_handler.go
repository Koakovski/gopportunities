package handler

import (
	"net/http"

	"github.com/Koakovski/gopportunities/schema"
	"github.com/gin-gonic/gin"
)

// @BasePath  /api/v1
// @Summary Get all openings
// @Description Get all job Openings
// @Tags Openings
// @Accept json
// @Produce json
// @Success 200 {object} GetAllOpeningResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [get]
func OpeningGetAllHandler(ctx *gin.Context) {
	openings := []schema.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		logger.Errorf("error get all openings: %v", err.Error())
		sendResponseError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendResponse(ctx, http.StatusOK, openings)
}
