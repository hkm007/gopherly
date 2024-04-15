package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hkm007/gopherly/models"
)

func GetEvents(context *gin.Context) {

	var event models.Event
	events, err := event.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch events!",
		})
		return
	}
	context.JSON(http.StatusOK, events)
}

func GetEvent(context *gin.Context) {

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse event id!",
		})
		return
	}
	
	var event models.Event
	foundEvent, err := event.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch event!",
		})
		return
	}
	context.JSON(http.StatusOK, foundEvent)
}

func CreateEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	event.ID = 1
	event.UserID = 1
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not create events!",
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "Event created!",
		"event": event,
	})
}
