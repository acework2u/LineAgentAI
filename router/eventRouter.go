package router

import (
	"github.com/gin-gonic/gin"
	"linechat/handler"
)

type EventRouter struct {
	eventHandler *handler.EventHandler
}

func NewEventRouter(eventHandler *handler.EventHandler) *EventRouter {
	return &EventRouter{eventHandler: eventHandler}
}

func (r *EventRouter) EventRouter(rg *gin.RouterGroup) {
	rg.GET("/events", r.eventHandler.GetEvents)
	rg.GET("/event/:eventId", r.eventHandler.GetEvent)
	rg.POST("/event", r.eventHandler.CreateEvent)
	rg.POST("/events", r.eventHandler.CreateEvent)
	rg.PUT("/event", r.eventHandler.UpdateEvent)
	rg.PUT("/events", r.eventHandler.UpdateEvent)
	rg.DELETE("/events/:eventId", r.eventHandler.DeleteEvent)
}
