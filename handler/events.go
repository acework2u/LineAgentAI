package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"linechat/services"
	"linechat/utils"
	"log"
	"net/http"
	"strconv"
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

	stage := c.Query("stage")
	log.Println(stage)
	sort := c.Query("sort")
	log.Println(sort)
	keyword := c.Query("keyword")
	log.Println(keyword)
	Page := c.Query("page")
	if Page == "" {
		Page = "0"
	}

	offset, _ := strconv.Atoi(Page)

	log.Println(Page)
	Limit := c.Query("limit")
	log.Println(Limit)
	limitInt, _ := strconv.Atoi(Limit)
	log.Println(limitInt)
	startDate := c.Query("start_date")
	if startDate == "" {
		startDate = time.Now().Format("2006-01-02")
	}
	startDate = startDate + " 00:00:00"

	startDateInt64, _ := time.Parse("2006-01-02 15:04:05", startDate)
	log.Println(startDate)
	endDate := c.Query("end_date")
	if endDate == "" {
		endDate = time.Now().Format("2006-01-02")
	}
	endDate = endDate + " 23:59:59"
	endDateInt64, _ := time.Parse("2006-01-02 15:04:05", endDate)

	log.Println(endDate)
	//status := c.Query("status")

	events, err := e.eventService.GetEvents(services.FilterEvent{
		Page:    offset,
		Limit:   limitInt,
		Sort:    sort,
		Keyword: keyword,
		Stages:  stage,
		Start:   startDateInt64.Unix(),
		End:     endDateInt64.Unix(),
	})
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
	//event := &services.EventImpl{}
	//err := c.ShouldBindJSON(event)
	event := &services.EventImpl{}
	err := c.ShouldBind(event)
	//err := c.Bind(event)
	cusErr := utils.NewCustomErrorHandler(c)
	if err != nil {
		log.Println("error creating event")
		log.Println(err.Error())
		cusErr.ValidateError(err)
		return
	}
	log.Println(event)
	// upload files
	form, err := c.MultipartForm()
	if err == nil {
		files := form.File["banners"]
		for _, file := range files {
			f, err := file.Open()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			defer f.Close()

			url, err := utils.UploadFile(f, file.Filename)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			event.Banners = append(event.Banners, url)
		}
	}

	log.Println(event)

	err = e.eventService.CreateEvent(event)
	if err != nil {
		log.Println("error creating event")
		log.Println(err.Error())
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(201, gin.H{"message": "creat the event successfully"})

}
func (e *EventHandler) UpdateEvent(c *gin.Context) {

	eventUpdate := &services.EventImpl{}
	err := c.ShouldBindJSON(eventUpdate)
	cusErr := utils.NewCustomErrorHandler(c)
	if err != nil {
		log.Println(err.Error())
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
		Role:        eventUpdate.Role,
		Published:   eventUpdate.Published,
		UpdatedDate: time.Now().Local().Unix(),
	}
	err = e.eventService.UpdateEvent(&dataEvent)
	if err != nil {
		log.Println("error updating event")
		log.Println(err.Error())
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	// Success
	c.JSON(200, gin.H{"message": "update the event successfully"})

}
func (e *EventHandler) DeleteEvent(c *gin.Context) {
	eventId := c.Param("eventId")
	err := e.eventService.DeleteEvent(eventId)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "delete the event successfully"})
}
