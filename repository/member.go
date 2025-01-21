package repository

type Member struct {
	Name         string `bson:"name"`
	LastName     string `bson:"lastname"`
	PinCode      int    `bson:"pincode"`
	Email        string `bson:"email"`
	Phone        string `bson:"phone"`
	BirthDate    int64  `bson:"birthdate"`
	Med          string `bson:"med"`
	Organization string `bson:"organization"`
	Position     string `bson:"position"`
	Course       string `bson:"course"`
	LineId       string `bson:"lineid"`
	Line         string `bson:"line"`
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

type MemberRepository interface {
	CreateMember(member *Member) error
	GetMemberByLineId(id string) (*Member, error)
	UpdateMember(lineId string, member *Member) error
	DeleteMember(id string) error
	GetMembers(filter Filter) ([]*Member, error)
}
