package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sebmaz93/book_my_event/models"
	"net/http"
	"strconv"
)

func registerForEvent(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event Id"})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch the event"})
		return
	}

	err = event.Register(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not register user for the event"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "registered for the event"})
}

func cancelRegistration(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event Id"})
		return
	}
	var event models.Event
	event.ID = eventId

	err = event.CancelRegistration(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not cancel registration for the event"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "cancelled registration"})
}
