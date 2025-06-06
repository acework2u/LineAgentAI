package repository

type EventsRepository interface {
	EventJoin(event *MemberEventImpl) error
	EventLeave(event *MemberEventImpl) error
	GetEvent(eventId string) (*MemberEventImpl, error)
	GetEvents(filter EventFilter) ([]*MemberEventImpl, error)
	CheckJoinEvent(eventId string, userId string) (bool, error)
	GetEventJoin(eventId string, userId string) (*MemberEventImpl, error)
	CheckInEvent(userId string, eventCheckIn *EventCheckIn) (bool, error)
	EventByUserId(userId string) ([]*Event, error)
	CreateEvent(event *Event) error
	UpdateEvent(eventId string, event *Event) error
	DeleteEvent(eventId string) error
	EventByEventId(eventId string) (*Event, error)
	EventsList(filter EventFilter) ([]*Event, error)
	EventsByClinic(eventId string) ([]*ClinicGroup, error)
	EventReport(filter *ReportFilter) ([]*Event, error)
	CountEvent(filter EventFilter) (int, error)
	CountMemberJoinEvents(filter EventFilter) (int, error)
	MembersJoinEvent(filter EventFilter) ([]*MemberJoinEvent, error)
}

type EventFilter struct {
	Page    int    `json:"page"`
	Limit   int    `json:"limit"`
	Total   int    `json:"total"`
	Sort    string `json:"sort"`
	Keyword string `json:"keyword"`
	Start   int64  `json:"start"`
	End     int64  `json:"end"`
	Stages  string `json:"stages"`
	Status  bool   `json:"status"`
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
	LineId       string             `bson:"lineId"`
	LineName     string             `bson:"lineName"`
	EventType    string             `bson:"eventType"`
	Members      []MemberEventImpl  `bson:"members,omitempty"`
	EventCheckIn []*EventCheckIn    `bson:"eventCheckIn,omitempty"`
	Published    bool               `bson:"Published"`
	Role         []string           `bson:"role,omitempty"`
	Certificate  []CertificateEvent `bson:"certificate,omitempty"`
	CreatedDate  int64              `bson:"createdDate,omitempty"`
	UpdatedDate  int64              `bson:"updatedDate,omitempty"`
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
	Role           []string        `bson:"role,omitempty"`
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
	Clinic       string `bson:"clinic"`
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

type MemberJoinEvent struct {
	EventTitle  string `bson:"eventTitle" json:"eventTitle"`
	EventDate   int64  `bson:"eventDate" json:"eventDate"`
	MemberCount int    `bson:"memberCount" json:"memberCount"`
}
