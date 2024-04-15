package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hkm007/gopherly/db"
	"github.com/hkm007/gopherly/handlers"
	"github.com/hkm007/gopherly/utils/constants"
)

func main() {

	// init database & server
	db.InitDB()
	server := gin.Default()

	// api routes
	server.GET("/events", handlers.GetEvents)
	server.POST("/events", handlers.CreateEvent)

	server.Run(constants.PORT)
}
