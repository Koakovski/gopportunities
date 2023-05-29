package handler

import (
	"fmt"
	"net/http"

	"github.com/Koakovski/gopportunities/schema"
	"github.com/gin-gonic/gin"
)

// @BasePath  /api/v1
// @Summary Get one opening
// @Description Get one job Opening by ID
// @Tags Openings
// @Accept json
// @Produce json
// @Param openingId path int true "Opening ID"
// @Success 200 {object} GetOneOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /opening/{openingId} [get]
func OpeningGetOneHandler(ctx *gin.Context) {
	openingId := ctx.Param("openingId")

	if openingId == "" {
		err := errParamIsRequired("openingId", "param")
		logger.Errorf("error get one opening: %v", err.Error())
		sendResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schema.Opening{}

	if err := db.First(&opening, openingId).Error; err != nil {
		logger.Errorf("error get one opening: %v", err.Error())
		sendResponseError(ctx, http.StatusNotFound, fmt.Errorf("opening with id %s not found", openingId).Error())
		return
	}

	sendResponse(ctx, http.StatusOK, opening)
}
