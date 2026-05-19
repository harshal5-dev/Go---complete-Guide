package routes

import (
	"example/com/rest-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse event id, try again later.",
		})
		return
	}

	event, err := models.GetEventById(id)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not fetch event.",
		})
		return
	}

	context.JSON(http.StatusOK, event)
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch events, try again later.",
		})
		return
	}

	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "could not parse",
		})
		return
	}

	event.UserID = 1
	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not create events, try again later.",
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "Event created!",
		"event":   event,
	})
}

func updateEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse event id, try again later.",
		})
		return
	}

	_, err = models.GetEventById(id)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not fetch event.",
		})
		return
	}

	var updateEvent models.Event
	err = context.ShouldBindJSON(&updateEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "could not parse",
		})
		return
	}

	updateEvent.Id = id
	err = updateEvent.Update()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not update events, try again later.",
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "Event updated!",
	})
}

func deleteEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse event id, try again later.",
		})
		return
	}

	event, err := models.GetEventById(id)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not fetch event.",
		})
		return
	}

	err = event.Delete()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not delete events, try again later.",
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "Event deleted!",
	})
}
