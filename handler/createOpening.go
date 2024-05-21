package handler

import (
	"github.com/gin-gonic/gin"
)

// CREATE OPENING
func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := db.Create(&request).Error; err != nil {
		logger.ErrF("error creating opening: %v", err.Error())
		return
	}

	logger.InfoF("Request Recebido: %+v", request)
}
