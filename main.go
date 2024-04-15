package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hkm007/gopherly/db"
	"github.com/hkm007/gopherly/routes"
	"github.com/hkm007/gopherly/utils/constants"
)

func main() {

	// init database & server
	db.InitDB()
	server := gin.Default()

	// auth api routes
	routes.RegisterAuthRoutes(server)
	
	// event api routes
	routes.RegisterEventRoutes(server)

	server.Run(constants.PORT)
}
