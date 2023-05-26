package handler

import (
	"fmt"
	"net/http"

	"github.com/Koakovski/gopportunities/schema"
	"github.com/gin-gonic/gin"
)

func OpeningUpdateHandler(ctx *gin.Context) {
	openingId := ctx.Param("openingId")

	if openingId == "" {
		err := errParamIsRequired("openingId", "param")
		logger.Errorf("error update opening: %v", err.Error())
		sendResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	request := UpdateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schema.Opening{}

	if err := db.First(&opening, openingId).Error; err != nil {
		logger.Errorf("error update opening: %v", err.Error())
		sendResponseError(ctx, http.StatusNotFound, fmt.Errorf("opening with id %s not found", openingId).Error())
		return
	}

	if request.Role != "" {
		opening.Role = request.Role
	}

	if request.Company != "" {
		opening.Company = request.Company
	}

	if request.Location != "" {
		opening.Location = request.Location
	}

	if request.Remote != nil {
		opening.Remote = *request.Remote
	}

	if request.Link != "" {
		opening.Link = request.Link
	}

	if request.Salary > 0 {
		opening.Salary = request.Salary
	}

	if err := db.Save(&opening).Error; err != nil {
		logger.Errorf("error update opening: %v", err.Error())
		sendResponseError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendResponse(ctx, http.StatusOK, opening)
}
