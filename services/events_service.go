package services

import (
	"errors"
	"linechat/repository"
	"linechat/utils"
	"time"
)

type eventsService struct {
	eventRepo repository.EventsRepository
}

func NewEventsService(eventsRepository repository.EventsRepository) EventsService {
	return &eventsService{eventRepo: eventsRepository}
}
func (s *eventsService) GetEvents() ([]*EventResponse, error) {

	resEvent, err := s.eventRepo.EventsList()
	if err != nil {
		return nil, err
	}
	eventList := []*EventResponse{}

	// banner

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

		item := EventResponse{
			EventId:     event.EventId,
			Title:       event.Title,
			Description: event.Description,
			StartDate:   startDate,
			EndDate:     endDate,
			Place:       event.Place,
			StartTime:   startTime,
			EndTime:     endTime,
			Banner:      banners,
			Location:    event.Location,
			Status:      event.Status,
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

	bannerImpl := []repository.EventBanner{}
	for _, banner := range event.Banner {
		bannerImpl = append(bannerImpl, repository.EventBanner{
			Url: banner.Url,
			Img: banner.Img,
		})
	}

	// Insert to repo
	err := s.eventRepo.CreateEvent(&repository.Event{
		EventId:     event.EventId,
		Title:       event.Title,
		Description: event.Description,
		StartDate:   utils.DateToTime(event.StartDate).Unix(),
		EndDate:     utils.DateToTime(event.EndDate).Unix(),
		Place:       event.Place,
		StartTime:   utils.TimeToTime(event.StartTime).Unix(),
		Banner:      bannerImpl,
		EndTime:     utils.TimeToTime(event.EndTime).Unix(),
		Location:    event.Location,
		Status:      true,
		CreatedDate: time.Now().Unix(),
		UpdatedDate: 0,
		LineId:      event.LineId,
		LineName:    event.LineName,
		EventType:   event.EventType,
	})
	if err != nil {
		return err
	}
	return nil
}
func (s *eventsService) UpdateEvent(event *Event) error {
	if event.EventId == "" {
		return errors.New("event id is required")
	}
	// update event with event repository
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
		StartDate:   event.StartDate,
		EndDate:     event.EndDate,
		Place:       event.Place,
		StartTime:   event.StartTime,
		Banner:      banerImpl,
		EndTime:     event.EndTime,
		Location:    event.Location,
		Status:      event.Status,
		UpdatedDate: time.Now().Unix(),
		LineId:      event.LineId,
		LineName:    event.LineName,
		EventType:   event.EventType,
	})
	if err != nil {
		return err
	}

	return nil
}
func (s *eventsService) DeleteEvent(eventId string) error {
	return nil
}
