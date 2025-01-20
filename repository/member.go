package repository

type Member struct {
	Name         string
	LastName     string
	PinCode      int
	Email        string
	Phone        string
	BirthDate    int64
	Med          string
	Organization string
	Position     string
	Course       string
	LineId       string
	Line         string
	Facebook     string
	Instagram    string
	FoodAllergy  string
	Religion     string
	RegisterDate int64
	UpdatedDate  int64
	Status       bool
}

type Filter struct {
	Page    int
	Limit   int
	Sort    string
	Keyword string
}

type MemberRepository interface {
	CreateMember(member *Member) error
	GetMember(id string) (*Member, error)
	UpdateMember(member *Member) error
	DeleteMember(id string) error
	GetMembers() ([]*Member, error)
}
