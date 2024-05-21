package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET ALL OPENING
func GetOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "GET Opening",
	})
}
