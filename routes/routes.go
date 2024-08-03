package routes

import (
	"example.com/go-rest-api/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/event/:id", getEvent)

	authenticated := server.Group("/")
	authenticated.Use(middleware.Authenticate)

	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegisterForEvent)
	
	server.POST("/signup", signup)
	server.POST("/login", login)
}