package handler

import (
	"net/http"

	"github.com/Koakovski/gopportunities/schema"
	"github.com/gin-gonic/gin"
)

func OpeningCreateHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	if err := ctx.BindJSON(&request); err != nil {
		logger.Errorf("request body error: %v", err.Error())
		sendResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schema.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Link:     request.Link,
		Salary:   request.Salary,
		Remote:   *request.Remote,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("error creating opening: %v", err.Error())
		sendResponseError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendResponse(ctx, http.StatusCreated, opening)
}
