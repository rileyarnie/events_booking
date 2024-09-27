package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rileyarnie/events_booking/db"
	"github.com/rileyarnie/events_booking/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")

}
