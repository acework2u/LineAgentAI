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
	eventRepo  repository.EventsRepository
}

func NewLineBotService(memberRepo repository.MemberRepository, eventRepo repository.EventsRepository) LineBotService {
	cfg, _ := conf.NewAppConfig()
	bot, _ := messaging_api.NewMessagingApiAPI(cfg.LineApp.ChannelToken)

	return &lineBotService{
		cfg:        cfg,
		bot:        bot,
		memberRepo: memberRepo,
		eventRepo:  eventRepo,
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
func (s *lineBotService) SendFlexMessage(replyToken string) error {

	// Create a Flex Message container using a Bubble layout
	flexMessage := messaging_api.FlexMessage{
		AltText: "this is a flex message",
		Contents: messaging_api.FlexBubble{
			Body: &messaging_api.FlexBox{
				Layout: messaging_api.FlexBoxLAYOUT_HORIZONTAL,
				Contents: []messaging_api.FlexComponentInterface{
					&messaging_api.FlexText{
						Text: "Hello",
					},
					&messaging_api.FlexText{
						Text: "World",
					},
				},
			},
		},
	}

	_, err := s.bot.ReplyMessage(&messaging_api.ReplyMessageRequest{
		ReplyToken: replyToken,
		Messages:   []messaging_api.MessageInterface{&flexMessage},
	})

	if err != nil {
		return err
	}

	return nil

}
func (s lineBotService) SendImageMessage(replyToken string, imageURL string) error {

	if imageURL == "" {
		return errors.New("imageURL is required")
	}
	news1 := []messaging_api.ImageCarouselColumn{
		{
			ImageUrl: imageURL,
			Action: messaging_api.UriAction{
				Label: "View detail",
				Uri:   "https://liff.line.me/2006793268-Rw1az8qr",
			},
		},
	}
	newList := []messaging_api.ImageCarouselColumn{
		{
			ImageUrl: imageURL,
			Action: messaging_api.UriAction{
				Label: "View detail",
				Uri:   "https://liff.line.me/2006793268-nJZqYeWE",
			},
		}, {
			ImageUrl: imageURL,
			Action: messaging_api.PostbackAction{
				Label: "Say hello",
				Data:  "action=buy&itemid=123",
			},
		}, {
			ImageUrl: imageURL,
			Action: messaging_api.MessageAction{
				Label: "Say hello",
				Text:  "hello",
			},
		}, {
			ImageUrl: imageURL,
			Action: messaging_api.DatetimePickerAction{
				Label: "Say hello",
				Mode:  "datetime",
				Data:  "action=buy&itemid=123",
			},
		},
	}
	_ = newList

	imageMessage := messaging_api.TemplateMessage{
		AltText: "news event carousel",
		Template: &messaging_api.ImageCarouselTemplate{
			Columns: news1,
		},
	}
	_, err := s.bot.ReplyMessage(&messaging_api.ReplyMessageRequest{
		ReplyToken: replyToken,
		Messages:   []messaging_api.MessageInterface{&imageMessage},
	})
	if err != nil {
		return err
	}
	return nil

}
func (s *lineBotService) SendFlexCarouselMessage(replyToken string) error {

	carouselMsg := messaging_api.FlexMessage{}
	carouselMsg.AltText = "this is a flex message"
	carouselMsg.Contents = messaging_api.FlexCarousel{
		Contents: []messaging_api.FlexBubble{
			{
				Body: &messaging_api.FlexBox{
					Layout: messaging_api.FlexBoxLAYOUT_VERTICAL,
					Contents: []messaging_api.FlexComponentInterface{
						&messaging_api.FlexText{
							Text: "Hello",
						},
						&messaging_api.FlexText{
							Text: "World",
						},
						&messaging_api.FlexText{
							Text: "Medical Volunteer",
						},
						&messaging_api.FlexImage{
							Url:         "https://www.linefriends.com/img/img_sec.jpg",
							Size:        "full",
							AspectRatio: "1.9:1",
						},
					},
				},
			},
		},
	}
	_, err := s.bot.ReplyMessage(&messaging_api.ReplyMessageRequest{
		ReplyToken: replyToken,
		Messages:   []messaging_api.MessageInterface{&carouselMsg},
	})
	if err != nil {
		return err
	}
	return nil
}
func (s *lineBotService) SendFlexJsonMessage(replyToken string, json string) error {
	//if json == "" {
	//	return errors.New("json is required")
	//}
	json = `{
  "type": "bubble",
  "hero": {
    "type": "image",
    "url": "https://developers-resource.landpress.line.me/fx/img/01_1_cafe.png",
    "size": "full",
    "aspectRatio": "20:13",
    "aspectMode": "cover",
    "action": {
      "type": "uri",
      "uri": "https://line.me/"
    }
  },
  "body": {
    "type": "box",
    "layout": "vertical",
    "contents": [
      {
        "type": "text",
        "text": "Brown Cafe",
        "weight": "bold",
        "size": "xl"
      },
      {
        "type": "box",
        "layout": "baseline",
        "margin": "md",
        "contents": [
          {
            "type": "icon",
            "size": "sm",
            "url": "https://developers-resource.landpress.line.me/fx/img/review_gold_star_28.png"
          },
          {
            "type": "icon",
            "size": "sm",
            "url": "https://developers-resource.landpress.line.me/fx/img/review_gold_star_28.png"
          },
          {
            "type": "icon",
            "size": "sm",
            "url": "https://developers-resource.landpress.line.me/fx/img/review_gold_star_28.png"
          },
          {
            "type": "icon",
            "size": "sm",
            "url": "https://developers-resource.landpress.line.me/fx/img/review_gold_star_28.png"
          },
          {
            "type": "icon",
            "size": "sm",
            "url": "https://developers-resource.landpress.line.me/fx/img/review_gray_star_28.png"
          },
          {
            "type": "text",
            "text": "4.0",
            "size": "sm",
            "color": "#999999",
            "margin": "md",
            "flex": 0
          }
        ]
      },
      {
        "type": "box",
        "layout": "vertical",
        "margin": "lg",
        "spacing": "sm",
        "contents": [
          {
            "type": "box",
            "layout": "baseline",
            "spacing": "sm",
            "contents": [
              {
                "type": "text",
                "text": "Place",
                "color": "#aaaaaa",
                "size": "sm",
                "flex": 1
              },
              {
                "type": "text",
                "text": "Flex Tower, 7-7-4 Midori-ku, Tokyo",
                "wrap": true,
                "color": "#666666",
                "size": "sm",
                "flex": 5
              }
            ]
          },
          {
            "type": "box",
            "layout": "baseline",
            "spacing": "sm",
            "contents": [
              {
                "type": "text",
                "text": "Time",
                "color": "#aaaaaa",
                "size": "sm",
                "flex": 1
              },
              {
                "type": "text",
                "text": "10:00 - 23:00",
                "wrap": true,
                "color": "#666666",
                "size": "sm",
                "flex": 5
              }
            ]
          }
        ]
      }
    ]
  },
  "footer": {
    "type": "box",
    "layout": "vertical",
    "spacing": "sm",
    "contents": [
      {
        "type": "button",
        "style": "link",
        "height": "sm",
        "action": {
          "type": "uri",
          "label": "CALL",
          "uri": "https://line.me/"
        }
      },
      {
        "type": "button",
        "style": "link",
        "height": "sm",
        "action": {
          "type": "uri",
          "label": "WEBSITE",
          "uri": "https://line.me/"
        }
      },
      {
        "type": "box",
        "layout": "vertical",
        "contents": [],
        "margin": "sm"
      }
    ],
    "flex": 0
  }
}`
	//contents, err := messaging_api.UnmarshalFlexContainer([]byte(jsonString))
	contents, err := messaging_api.UnmarshalFlexContainer([]byte(json))
	if err != nil {
		return err
	}
	_, err = s.bot.ReplyMessage(&messaging_api.ReplyMessageRequest{
		ReplyToken: replyToken,
		Messages: []messaging_api.MessageInterface{
			&messaging_api.FlexMessage{
				AltText:  "this is a flex message",
				Contents: contents,
			},
		},
	})
	if err != nil {
		return err
	}
	return nil

}
func (s *lineBotService) SendQuickReplyMessage(replyToken string) error {
	quickReplyMessage := messaging_api.TextMessage{
		Text: "Hello, world",
		QuickReply: &messaging_api.QuickReply{
			Items: []messaging_api.QuickReplyItem{
				{
					ImageUrl: "https://scdn.line-apps.com/n/channel_devcenter/img/fx/01_1_cafe.png",
					Action: messaging_api.UriAction{
						Label: "View detail",
						Uri:   "https://dca3a8ac633b.ngrok.app/register",
					},
				},
				{
					Action: messaging_api.PostbackAction{

						Label: "Say hello",
						Data:  "action=buy&itemid=123",
					},
				},
			},
		},
	}
	_, err := s.bot.ReplyMessage(&messaging_api.ReplyMessageRequest{
		ReplyToken: replyToken,
		Messages:   []messaging_api.MessageInterface{&quickReplyMessage},
	})
	if err != nil {
		return err
	}

	return nil
}
func (s *lineBotService) SendDateTimePickerMessage(replyToken string) error {
	dateTimePickerMessage := messaging_api.TemplateMessage{
		AltText: "Datetime Picker",
		Template: &messaging_api.ButtonsTemplate{
			Text: "When do you want to order?",
			Actions: []messaging_api.ActionInterface{
				&messaging_api.DatetimePickerAction{
					Label: "Date",
					Mode:  messaging_api.DatetimePickerActionMODE_DATE,
					Data:  "DATE",
				},
				&messaging_api.DatetimePickerAction{
					Label: "Time",
					Mode:  messaging_api.DatetimePickerActionMODE_TIME,
					Data:  "TIME",
				},
				&messaging_api.DatetimePickerAction{
					Label: "Date & Time",
					Mode:  messaging_api.DatetimePickerActionMODE_DATETIME,
					Data:  "DATETIME",
				},
			},
		},
	}

	_, err := s.bot.ReplyMessage(&messaging_api.ReplyMessageRequest{
		ReplyToken: replyToken,
		Messages:   []messaging_api.MessageInterface{&dateTimePickerMessage},
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

	//log.Println("user: getLineProfile")
	//log.Println(user)

	rs, err := s.memberRepo.GetMemberByLineId(userId)
	if err != nil {
		return nil, err
	}

	memberInfo := Member{
		Title:        rs.Title,
		Name:         rs.Name,
		LastName:     rs.LastName,
		PinCode:      rs.PinCode,
		Email:        rs.Email,
		Phone:        rs.Phone,
		BirthDate:    rs.BirthDate,
		Med:          rs.Med,
		MedExtraInfo: rs.MedExtraInfo,
		Organization: rs.Organization,
		Position:     rs.Position,
		Course:       rs.Course,
		LineId:       rs.LineId,
		LineName:     rs.LineName,
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

	member.Status = true
	err := s.memberRepo.UpdateMember(userId, &repository.Member{
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
		UpdatedDate:  time.Now().Unix(),
		Course:       member.Course,
		LineId:       member.LineId,
		LineName:     member.LineName,
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
func (s *lineBotService) EventJoin(event *JoinEventImpl) error {
	if len(event.EventId) == 0 {
		return errors.New("event id is required")
	}
	//if len(event.Name) == 0 || len(event.LastName) == 0 {
	//	return errors.New("name and lastname is required")
	//}

	//init the loc
	loc, _ := time.LoadLocation("Asia/Bangkok")
	//set timezone,
	now := time.Now().In(loc)

	err := s.eventRepo.EventJoin(&repository.MemberEventImpl{
		EventId:      event.EventId,
		UserId:       event.UserId,
		JoinTime:     now.Unix(),
		Name:         event.Name,
		LastName:     event.LastName,
		Organization: event.Organization,
		Position:     event.Position,
		Course:       event.Course,
		LineId:       event.LineId,
		Tel:          event.Phone,
		Clinic:       event.Clinic,
	})

	if err != nil {
		return err
	}
	return nil
}
func (s *lineBotService) CheckEventJoin(eventId string, userId string) (bool, error) {

	if eventId == "" {
		return false, errors.New("event id is required")
	}
	if userId == "" {
		return false, errors.New("user id is required")
	}
	rs, err := s.eventRepo.CheckJoinEvent(eventId, userId)
	if err != nil {
		return false, err
	}
	if rs {
		return true, nil
	}

	return false, nil
}
func (s *lineBotService) GetEventJoin(eventId string, userId string) (*MemberJoinEvent, error) {

	if eventId == "" {
		return nil, errors.New("event id is required")
	}
	if userId == "" {
		return nil, errors.New("user id is required")
	}

	res, err := s.eventRepo.GetEventJoin(eventId, userId)
	if err != nil {
		return nil, err
	}
	//joinTimeStr := time.Unix(res.JoinTime, 0).Format("2006-01-02 15:04:05")
	joinTimeStr := time.Unix(res.JoinTime, 0).Local().Format("2006-01-02 15:04:05")

	// bind checkin the events
	memberCheckIn := []*EventCheckIn{}
	if res.EventCheckIn != nil {
		for _, v := range res.EventCheckIn {
			if v.UserId == userId {
				checkIn := &EventCheckIn{
					CheckIn:      v.CheckIn,
					CheckInTime:  v.CheckInTime,
					CheckInPlace: v.CheckInPlace,
					CheckOut:     v.CheckOut,
					CheckOutTime: v.CheckOutTime,
				}
				memberCheckIn = append(memberCheckIn, checkIn)
			}
		}
	}

	memberJoinEvent := &MemberJoinEvent{
		EventId:        res.EventId,
		TitleEvent:     "",
		UserId:         res.UserId,
		JoinTime:       joinTimeStr,
		Name:           res.Name,
		LastName:       res.LastName,
		Organization:   res.Organization,
		Position:       res.Position,
		Course:         res.Course,
		LineId:         res.LineId,
		LineName:       res.LineName,
		Tel:            res.Tel,
		ReferenceName:  res.ReferenceName,
		ReferencePhone: res.ReferencePhone,
		Clinic:         res.Clinic,
		EventCheckIn:   memberCheckIn,
	}

	return memberJoinEvent, nil
}
func (s *lineBotService) CheckInEvent(eventCheckIn *EventCheckIn) (bool, error) {

	if eventCheckIn.EventId == "" {
		return false, errors.New("event id is required")
	}
	if eventCheckIn.UserId == "" {
		return false, errors.New("user id is required")
	}
	//checkIntime := time.Unix(eventCheckIn.CheckInTime, 0).Format("2006-01-02 15:04:05")

	res, err := s.eventRepo.CheckInEvent(eventCheckIn.UserId, &repository.EventCheckIn{
		EventId:      eventCheckIn.EventId,
		UserId:       eventCheckIn.UserId,
		CheckIn:      eventCheckIn.CheckIn,
		CheckInTime:  eventCheckIn.CheckInTime,
		CheckInPlace: eventCheckIn.CheckInPlace,
		CheckOut:     eventCheckIn.CheckOut,
		CheckOutTime: eventCheckIn.CheckOutTime,
		Clinic:       eventCheckIn.Clinic,
	})
	if err != nil {
		return false, err
	}
	return res, nil

}
func (s *lineBotService) MyEvents(userId string) ([]*EventResponse, error) {
	if userId == "" {
		return nil, errors.New("user id is required")
	}

	myEvents, err := s.eventRepo.EventByUserId(userId)
	if err != nil {
		return nil, err
	}
	eventResponse := make([]*EventResponse, 0, len(myEvents))
	for _, event := range myEvents {
		banners := make([]EventBanner, 0, len(event.Banner))
		if event.Banner != nil {
			for _, v := range event.Banner {
				banners = append(banners, EventBanner{
					Url: v.Url,
					Img: v.Img,
				})
			}
		}

		// Unix time to Datetime string format
		stDate := time.Unix(event.StartDate, 0).Format("2006-01-02")
		enDate := time.Unix(event.EndDate, 0).Format("2006-01-02")
		stTime := time.Unix(event.StartTime, 0).Format("15:04")
		enTime := time.Unix(event.EndTime, 0).Format("15:04")

		memberJoinEvent := &EventResponse{
			EventId:     event.EventId,
			Title:       event.Title,
			Description: event.Description,
			StartDate:   stDate,
			StartTime:   stTime,
			EndDate:     enDate,
			EndTime:     enTime,
			Place:       event.Place,
			Banner:      banners,
			Location:    event.Location,
			Status:      event.Status,
		}
		eventResponse = append(eventResponse, memberJoinEvent)
	}
	return eventResponse, nil

	//memberJoinEvents := make([]*MemberJoinEvent, len(myEvents))
	//for i, v := range myEvents {
	//	joinTimeStr := time.Unix(v.JoinTime, 0).Format("2006-01-02 15:04:05")
	//	memberJoinEvents[i] = &MemberJoinEvent{
	//		EventId:      v.EventId,
	//		UserId:       v.UserId,
	//		JoinTime:     joinTimeStr,
	//		Name:         v.Name,
	//		LastName:     v.LastName,
	//		Organization: v.Organization,
	//	}
	//}
	//return memberJoinEvents, nil
}
func (s *lineBotService) ReportFlexCarouselMessage(replyToken string) error {

	//baseUrl := "https://dca3a8ac633b.ngrok.app"
	enventImg := "https://hostpital-sd.s3.ap-southeast-7.amazonaws.com/events-report.png"
	memberImg := "https://hostpital-sd.s3.ap-southeast-7.amazonaws.com/members-report.png"
	eventReportLink := "https://dca3a8ac633b.ngrok.app/api/v1/reports/events/excel"
	memberReportLink := "https://dca3a8ac633b.ngrok.app/api/v1/reports/events/excel"
	//
	//reportFlexCarouselMessage := messaging_api.CarouselTemplate{
	//	Columns: []messaging_api.CarouselColumn{
	//		{ThumbnailImageUrl: enventImg,
	//			Title: "Events Report",
	//			Text:  "Download Events Report",
	//			DefaultAction: &messaging_api.UriAction{
	//				Label: "Download",
	//				Uri:   enventReportLink,
	//			},
	//			Actions: []messaging_api.ActionInterface{
	//				messaging_api.UriAction{
	//					Label: "Download",
	//					Uri:   enventReportLink,
	//				},
	//			}}, {
	//			ThumbnailImageUrl: memberImg,
	//			Title:             "Member Report",
	//			Text:              "Download Member Report",
	//			DefaultAction: &messaging_api.UriAction{
	//				Label: "Download",
	//				Uri:   memberReportLink,
	//			},
	//			Actions: []messaging_api.ActionInterface{
	//				messaging_api.UriAction{
	//					Label: "Download",
	//					Uri:   memberReportLink,
	//				},
	//			},
	//		},
	//	},
	//}

	reportFlexCarouselMessage := messaging_api.FlexMessage{}
	reportFlexCarouselMessage.AltText = "Report"
	reportFlexCarouselMessage.Contents = messaging_api.FlexCarousel{
		Contents: []messaging_api.FlexBubble{
			{
				Body: &messaging_api.FlexBox{
					Layout: messaging_api.FlexBoxLAYOUT_VERTICAL,
					Contents: []messaging_api.FlexComponentInterface{
						&messaging_api.FlexImage{
							Url:         memberImg,
							Size:        "full",
							AspectRatio: "1.9:1",
							Action: &messaging_api.UriAction{
								Uri:   memberReportLink,
								Label: "Member Report download",
							},
						},
						&messaging_api.FlexImage{
							Url:         enventImg,
							Size:        "full",
							AspectRatio: "1.9:1",
							Action: &messaging_api.UriAction{
								Uri:   eventReportLink,
								Label: "Events Report download",
							},
						},
						&messaging_api.FlexImage{
							Url:         "https://www.linefriends.com/img/img_sec.jpg",
							Size:        "full",
							AspectRatio: "1.9:1",
							Action: &messaging_api.UriAction{
								Uri:   "https://www.linefriends.com",
								Label: "https://www.linefriends.com",
							},
						},
					},
				},
			},
		},
	}

	_, err := s.bot.ReplyMessage(&messaging_api.ReplyMessageRequest{
		ReplyToken: replyToken,
		Messages:   []messaging_api.MessageInterface{&reportFlexCarouselMessage},
	})
	if err != nil {
		log.Println("error")
		log.Println(err)
		return err
	}

	return nil
}
func validation(member *Member) error {
	if member.Name == "" {
		return errors.New("name is required")
	}
	if member.Email == "" {
		return errors.New("email is required")
	}
	if member.LineId == "" {
		return errors.New("lineId is required")
	}
	return nil
}
