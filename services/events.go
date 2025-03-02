package services

type EventsService interface {
	GetEvents() ([]*EventResponse, error)
	GetEventById(eventId string) (*Event, error)
	CreateEvent(event *EventImpl) error
	UpdateEvent(event *EventImpl) error
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
	EventId     string                     `json:"eventId"`
	Title       string                     `json:"title"`
	Description string                     `json:"description"`
	StartDate   string                     `json:"startDate,omitempty"`
	EndDate     string                     `json:"endDate"`
	Place       string                     `json:"place"`
	StartTime   string                     `json:"startTime"`
	Banner      []EventBanner              `json:"banner"`
	EndTime     string                     `json:"endTime"`
	Location    string                     `json:"location"`
	Status      bool                       `json:"status"`
	CreatedDate int64                      `json:"createdDate"`
	UpdatedDate int64                      `json:"updatedDate"`
	LineId      string                     `json:"lineId"`
	LineName    string                     `json:"lineName"`
	EventType   string                     `json:"eventType"`
	Members     []*MemberJoinEventResponse `json:"members"`
}
type MemberJoinEventResponse struct {
	EventId  string `json:"eventId"`
	UserId   string `json:"userId"`
	Clinic   string `json:"clinic"`
	JoinTime int64  `json:"joinTime,omitempty"`
	IsJoined bool   `json:"isJoined"`
}
type EventImpl struct {
	EventId      string             `json:"eventId" form:"eventId" binding:"required" validate:"required,min=3,max=20"`
	Title        string             `json:"title" form:"title" binding:"required" validate:"required,min=3,max=20"`
	Description  string             `json:"description" form:"description"  binding:"required"`
	StartDate    string             `json:"startDate,omitempty" form:"startDate" binding:"required"`
	EndDate      string             `json:"endDate" form:"endDate"`
	Place        string             `json:"place" form:"place"`
	StartTime    string             `json:"startTime" form:"startTime"`
	Banner       []EventBanner      `json:"banner" form:"banner"`
	Banners      []string           `json:"banners" form:"banners"`
	EndTime      string             `json:"endTime" form:"endTime"`
	Location     string             `json:"location" form:"location"`
	Status       bool               `json:"status" form:"status"`
	CreatedDate  int64              `json:"createdDate" form:"createdDate"`
	UpdatedDate  int64              `json:"updatedDate" form:"updatedDate"`
	LineId       string             `json:"lineId" form:"lineId"`
	LineName     string             `json:"lineName" form:"lineName"`
	EventType    string             `json:"eventType" form:"eventType"`
	Members      []*MemberJoinEvent `json:"members" form:"members"`
	EventCheckIn []*EventCheckIn    `json:"eventCheckIn,omitempty" form:"eventCheckIn"`
	Published    bool               `json:"Published" form:"Published"`
	Role         []string           `json:"role,omitempty" form:"role"`
}

type EventBanner struct {
	Url string `json:"url"`
	Img string `json:"img"`
}
