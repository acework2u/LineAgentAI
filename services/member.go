package services

type MemberService interface {
	GetMemberByLineId(lineId string) (member *Member, err error)
	GetMembers() ([]*Member, error)
}
