package handler

import (
	"github.com/gin-gonic/gin"
	"linechat/services"
	"linechat/utils"
)

type EventHandler struct {
	eventService services.EventsService
}

func NewEventHandler(eventService services.EventsService) *EventHandler {
	return &EventHandler{
		eventService: eventService,
	}
}
func (e *EventHandler) GetEvents(c *gin.Context) {
	events, err := e.eventService.GetEvents()
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"events": events})
}
func (e *EventHandler) GetEvent(c *gin.Context) {
	eventId := c.Param("eventId")

	res, err := e.eventService.GetEventById(eventId)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": res})
}
func (e *EventHandler) CreateEvent(c *gin.Context) {
	event := &services.EventImpl{}
	err := c.ShouldBindJSON(event)
	cusErr := utils.NewCustomErrorHandler(c)
	if err != nil {
		cusErr.ValidateError(err)
		return
	}

	err = e.eventService.CreateEvent(event)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "creat the event successfully"})

}
