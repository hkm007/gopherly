package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hkm007/gopherly/handlers"
	"github.com/hkm007/gopherly/middlewares"
)

func RegisterEventRoutes(server *gin.Engine) {

	server.GET("/events", handlers.GetEvents)
	server.GET("/events/:id", handlers.GetEvent)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.IsLoggedIn)
	authenticated.POST("/events", handlers.CreateEvent)
	authenticated.PUT("/events/:id", handlers.UpdateEvent)
	authenticated.DELETE("/events/:id", handlers.DeleteEvent)
}
