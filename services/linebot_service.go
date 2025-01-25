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
	imageMessage := messaging_api.TemplateMessage{
		AltText: "Image carousel alt text",
		Template: &messaging_api.ImageCarouselTemplate{
			Columns: []messaging_api.ImageCarouselColumn{
				{
					ImageUrl: imageURL,
					Action: messaging_api.UriAction{
						Label: "View detail",
						Uri:   "https://dca3a8ac633b.ngrok.app/attend",
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
			},
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

	member.Status = true
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
