package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rileyarnie/events_booking/models"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		fmt.Printf("errroooooorrrrrr ,%v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't fetch events. Try again"})
		return
	}
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't parse request data"})
		return
	}

	event.UserID = 1
	err = event.Save()
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't create event. Try again"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": event})
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't parse event id. Try again"})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't fetch event. Try again"})
		return
	}
	context.JSON(http.StatusOK, event)

}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't parse event id. Try again"})
		return
	}

	_, err = models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't fetch event. Try again"})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't parse request data"})
		return
	}

	updatedEvent.ID = eventId

	err = updatedEvent.Update()

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't update event. Try again"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully"})

}

func deleteEvent(context *gin.Context) {
	id := context.Param("id")
	eventId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't parse event id"})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't fetch event"})
		return
	}

	err = event.Delete()

	if err != nil {
		fmt.Printf("here is the error %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't delete event "})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})

}
