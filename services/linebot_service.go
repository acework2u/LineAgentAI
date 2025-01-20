package services

import (
	"fmt"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"linechat/conf"
)

type lineBotService struct {
	cfg *conf.AppConfig
	bot *messaging_api.MessagingApiAPI
}

func NewLineBotService() LineBotService {
	cfg, _ := conf.NewAppConfig()
	bot, _ := messaging_api.NewMessagingApiAPI(cfg.LineApp.ChannelToken)

	return &lineBotService{bot: bot}
}
func (s *lineBotService) SendTextMessage(text string) error {
	return nil
}
func (s *lineBotService) ReplyMessage(replyToken string, meg string) error {

	replayTxt := fmt.Sprintf("%s", meg)
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
