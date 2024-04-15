package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hkm007/gopherly/handlers"
)

func RegisterEventRoutes(server *gin.Engine) {

	server.GET("/events", handlers.GetEvents)
	server.GET("/events/:id", handlers.GetEvent)
	server.POST("/events", handlers.CreateEvent)
	server.PUT("/events/:id", handlers.UpdateEvent)
	server.DELETE("/events/:id", handlers.DeleteEvent)
}
