package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"linechat/services"
	"linechat/utils"
	"time"
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
func (e *EventHandler) UpdateEvent(c *gin.Context) {

	eventUpdate := &services.EventImpl{}
	err := c.ShouldBindJSON(eventUpdate)
	cusErr := utils.NewCustomErrorHandler(c)
	if err != nil {
		cusErr.ValidateError(err)
		return
	}

	// convert datetime string to date format "YYYY-MM-DDD"

	stDate := fmt.Sprintf("%s %s", eventUpdate.StartDate, eventUpdate.StartTime)
	enDate := fmt.Sprintf("%s %s", eventUpdate.EndDate, eventUpdate.EndTime)

	dataEvent := services.EventImpl{
		EventId:     eventUpdate.EventId,
		Title:       eventUpdate.Title,
		Description: eventUpdate.Description,
		StartDate:   stDate,
		StartTime:   stDate,
		EndDate:     enDate,
		EndTime:     enDate,
		Place:       eventUpdate.Place,
		Banner:      eventUpdate.Banner,
		Location:    eventUpdate.Location,
		Status:      eventUpdate.Status,
		LineId:      eventUpdate.LineId,
		LineName:    eventUpdate.LineName,
		EventType:   eventUpdate.EventType,
		UpdatedDate: time.Now().Local().Unix(),
	}
	err = e.eventService.UpdateEvent(&dataEvent)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	// Success
	c.JSON(200, gin.H{"message": "update the event successfully"})

}
