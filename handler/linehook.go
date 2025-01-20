package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
	"linechat/conf"
	"linechat/services"
	"log"
	"net/http"
)

type LineWebhookHandler struct {
	cfg         *conf.AppConfig
	linebot     *messaging_api.MessagingApiAPI
	lineService services.LineBotService
}

func NewLineWebhookHandler(lineService services.LineBotService) *LineWebhookHandler {

	cfg, _ := conf.NewAppConfig()
	bot, err := messaging_api.NewMessagingApiAPI(cfg.LineApp.ChannelToken)
	if err != nil {
		log.Println(err)
	}

	return &LineWebhookHandler{
		cfg:         cfg,
		linebot:     bot,
		lineService: lineService,
	}
}

func (h *LineWebhookHandler) LineHookHandle(c *gin.Context) {
	cb, err := webhook.ParseRequest(h.cfg.LineApp.ChannelSecret, c.Request)
	if err != nil {
		if err == webhook.ErrInvalidSignature {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid signature"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		}
		return
	}

	for _, event := range cb.Events {

		switch e := event.(type) {
		case webhook.MessageEvent:
			// Do Something...

			data, _ := json.Marshal(e.Source)
			eventSource := make(map[string]interface{})
			_ = json.Unmarshal(data, &eventSource)
			userId := eventSource["userId"].(string)

			log.Printf("MessageEv: %v", e.Message)

			h.lineService.ReplyMessage(e.ReplyToken, "Hello Doctor"+" "+userId)
			// 	roles, ok := jwtClaims["payload"].(map[string]interface{})["acl"].([]interface{})

			log.Printf("MessageEvent: %v", eventSource["userId"])
			//replyToken := e.ReplyToken

		case webhook.StickerMessageContent:
			// Do Something...
		}
	}

	c.JSON(200, gin.H{"message": "OK"})
}
