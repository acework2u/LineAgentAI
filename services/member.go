package services

type MemberService interface {
	GetMemberByLineId(lineId string) (member *Member, err error)
	GetMembers() ([]*Member, error)
	UpdateMemberStatus(lineId string, status bool) error
}

type MemberStatus struct {
	Status bool `json:"status"`
}
