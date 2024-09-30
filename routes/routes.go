package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rileyarnie/events_booking/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	//configure middleware
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	//event routes
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	//user routes
	server.POST("/signup", signup)
	server.POST("/login", login)
	//register routes
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)
}
