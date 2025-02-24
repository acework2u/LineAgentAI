package repository

type EventsRepository interface {
	EventJoin(event *MemberEventImpl) error
	EventLeave(event *MemberEventImpl) error
	GetEvent(eventId string) (*MemberEventImpl, error)
	GetEvents(filter Fillter) ([]*MemberEventImpl, error)
	CheckJoinEvent(eventId string, userId string) (bool, error)
	GetEventJoin(eventId string, userId string) (*MemberEventImpl, error)
	CheckInEvent(userId string, eventCheckIn *EventCheckIn) (bool, error)
	EventByUserId(userId string) ([]*Event, error)
	CreateEvent(event *Event) error
	UpdateEvent(eventId string, event *Event) error
	DeleteEvent(eventId string) error
	EventByEventId(eventId string) (*Event, error)
	EventsList() ([]*Event, error)
	EventsByClinic(eventId string) ([]*ClinicGroup, error)
	EventReport(filter *ReportFilter) ([]*Event, error)
}

type Fillter struct {
	Page    int    `json:"page"`
	Limit   int    `json:"limit"`
	Sort    string `json:"sort"`
	Keyword string `json:"keyword"`
}

type ReportFilter struct {
	StartDate int64  `json:"startDate"`
	EndDate   int64  `json:"endDate"`
	Keyword   string `json:"keyword"`
	Page      int    `json:"page"`
	Limit     int    `json:"limit"`
	Sort      string `json:"sort"`
}

type Event struct {
	EventId      string             `bson:"eventId"`
	Title        string             `bson:"title"`
	Description  string             `bson:"description"`
	StartDate    int64              `bson:"startDate,omitempty"`
	EndDate      int64              `bson:"endDate,omitempty"`
	Place        string             `bson:"place"`
	StartTime    int64              `bson:"startTime,omitempty"`
	Banner       []EventBanner      `bson:"banner,omitempty"`
	EndTime      int64              `bson:"endTime"`
	Location     string             `bson:"location"`
	Status       bool               `bson:"status"`
	CreatedDate  int64              `bson:"createdDate,omitempty"`
	UpdatedDate  int64              `bson:"updatedDate,omitempty"`
	LineId       string             `bson:"lineId"`
	LineName     string             `bson:"lineName"`
	EventType    string             `bson:"eventType"`
	Members      []MemberEventImpl  `bson:"members,omitempty"`
	EventCheckIn []*EventCheckIn    `bson:"eventCheckIn,omitempty"`
	Public       bool               `bson:"public"`
	Certificate  []CertificateEvent `bson:"certificate,omitempty"`
}
type EventUpdateImpl struct {
	EventId     string        `bson:"eventId"`
	Title       string        `bson:"title"`
	Description string        `bson:"description"`
	StartDate   int64         `bson:"startDate,omitempty"`
	EndDate     int64         `bson:"endDate"`
	Place       string        `bson:"place"`
	StartTime   int64         `bson:"startTime"`
	Banner      []EventBanner `bson:"banner"`
	EndTime     int64         `bson:"endTime"`
	Location    string        `bson:"location"`
	Status      bool          `bson:"status"`
	UpdatedDate int64         `bson:"updatedDate"`
	LineId      string        `bson:"lineId"`
	LineName    string        `bson:"lineName"`
	EventType   string        `bson:"eventType"`
}
type EventBanner struct {
	Url string `bson:"url"`
	Img string `bson:"img"`
}

type MemberEventImpl struct {
	EventId        string          `bson:"eventId"`
	UserId         string          `bson:"userId"`
	JoinTime       int64           `bson:"joinTime,omitempty"`
	Name           string          `bson:"name"`
	LastName       string          `bson:"lastName"`
	Organization   string          `bson:"organization"`
	Position       string          `bson:"position"`
	Course         string          `bson:"course"`
	LineId         string          `bson:"lineId"`
	LineName       string          `bson:"lineName"`
	Tel            string          `bson:"tel"`
	ReferenceName  string          `bson:"referenceName"`
	ReferencePhone string          `bson:"referencePhone"`
	Clinic         string          `bson:"clinic"`
	EventCheckIn   []*EventCheckIn `bson:"eventCheckIn,omitempty"`
}

type EventResponse struct {
	EventId          string        `json:"eventId"`
	EventName        string        `json:"evnetName"`
	EventDescription string        `json:"eventDescription"`
	EventStartDate   string        `json:"eventStartDate,omitempty"`
	EventEndDate     string        `json:"eventEndDate"`
	EventPlace       string        `json:"eventPlace"`
	EventStartTime   string        `json:"eventStartTime"`
	EventBanner      []EventBanner `json:"eventBanner"`
	EventEndTime     string        `json:"eventEndTime"`
	IsJoin           bool          `json:"isJoin"`
	EventStatus      bool          `json:"eventStatus"`
}

type EventCheckIn struct {
	EventId      string `bson:"eventId"`
	UserId       string `bson:"userId"`
	CheckIn      bool   `bson:"checkIn"`
	CheckOut     bool   `bson:"checkOut"`
	CheckInTime  int64  `bson:"checkInTime,omitempty"`
	CheckOutTime int64  `bson:"checkOutTime,omitempty"`
	CheckInPlace string `bson:"checkInPlace"`
}
type CertificateEvent struct {
	EventId string `bson:"eventId"`
	UserId  string `bson:"userId"`
}

// ClinicGroup represents the result of grouping members by clinic
type ClinicGroup struct {
	Clinic  string   `bson:"_id" json:"clinic"`
	Members []Member `bson:"members" json:"members"`
	Count   int      `bson:"count" json:"count"`
}
