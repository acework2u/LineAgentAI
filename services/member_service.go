package services

import (
	"errors"
	"linechat/repository"
)

type memberService struct {
	memberRepo repository.MemberRepository
}

func NewMemberService(memberRepo repository.MemberRepository) MemberService {
	return &memberService{memberRepo: memberRepo}
}
func (s *memberService) GetMemberByLineId(lineId string) (member *Member, err error) {

	if lineId == "" {
		return nil, errors.New("lineId is empty")
	}

	res, err := s.memberRepo.GetMemberByLineId(lineId)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return nil, errors.New("member not found")
		}
		return nil, err
	}
	member = &Member{
		Title:        res.Title,
		Name:         res.Name,
		LastName:     res.LastName,
		PinCode:      res.PinCode,
		Email:        res.Email,
		Phone:        res.Phone,
		BirthDate:    res.BirthDate,
		Med:          res.Med,
		MedExtraInfo: res.MedExtraInfo,
		Organization: res.Organization,
		Position:     res.Position,
		Course:       res.Course,
		LineId:       res.LineId,
		LineName:     res.LineName,
		Facebook:     res.Facebook,
		Instagram:    res.Instagram,
		FoodAllergy:  res.FoodAllergy,
		Religion:     res.Religion,
		RegisterDate: res.RegisterDate,
		UpdatedDate:  res.UpdatedDate,
		Status:       res.Status,
	}

	return member, nil
}
func (s *memberService) GetMembers() ([]*Member, error) {
	res, err := s.memberRepo.MemberList()
	if err != nil {
		return nil, err
	}
	members := []*Member{}
	for _, member := range res {
		members = append(members, &Member{
			Title:        member.Title,
			Name:         member.Name,
			LastName:     member.LastName,
			PinCode:      member.PinCode,
			Email:        member.Email,
			Phone:        member.Phone,
			BirthDate:    member.BirthDate,
			Med:          member.Med,
			MedExtraInfo: member.MedExtraInfo,
			Organization: member.Organization,
			Position:     member.Position,
			Course:       member.Course,
			LineId:       member.LineId,
			LineName:     member.LineName,
			Facebook:     member.Facebook,
			Instagram:    member.Instagram,
			FoodAllergy:  member.FoodAllergy,
			Religion:     member.Religion,
			RegisterDate: member.RegisterDate,
			UpdatedDate:  member.UpdatedDate,
			Status:       member.Status,
		})
	}

	return members, nil
}
