package handler

import (
	"fmt"
	"net/http"

	"github.com/Koakovski/gopportunities/schema"
	"github.com/gin-gonic/gin"
)

// @BasePath  /api/v1
// @Summary Delete opening
// @Description Delete a job Opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param openingId path int true "Opening ID"
// @Success 204
// @Failure 404 {object} ErrorResponse
// @Router /opening/{openingId} [delete]
func OpeningDeleteHandler(ctx *gin.Context) {
	openingId := ctx.Param("openingId")

	if openingId == "" {
		err := errParamIsRequired("openingId", "param")
		logger.Errorf("error deleting opening: %v", err.Error())
		sendResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schema.Opening{}

	if err := db.First(&opening, openingId).Error; err != nil {
		logger.Errorf("error deleting opening: %v", err.Error())
		sendResponseError(ctx, http.StatusNotFound, fmt.Errorf("opening with id %s not found", openingId).Error())
		return
	}

	if err := db.Delete(&opening).Error; err != nil {
		logger.Errorf("error deleting opening: %v", err.Error())
		sendResponseError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendResponse(ctx, http.StatusNoContent, nil)
}
