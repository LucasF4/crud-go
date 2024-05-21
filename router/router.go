package router

import (
	"github.com/gin-gonic/gin"
)

func Initializer() {
	// Initializer Router
	r := gin.Default()
	// Inicializer Routes
	initializerRoutes(r)
	// Run Server
	r.Run(":3001")
}
