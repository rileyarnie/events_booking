package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rileyarnie/events_booking/models"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't parse event id. Try again"})
		return
	}
	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't get event. Try again"})
		return
	}

	err = event.Register(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't register for event. Try again"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "you have successfully registered for the event"})

}
func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't parse event id. Try again"})
		return
	}
	var event models.Event
	event.ID = eventId
	event.CancelRegistration(userId)

	context.JSON(http.StatusOK, gin.H{"message": "Registration has been cancelled."})
}
