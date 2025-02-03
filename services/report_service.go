package services

import (
	"linechat/repository"
	"time"
)

type reportService struct {
	eventRepo  repository.EventsRepository
	memberRepo repository.MemberRepository
}

func NewReportService(eventRepo repository.EventsRepository, memberRepo repository.MemberRepository) ReportService {
	return &reportService{
		eventRepo:  eventRepo,
		memberRepo: memberRepo,
	}
}

func (s *reportService) ExportMemberReport() ([]*MemberReport, error) {
	members, err := s.memberRepo.MemberList()
	if err != nil {
		return nil, err
	}
	memberReport := []*MemberReport{}
	for _, member := range members {

		// unix time to string time format
		regDateStr := time.Unix(member.RegisterDate, 0).Format("2006-01-02 15:04:05")

		inMember := MemberReport{
			MemberId:       member.LineId,
			Name:           member.Name,
			LastName:       member.LastName,
			Phone:          member.Phone,
			Email:          member.Email,
			Position:       member.Position,
			Organization:   member.Organization,
			Course:         member.Course,
			MemberType:     member.Med,
			ExtraInfo:      member.MedExtraInfo,
			EventId:        "",
			EventTitle:     "",
			RegisteredDate: regDateStr,
			LineName:       member.LineName,
			EventName:      "",
			LineId:         "",
			ClinicName:     "",
			Status:         member.Status,
		}
		memberReport = append(memberReport, &inMember)
	}

	return memberReport, nil

}
func (s *reportService) ExportEventReport() ([]*EventReport, error) {
	eventList := []*EventReport{}

	events, err := s.eventRepo.EventsList()
	if err != nil {
		return nil, err
	}

	memberList, err := s.memberRepo.MemberList()
	if err != nil {
		return nil, err
	}
	memberMap := []*Member{}
	for _, member := range memberList {
		memberMap = append(memberMap, &Member{
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

	memberJoin := []*Member{}
	for _, event := range events {
		// mapping members
		for _, member := range event.Members {
			for _, memMap := range memberMap {
				if memMap.LineId == member.LineId {
					memberJoin = append(memberJoin, memMap)
				}
			}
		}

		itemView := EventReport{
			EventId:     event.EventId,
			Title:       event.Title,
			Description: event.Description,
			Location:    event.Location,
			StartDate:   time.Unix(event.StartDate, 0).Format("2006-01-02"),
			StartTime:   time.Unix(event.StartTime, 0).Format("15:04:05"),
			EndTime:     time.Unix(event.EndTime, 0).Format("15:04:05"),
			EndDate:     time.Unix(event.EndDate, 0).Format("2006-01-02"),
			Status:      event.Status,
			Date:        time.Unix(event.StartDate, 0).Format("2006-01-02 15:04:05"),
			Members:     memberJoin,
			CountMember: len(memberJoin),
		}
		eventList = append(eventList, &itemView)
		memberJoin = []*Member{}
	}

	return eventList, nil

}
func (s *reportService) ExportClinicReport() ([]*ClinicReport, error) {
	return nil, nil
}
