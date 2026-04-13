package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID", "error": err.Error()})
		return
	}
	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve event", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, event)

}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve events", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	event.ID = 1
	event.UserId = 1
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create event", "error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created successfully", "event": event})
}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID", "error": err.Error()})
		return
	}
	_, err = models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve event", "error": err.Error()})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update event", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully", "event": updatedEvent})
}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID", "error": err.Error()})
		return
	}
	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve event", "error": err.Error()})
		return
	}
	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete event", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}
