package repository

type EventsRepository interface {
	EventJoin(event *MemberEventImpl) error
	EventLeave(event *MemberEventImpl) error
	GetEvent(eventId string) (*MemberEventImpl, error)
	GetEvents(filter Filter) ([]*MemberEventImpl, error)
	CheckJoinEvent(eventId string, userId string) (bool, error)
	GetEventJoin(eventId string, userId string) (*MemberEventImpl, error)
	CheckInEvent(userId string, eventCheckIn *EventCheckIn) (bool, error)
}

type MemberEventImpl struct {
	EventId        string `bson:"eventId"`
	UserId         string `bson:"userId"`
	JoinTime       int64  `bson:"joinTime,omitempty"`
	Name           string `bson:"name"`
	LastName       string `bson:"lastName"`
	Organization   string `bson:"organization"`
	Position       string `bson:"position"`
	Course         string `bson:"course"`
	LineId         string `bson:"lineId"`
	LineName       string `bson:"lineName"`
	Tel            string `bson:"tel"`
	ReferenceName  string `bson:"`
	ReferencePhone string `bson:"referencePhone"`
	Clinic         string `bson:"clinic"`
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
