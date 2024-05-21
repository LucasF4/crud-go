package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// EDIT OPENING
func EditOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "GET Opening",
	})
}
