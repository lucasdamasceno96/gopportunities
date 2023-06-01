package handler

import (
	"net/http"

	"gopportunities/schemas"

	"github.com/gin-gonic/gin"
)

func ShowOpeningsHandler(ctx *gin.Context){
		openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}

	sendSuccess(ctx, "list-openings", openings)
}
