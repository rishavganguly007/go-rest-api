package main

import (
	"net/http"

	"example.com/go-rest-api/models"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default() // configures a http server
	
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
	server.Run(":8080") //localhost port

}

func createEvent(context *gin.Context){
	var event models.Event
	context.ShouldBind(&event)
}
func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}