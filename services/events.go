package services

type EventsService interface {
	GetEvents() ([]*EventResponse, error)
	GetEventById(eventId string) (*Event, error)
	CreateEvent(event *EventImpl) error
	UpdateEvent(event *Event) error
	DeleteEvent(eventId string) error
}

type FilterEvent struct {
	Page    int    `json:"page"`
	Limit   int    `json:"limit"`
	Sort    string `json:"sort"`
	Keyword string `json:"keyword"`
}
type Event struct {
	EventId     string            `json:"eventId"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	StartDate   int64             `json:"startDate,omitempty"`
	EndDate     int64             `json:"endDate"`
	Place       string            `json:"place"`
	StartTime   int64             `json:"startTime"`
	Banner      []EventBanner     `json:"banner"`
	EndTime     int64             `json:"endTime"`
	Location    string            `json:"location"`
	Status      bool              `json:"status"`
	CreatedDate int64             `json:"createdDate"`
	UpdatedDate int64             `json:"updatedDate"`
	LineId      string            `json:"lineId"`
	LineName    string            `json:"lineName"`
	EventType   string            `json:"eventType"`
	Members     []MemberJoinEvent `json:"members"`
}
type EventResponse struct {
	EventId     string        `json:"eventId"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	StartDate   string        `json:"startDate,omitempty"`
	EndDate     string        `json:"endDate"`
	Place       string        `json:"place"`
	StartTime   string        `json:"startTime"`
	Banner      []EventBanner `json:"banner"`
	EndTime     string        `json:"endTime"`
	Location    string        `json:"location"`
	Status      bool          `json:"status"`
	CreatedDate int64         `json:"createdDate"`
	UpdatedDate int64         `json:"updatedDate"`
	LineId      string        `json:"lineId"`
	LineName    string        `json:"lineName"`
	EventType   string        `json:"eventType"`
}
type EventImpl struct {
	EventId     string             `json:"eventId" binding:"required" validate:"required,min=3,max=20"`
	Title       string             `json:"title" binding:"required" validate:"required,min=3,max=20"`
	Description string             `json:"description" binding:"required"`
	StartDate   string             `json:"startDate,omitempty" binding:"required"`
	EndDate     string             `json:"endDate" binding:"required"`
	Place       string             `json:"place" binding:"required"`
	StartTime   string             `json:"startTime" binding:"required"`
	Banner      []EventBanner      `json:"banner" binding:"required"`
	EndTime     string             `json:"endTime" binding:"required"`
	Location    string             `json:"location" binding:"required"`
	Status      bool               `json:"status"`
	CreatedDate int64              `json:"createdDate"`
	UpdatedDate int64              `json:"updatedDate"`
	LineId      string             `json:"lineId"`
	LineName    string             `json:"lineName"`
	EventType   string             `json:"eventType"`
	Members     []*MemberJoinEvent `json:"members"`
}

type EventBanner struct {
	Url string `json:"url"`
	Img string `json:"img"`
}
