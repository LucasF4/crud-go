package router

import (
	"github.com/gin-gonic/gin"

	"api/handler"
)

func initializerRoutes(router *gin.Engine) {
	handler.InitializerHandler()

	// ROUTER GROUPS
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", handler.GetOpeningHandler)

		v1.POST("/opening", handler.CreateOpeningHandler)

		v1.PUT("/opening", handler.EditOpeningHandler)

		v1.DELETE("/opening", handler.DeleteOpeningHandler)
	}

}
