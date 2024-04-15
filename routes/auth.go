package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hkm007/gopherly/handlers"
)

func RegisterAuthRoutes(server *gin.Engine) {

	server.POST("/auth/login", handlers.Login)
	server.POST("/auth/signup", handlers.SignUp)
}
