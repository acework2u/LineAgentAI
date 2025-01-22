package services

import (
	"errors"
	"fmt"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"linechat/conf"
	"linechat/repository"
	"log"
	"time"
)

type lineBotService struct {
	cfg        *conf.AppConfig
	bot        *messaging_api.MessagingApiAPI
	memberRepo repository.MemberRepository
}

func NewLineBotService(memberRepo repository.MemberRepository) LineBotService {
	cfg, _ := conf.NewAppConfig()
	bot, _ := messaging_api.NewMessagingApiAPI(cfg.LineApp.ChannelToken)

	return &lineBotService{
		cfg:        cfg,
		bot:        bot,
		memberRepo: memberRepo,
	}
}
func (s *lineBotService) SendTextMessage(text string) error {
	return nil
}
func (s *lineBotService) ReplyMessage(replyToken string, receiveMessage string) error {

	// member reply logic
	replayTxt := fmt.Sprintf("%s", receiveMessage)
	_, err := s.bot.ReplyMessage(&messaging_api.ReplyMessageRequest{
		ReplyToken: replyToken,
		Messages: []messaging_api.MessageInterface{
			messaging_api.TextMessage{
				Text: replayTxt,
			},
		},
	})

	if err != nil {
		return err
	}
	return nil
}
func (s *lineBotService) RegisterMember(member *Member) error {

	if err := validation(member); err != nil {
		return err
	}

	now := time.Now().Unix()
	err := s.memberRepo.CreateMember(&repository.Member{
		Name:         member.Name,
		LastName:     member.LastName,
		PinCode:      member.PinCode,
		Email:        member.Email,
		Phone:        member.Phone,
		BirthDate:    member.BirthDate,
		Med:          member.Med,
		Organization: member.Organization,
		Position:     member.Position,
		Course:       member.Course,
		LineId:       member.LineId,
		Line:         member.Line,
		Facebook:     member.Facebook,
		Instagram:    member.Instagram,
		FoodAllergy:  member.FoodAllergy,
		Religion:     member.Religion,
		RegisterDate: now,
		UpdatedDate:  0,
		Status:       true,
	})

	if err != nil {
		return err
	}

	return nil
}
func (s *lineBotService) GetLineProfile(userId string) (*UserInfo, error) {
	if userId == "" {
		return nil, errors.New("userId is required")
	}
	user, err := s.bot.GetProfile(userId)
	if err != nil {
		return nil, err
	}
	log.Println("user: getLineProfile")
	log.Println(user)

	rs, err := s.memberRepo.GetMemberByLineId(userId)
	if err != nil {
		return nil, err
	}

	memberInfo := Member{
		Name:         rs.Name,
		LastName:     rs.LastName,
		PinCode:      rs.PinCode,
		Email:        rs.Email,
		Phone:        rs.Phone,
		BirthDate:    rs.BirthDate,
		Med:          rs.Med,
		Organization: rs.Organization,
		Position:     rs.Position,
		Course:       rs.Course,
		LineId:       rs.LineId,
		Line:         rs.Line,
		Facebook:     rs.Facebook,
		Instagram:    rs.Instagram,
		FoodAllergy:  rs.FoodAllergy,
		Religion:     rs.Religion,
		RegisterDate: rs.RegisterDate,
		UpdatedDate:  rs.UpdatedDate,
		Status:       rs.Status,
	}

	userInfo := &UserInfo{
		UserID:     userId,
		Name:       rs.Name,
		LastName:   rs.LastName,
		PictureURL: user.PictureUrl,
		LineID:     user.UserId,
		Member:     memberInfo,
		Status:     rs.Status,
	}

	return userInfo, nil
}
func (s *lineBotService) UpdateMemberProfile(userId string, member *Member) error {
	if userId == "" {
		return errors.New("userId is required")
	}
	if err := validation(member); err != nil {
		return err
	}
	s.memberRepo.GetMemberByLineId(userId)

	err := s.memberRepo.UpdateMember(userId, &repository.Member{
		Name:         member.Name,
		LastName:     member.LastName,
		PinCode:      member.PinCode,
		Email:        member.Email,
		Phone:        member.Phone,
		BirthDate:    member.BirthDate,
		Med:          member.Med,
		Organization: member.Organization,
		Position:     member.Position,
		UpdatedDate:  time.Now().Unix(),
		Course:       member.Course,
		LineId:       member.LineId,
		Line:         member.Line,
		Facebook:     member.Facebook,
		Instagram:    member.Instagram,
		FoodAllergy:  member.FoodAllergy,
		Religion:     member.Religion,
		RegisterDate: member.RegisterDate,
		Status:       member.Status,
	})
	if err != nil {
		return err
	}
	return nil

}
func (s *lineBotService) CheckMemberRegister(userId string) (bool, error) {
	if userId == "" {
		return false, errors.New("userId is required")
	}
	rs, err := s.memberRepo.GetMemberByLineId(userId)
	if err != nil {
		return false, err
	}
	if rs.Status == true {
		return true, nil
	}
	return false, nil
}
func validation(member *Member) error {
	if member.Name == "" {
		return errors.New("name is required")
	}
	if member.LastName == "" {
		return errors.New("lastname is required")
	}
	if member.Email == "" {
		return errors.New("email is required")
	}
	if member.LineId == "" {
		return errors.New("lineId is required")
	}
	return nil
}
