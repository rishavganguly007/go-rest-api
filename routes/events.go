package routes

import (
	"net/http"
	"strconv"

	"example.com/go-rest-api/models"

	"github.com/gin-gonic/gin"
)

func createEvent(context *gin.Context){
	var event models.Event
	err := context.ShouldBind(&event)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not parse request data"})
		return
	}

	event.Id = 1
	event.UserId = 1

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not create event"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": event})
}
func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not parse request data"})
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": events})
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse eventId"})
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "id not available"})
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event fetched", "event": event})
}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse eventId"})
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "id not available"})
		return
	}

	var updated models.Event
	err = context.ShouldBind(&event)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not parse request data"})
		return
	}

	updated.Id = eventId
	err = updated.Update(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not parse request data"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event fetched", "event": event})
}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse eventId"})
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "id not available"})
		return
	}

	err = context.ShouldBind(&event)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not parse request data"})
		return
	}

	err = event.Delete(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not parse request data"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event fetched", "event": event})
}