package services

type ReportService interface {
	ExportMemberReport() ([]*MemberReport, error)
	ExportEventReport() ([]*EventReport, error)
	ExportClinicReport(eventId string) ([]*ClinicReport, error)
}

type ReportType string

const (
	EVEType  ReportType = "event"
	CLINType ReportType = "clinic"
	MEMBType ReportType = "member"
)

type ReportFilter struct {
	Type    ReportType
	Date    string
	Keyword string
	Page    int
	Limit   int
	Sort    string
	Status  bool
	LineId  string
	EventId string
}

type ReportRequest struct {
	MemberId string `json:"memberId"`
	EventId  string `json:"eventId"`
	LineId   string `bson:"lineId"`
	Type     string `json:"type"`
	Date     string `json:"date"`
	Keyword  string `json:"keyword"`
	Page     int    `json:"page"`
	Limit    int    `json:"limit"`
	Sort     string `json:"sort"`
	Status   bool   `json:"status"`
}
type EventReport struct {
	EventId     string    `json:"eventId"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartDate   string    `json:"startDate,omitempty"`
	EndDate     string    `json:"endDate"`
	StartTime   string    `json:"startTime"`
	EndTime     string    `json:"endTime"`
	EventType   string    `json:"eventType"`
	Location    string    `json:"location"`
	Status      bool      `json:"status"`
	Date        string    `json:"date"`
	Members     []*Member `json:"members"`
	CountMember int       `json:"countMember"`
}
type ClinicReport struct {
	ClinicId    string   `json:"_id" bson:"_id"`
	ClinicName  string   `json:"clinic"`
	CountEvent  int      `json:"countEvent"`
	CountMember int      `json:"countMember"`
	Status      bool     `json:"status"`
	Member      []Member `json:"member"`
}

type MemberReport struct {
	MemberId       string `json:"memberId"`
	Name           string `json:"name"`
	LastName       string `json:"lastName"`
	Phone          string `json:"phone"`
	Email          string `json:"email"`
	Position       string `json:"position"`
	Organization   string `json:"organization"`
	Course         string `json:"course"`
	MemberType     string `json:"memberType"`
	ExtraInfo      string `json:"extraInfo"`
	EventId        string `json:"eventId"`
	EventTitle     string `json:"eventTitle"`
	RegisteredDate string `json:"registeredDate"`
	LineName       string `json:"lineName"`
	EventName      string `json:"eventName"`
	LineId         string `json:"lineId"`
	ClinicName     string `json:"clinicName"`
	Status         bool   `json:"status"`
}
type MemberJoinEventReport struct {
	MemberId     string `json:"memberId"`
	Name         string `json:"name"`
	LastName     string `json:"lastName"`
	Phone        string `json:"phone"`
	Email        string `json:"email"`
	Position     string `json:"position"`
	Organization string `json:"organization"`
	Course       string `json:"course"`
	Med          string `json:"med"`
	ExtraInfo    string `json:"extraInfo"`
	EventId      string `json:"eventId"`
	EventTitle   string `json:"eventTitle"`
	JoinTime     string `json:"joinTime"`
	LineName     string `json:"lineName"`
	EventName    string `json:"eventName"`
	LineId       string `json:"lineId"`
	ClinicName   string `json:"clinicName"`
	Status       bool   `json:"status"`
}
