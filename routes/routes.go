package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sebmaz93/book_my_event/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	// events
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	// authenticated
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)

	// users
	server.POST("/signup", signup)
	server.POST("/login", login)
}
