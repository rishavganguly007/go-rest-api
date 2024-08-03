package routes

import (
	"net/http"
	"strconv"

	"example.com/go-rest-api/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context){
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse eventId"})
		return 
	} 

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse eventId"})
		return 
	} 

	err = event.Register(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not parse eventId"})
		return 
	} 
	context.JSON(http.StatusCreated, gin.H{"message": "registered"})
}

func cancelRegisterForEvent(context *gin.Context){
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse eventId"})
	}

	var event models.Event
	event.Id = eventId
	err = event.CancelRegister(userId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse eventId"})
	}

	context.JSON(http.StatusCreated, gin.H{"message": "cancelled"})
	
}