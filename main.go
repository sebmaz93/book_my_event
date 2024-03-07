package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sebmaz93/book_my_event/db"
	"github.com/sebmaz93/book_my_event/models"
	"net/http"
)

func main() {
	server := gin.Default()
	db.Init()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	err := server.Run(":777")
	if err != nil {
		return
	}

}

func getEvents(ctx *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch", "error": err})
		return
	}
	ctx.JSON(http.StatusOK, events)
}

func createEvent(ctx *gin.Context) {
	var event models.Event
	err := ctx.ShouldBindJSON(&event)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data", "error": err})
	}

	err = event.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not create event", "error": err})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Event created", "data": event})
}
