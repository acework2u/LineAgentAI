package repository

type MemberRepository interface {
	CreateMember(member *Member) error
	GetMemberByLineId(id string) (*Member, error)
	UpdateMember(lineId string, member *Member) error
	DeleteMember(id string) error
	GetMembers(filter Filter) ([]*Member, error)
	CreateJoinEvent(event *JoinEventImpl) error
	GetJoinEvent(eventId string) (*JoinEventImpl, error)
	CheckJoinEvent(eventId string, userId string) (bool, error)
	MemberList() ([]*Member, error)
}

type Member struct {
	Title        string `bson:"title"`
	Name         string `bson:"name"`
	LastName     string `bson:"lastname"`
	PinCode      int    `bson:"pincode"`
	Email        string `bson:"email"`
	Phone        string `bson:"phone"`
	BirthDate    int64  `bson:"birthdate"`
	Med          string `bson:"med"`
	MedExtraInfo string `bson:"medextrainfo"`
	Organization string `bson:"organization"`
	Position     string `bson:"position"`
	Course       string `bson:"course"`
	LineId       string `bson:"lineid"`
	LineName     string `bson:"lineName"`
	Facebook     string `bson:"facebook"`
	Instagram    string `bson:"instagram"`
	FoodAllergy  string `bson:"foodallergy"`
	Religion     string `bson:"religion"`
	RegisterDate int64  `bson:"registerdate"`
	UpdatedDate  int64  `bson:"updateddate"`
	Status       bool   `bson:"status"`
}

type Filter struct {
	Page    int      `json:"page"`
	Limit   int      `json:"limit"`
	Sort    string   `json:"sort"`
	Keyword string   `json:"keyword"`
	Members []string `json:"members"`
}
type JoinEventImpl struct {
	EventId        string `json:"eventId" binding:"required" validate:"required,min=3,max=20"`
	UserId         string `json:"userId" binding:"required"`
	JoinTime       int64  `json:"joinTime,omitempty"`
	Name           string `json:"name" binding:"required" validate:"required,min=3,max=20"`
	LastName       string `json:"lastName" binding:"required" validate:"required,min=3,max=20"`
	Organization   string `json:"organization" binding:"required"`
	Position       string `json:"position" binding:"required"`
	Course         string `json:"course" binding:"required"`
	LineId         string `json:"lineId" binding:"required"`
	LineName       string `json:"lineName" binding:"required"`
	Tel            string `json:"tel" binding:"required" validate:"required,numeric,min=10,max=10"`
	ReferenceName  string `json:"referenceName" binding:"required" validate:"required,min=3,max=20"`
	ReferencePhone string `json:"referencePhone" binding:"required" validate:"required,numeric,min=10,max=10"`
	Clinic         string `json:"clinic" binding:"required"`
}
