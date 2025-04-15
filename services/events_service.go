package services

import (
	"errors"
	"fmt"
	"linechat/repository"
	"log"
	"time"
)

type eventsService struct {
	eventRepo repository.EventsRepository
}

func NewEventsService(eventsRepository repository.EventsRepository) EventsService {
	return &eventsService{eventRepo: eventsRepository}
}
func (s *eventsService) GetEvents(event FilterEvent) ([]*EventResponse, error) {

	//filter

	//init the loc
	loc, _ := time.LoadLocation("Asia/Bangkok")
	_ = loc

	page := 1
	if event.Page != 0 {
		page = event.Page
	}
	limit := 10
	if event.Limit != 0 {
		limit = event.Limit
	}
	sort := ""
	if event.Sort != "" {
		sort = event.Sort
	}
	eventStart := event.Start
	eventEnd := event.End
	switch event.Stages {
	case "all":
		// first month
		//now := time.Now()
		//eventStart = now.Add(-30 * 24 * time.Hour).Unix()
		//eventEnd = now.Add(30 * 24 * time.Hour).Unix()
		eventStart = -1
		eventEnd = -1

	case "upcoming":
		now := time.Now()
		eventStart = now.Add(-24 * time.Hour).Unix()
		eventEnd = now.Add(30 * 24 * time.Hour).Unix()
	case "ongoing":
		now := time.Now()
		eventStart = now.Add(-24 * time.Hour).Unix()
		end := now.Add(24 * time.Hour)
		eventEnd = end.Unix()

	}
	resEvent, err := s.eventRepo.EventsList(repository.EventFilter{
		Page:    page,
		Limit:   limit,
		Sort:    sort,
		Keyword: event.Keyword,
		Start:   eventStart,
		End:     eventEnd,
		Stages:  event.Stages,
		Status:  event.Status,
	})

	if err != nil {
		return nil, err
	}
	eventList := []*EventResponse{}
	for _, event := range resEvent {
		banners := []EventBanner{}

		startDate := time.Unix(event.StartDate, 0).Format("2006-01-02")
		startTime := time.Unix(event.StartTime, 0).Format("15:04")
		endDate := time.Unix(event.EndDate, 0).Format("2006-01-02")
		endTime := time.Unix(event.EndTime, 0).Format("15:04")

		for _, banner := range event.Banner {
			banners = append(banners, EventBanner{
				Url: banner.Url,
				Img: banner.Img,
			})
		}
		// members join on this event
		memberJoined := []*MemberJoinEventResponse{}
		for _, member := range event.Members {
			memberJoined = append(memberJoined, &MemberJoinEventResponse{
				EventId:  member.EventId,
				UserId:   member.UserId,
				JoinTime: member.JoinTime,
				Clinic:   member.Clinic,
				IsJoined: true,
			})

		}
		memberCheckin := make([]*EventCheckIn, 0, len(event.EventCheckIn))
		for _, checkin := range event.EventCheckIn {
			memberCheckin = append(memberCheckin, &EventCheckIn{
				EventId:      checkin.EventId,
				UserId:       checkin.UserId,
				CheckIn:      checkin.CheckIn,
				CheckOut:     false,
				CheckInTime:  checkin.CheckInTime,
				CheckOutTime: checkin.CheckOutTime,
				CheckInPlace: checkin.CheckInPlace,
				Clinic:       checkin.Clinic,
			})
		}

		item := EventResponse{
			EventId:      event.EventId,
			Title:        event.Title,
			Description:  event.Description,
			StartDate:    startDate,
			EndDate:      endDate,
			Place:        event.Place,
			StartTime:    startTime,
			EndTime:      endTime,
			Banner:       banners,
			Location:     event.Location,
			Status:       event.Status,
			Members:      memberJoined,
			Role:         event.Role,
			Published:    event.Published,
			EventType:    event.EventType,
			EventCheckIn: memberCheckin,
		}
		eventList = append(eventList, &item)
	}
	return eventList, nil
}
func (s *eventsService) GetEventById(eventId string) (*Event, error) {
	if eventId == "" {
		return nil, errors.New("event id is required")
	}
	resEvent, err := s.eventRepo.EventByEventId(eventId)
	if err != nil {
		return nil, err
	}
	event := &Event{
		EventId:     resEvent.EventId,
		Title:       resEvent.Title,
		Description: resEvent.Description,
		StartDate:   resEvent.StartDate,
		EndDate:     resEvent.EndDate,
		Place:       resEvent.Place,
		StartTime:   resEvent.StartTime,
		EndTime:     resEvent.EndTime,
		Location:    resEvent.Location,
	}
	return event, nil

}
func (s *eventsService) CreateEvent(event *EventImpl) error {

	//init the loc
	loc, _ := time.LoadLocation("Asia/Bangkok")
	//set timezone,
	now := time.Now().In(loc)

	bannerImpl := []repository.EventBanner{}
	for _, banner := range event.Banner {
		bannerImpl = append(bannerImpl, repository.EventBanner{
			Url: banner.Url,
			Img: banner.Img,
		})
	}

	// convert string datetime to Time
	// join date and time string to Time format
	eventStdStr := event.StartDate + " " + event.StartTime
	eventEndStr := event.StartDate + " " + event.StartTime
	bangKok, ok := time.LoadLocation("Asia/Bangkok")
	if ok != nil {
		fmt.Println("Error loading location:", ok)
	}
	timeLayout := "2006-01-02 15:04"
	eventStart, _ := time.ParseInLocation(timeLayout, eventStdStr, bangKok)
	eventEnd, _ := time.ParseInLocation(timeLayout, eventEndStr, bangKok)

	// Insert to repo
	err := s.eventRepo.CreateEvent(&repository.Event{
		EventId:      event.EventId,
		Title:        event.Title,
		Description:  event.Description,
		StartDate:    eventStart.Unix(),
		EndDate:      eventEnd.Unix(),
		Place:        event.Place,
		StartTime:    eventStart.Unix(),
		Banner:       bannerImpl,
		EndTime:      eventEnd.Unix(),
		Location:     event.Location,
		Status:       true,
		CreatedDate:  now.Unix(),
		UpdatedDate:  0,
		LineId:       event.LineId,
		LineName:     event.LineName,
		EventType:    event.EventType,
		Members:      make([]repository.MemberEventImpl, 0),
		EventCheckIn: make([]*repository.EventCheckIn, 0),
		Published:    event.Published,
		Role:         event.Role,
	})
	if err != nil {
		return err
	}
	return nil
}
func (s *eventsService) UpdateEvent(event *EventImpl) error {
	if event.EventId == "" {
		return errors.New("event id is required")
	}
	// update event with event repository
	//init the loc
	bangKok, ok := time.LoadLocation("Asia/Bangkok")
	if ok != nil {
		fmt.Println("Error loading location:", ok)
	}
	////set timezone,
	//now := time.Now().In(bangKok)
	layout := "2006-01-02 15:04"

	eventStart, _ := time.Parse(layout, event.StartDate)
	log.Println(time.Now().Format(layout))
	eventStartDate, _ := time.ParseInLocation(layout, event.StartDate, bangKok)
	eventStartTime, _ := time.ParseInLocation(layout, event.StartTime, bangKok)
	eventEndDate, _ := time.ParseInLocation(layout, event.EndDate, bangKok)
	eventEndTime, _ := time.ParseInLocation(layout, event.EndTime, bangKok)

	log.Println("Event Start:", event.StartDate, " ", event.EndDate)
	log.Println(eventStart)
	log.Println(eventStartDate)
	log.Println(eventStartTime)
	log.Println(eventEndDate)
	log.Println(eventEndTime)

	banerImpl := []repository.EventBanner{}
	for _, banner := range event.Banner {
		banerImpl = append(banerImpl, repository.EventBanner{
			Url: banner.Url,
			Img: banner.Img,
		})
	}

	err := s.eventRepo.UpdateEvent(event.EventId, &repository.Event{
		EventId:     event.EventId,
		Title:       event.Title,
		Description: event.Description,
		StartDate:   eventStartDate.Unix(),
		EndDate:     eventEndDate.Unix(),
		Place:       event.Place,
		StartTime:   eventStartTime.Unix(),
		EndTime:     eventEndTime.Unix(),
		Banner:      banerImpl,
		Location:    event.Location,
		Status:      event.Status,
		UpdatedDate: time.Now().Unix(),
		LineId:      event.LineId,
		LineName:    event.LineName,
		EventType:   event.EventType,
		Role:        event.Role,
		Published:   event.Published,
	})
	if err != nil {
		return err
	}

	return nil
}
func (s *eventsService) DeleteEvent(eventId string) error {

	if eventId == "" {
		return errors.New("event id is required")
	}
	err := s.eventRepo.DeleteEvent(eventId)
	if err != nil {
		return err
	}
	return nil
}
func (s *eventsService) CountEvent(fileterEvent FilterEvent) (int, error) {

	totalEvent, err := s.eventRepo.CountEvent(repository.EventFilter{})
	if err != nil {
		return 0, err
	}
	return totalEvent, nil
}
func (s *eventsService) CountMemberJoinEvents(fileterEvent FilterEvent) (int, error) {
	totalEvent, err := s.eventRepo.CountMemberJoinEvents(repository.EventFilter{})
	if err != nil {
		return 0, err
	}
	return totalEvent, nil
}
