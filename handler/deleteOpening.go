package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DELETE OPENING
func DeleteOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "GET Opening",
	})
}
