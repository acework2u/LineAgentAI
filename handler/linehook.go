package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
	"linechat/conf"
	"linechat/services"
	"linechat/utils"
	"log"
	"net/http"
	"net/url"
	"time"
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
			eSource := e.Source.(webhook.UserSource)
			log.Printf("SourceMsg: %v", eSource.UserId)

			//data, _ := json.Marshal(e.Source)
			//eventSource := make(map[string]interface{})
			//_ = json.Unmarshal(data, &eventSource)
			//userId := eventSource["userId"].(string)

			log.Printf("MessageEv: %v", e.Message)
			msg2 := e.Message.(webhook.TextMessageContent)
			log.Printf("MessageWebHook: %v", msg2.Text)

			rawMsg, _ := json.Marshal(e.Message)
			msg := make(map[string]interface{})
			_ = json.Unmarshal(rawMsg, &msg)
			if msg["type"].(string) == "text" {
				switch msg["text"].(string) {
				case "medical volunteer":
				case "ร่วมงานแพทย์อาสา":

					err := h.lineService.SendFlexMessage(e.ReplyToken)
					if err != nil {
						log.Println(err)
					}

					//HelloTxt := fmt.Sprintf("Hi %v", eSource.UserId)
					//h.lineService.ReplyMessage(e.ReplyToken, HelloTxt)
				case "news":
				case "ประชาสัมพันธ์":
					//imgUrl := "https://drive.google.com/file/d/1Zz0q8lf6fvNCQoNzZKkDw8OMaND95-Yp/view?usp=sharing"
					imgUrl := "https://hostpital-sd.s3.ap-southeast-7.amazonaws.com/S__29261992.png"
					h.lineService.SendImageMessage(e.ReplyToken, imgUrl)
				case "calendar":
				case "ปฏิทินงาน":
					//h.lineService.SendDateTimePickerMessage(e.ReplyToken)
					//h.lineService.SendFlexCarouselMessage(e.ReplyToken)
					//h.lineService.SendFlexJsonMessage(e.ReplyToken, "")
				case "contact us":
				case "ติดต่อเรา":
					//h.lineService.SendFlexJsonMessage(e.ReplyToken, "")
					//h.lineService.SendQuickReplyMessage(e.ReplyToken)

				case "custom":
					log.Printf("custom: %v", msg["text"])
					h.lineService.SendFlexJsonMessage(e.ReplyToken, "")

				default:
					break
					// no action
				}
			}

			//replyToken := e.ReplyToken

		case webhook.StickerMessageContent:
			// Do Something...
		case webhook.TextMessageContent:
			log.Printf("TextMessageContent: %v", e.Text)
			// Do something
		case webhook.PostbackEvent:
			postbackData := e.Postback.Data
			log.Printf("Postback Data: %s", postbackData)

			if postbackData == "DATE" || postbackData == "TIME" || postbackData == "DATETIME" {
				// Check if there are parameters (e.g., datetime picker)
				if e.Postback.Params != nil {
					log.Printf("Postback Params: %v", e.Postback.Params)
					data := e.Postback.Params
					date := data["date"]
					time := data["time"]
					datetime := data["datetime"]
					if date != "" {
						log.Printf("Selected Date: %s", date)
					}
					if time != "" {
						log.Printf("Selected Time: %s", time)
					}
					if datetime != "" {
						log.Printf("Selected DateTime: %s", datetime)
					}
				}
			}

			// Parse the query string
			values, err := url.ParseQuery(postbackData)
			if err != nil {
				log.Printf("Error parsing postback data: %v", err)
				return
			}
			// Extract individual values
			action := values.Get("action")
			itemID := values.Get("itemid")

			log.Printf("Action: %s, ItemID: %s", action, itemID)

			// Reply to the user based on extracted values
			replyMessage := "Action: " + action + ", Item ID: " + itemID
			h.lineService.ReplyMessage(e.ReplyToken, replyMessage)

		} // end of switch

	}

	c.JSON(200, gin.H{"message": "OK"})
}
func (h *LineWebhookHandler) LineCallback(c *gin.Context) {
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
		log.Printf("Event: %v", event)
	}

	c.JSON(200, gin.H{"message": "OK"})
}
func (h *LineWebhookHandler) LineRegister(c *gin.Context) {
	member := &services.Member{}
	err := c.ShouldBindJSON(member)
	cusErr := utils.NewCustomErrorHandler(c)

	if err != nil {

		cusErr.ValidateError(err)
		//c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err = h.lineService.RegisterMember(member)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	//c.JSON(200, gin.H{"member register a message": member})
	c.JSON(200, gin.H{"message": "Member registration successful"})
}
func (h *LineWebhookHandler) LineLogin(c *gin.Context) {
	c.JSON(200, gin.H{"member login a message": "OK"})
}
func (h *LineWebhookHandler) LineLogout(c *gin.Context) {
	c.JSON(200, gin.H{"member logout a message": "OK"})
}
func (h *LineWebhookHandler) LineChat(c *gin.Context) {
	c.JSON(200, gin.H{"member chat a message": "OK"})
}
func (h *LineWebhookHandler) GetLineProfile(c *gin.Context) {

	userId := c.Query("userId")
	if userId == "" {
		c.JSON(400, gin.H{"error": "userId is required"})
		return
	}
	userInfo, err := h.lineService.GetLineProfile(userId)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"userInfo": userInfo})

}
func (h *LineWebhookHandler) PutLineProfile(c *gin.Context) {
	member := &services.Member{}
	err := c.ShouldBindJSON(member)

	cusErr := utils.NewCustomErrorHandler(c)
	if err != nil {
		cusErr.ValidateError(err)
		return
	}
	err = h.lineService.UpdateMemberProfile(member.LineId, member)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	c.JSON(200, gin.H{"message": "Member profile updated"})
}
func (h *LineWebhookHandler) CheckMemberRegister(c *gin.Context) {
	//userId := c.Query("userId")
	var userId string
	userMiniInfo := services.UserMinInfo{}
	err := c.ShouldBindJSON(&userMiniInfo)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	userId = userMiniInfo.UserID

	if userId == "" {
		c.JSON(400, gin.H{"error": "userId is required"})
		return
	}
	isRegister, err := h.lineService.CheckMemberRegister(userId)
	if err != nil {

		if err.Error() == "mongo: no documents in result" {
			c.JSON(200, gin.H{"isRegistered": false})
			return
		}

		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"isRegistered": isRegister})

}
func (h *LineWebhookHandler) PostUpdateMember(c *gin.Context) {
	member := &services.Member{}
	err := c.ShouldBindJSON(member)
	cusErr := utils.NewCustomErrorHandler(c)
	if err != nil {
		cusErr.ValidateError(err)
		return

	}
	err = h.lineService.UpdateMemberProfile(member.LineId, member)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	log.Println(member)
	c.JSON(200, gin.H{"message": "update member successful"})
}
func (h *LineWebhookHandler) PostJoinEvent(c *gin.Context) {

	eventJoin := services.MemberJoinEvent{}
	err := c.ShouldBindJSON(&eventJoin)
	//log.Println("eventJoin:")
	//log.Println(eventJoin)
	cusErr := utils.NewCustomErrorHandler(c)
	if err != nil {
		log.Println("In Join Event: error")
		log.Println(err)
		cusErr.ValidateError(err)
		return
	}

	err = h.lineService.EventJoin(&eventJoin)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "join event successful"})

}
func (h *LineWebhookHandler) GetCheckEventJoin(c *gin.Context) {
	userId := c.Query("userId")
	eventID := c.Query("eventId")
	if userId == "" {
		c.JSON(400, gin.H{"error": "userId is required"})
		return
	}
	if eventID == "" {
		c.JSON(400, gin.H{"error": "eventId is required"})
		return
	}
	eventJoin, err := h.lineService.CheckEventJoin(eventID, userId)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"eventJoin": eventJoin})

}
func (h *LineWebhookHandler) GetEventJoin(c *gin.Context) {
	userId := c.Query("userId")
	eventID := c.Query("eventId")
	if userId == "" {
		c.JSON(400, gin.H{"error": "userId is required"})
		return
	}
	if eventID == "" {
		c.JSON(400, gin.H{"error": "eventId is required"})
		return
	}
	eventJoin, err := h.lineService.GetEventJoin(eventID, userId)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": eventJoin})
}
func (h *LineWebhookHandler) PostCheckInEvent(c *gin.Context) {
	qrcode := services.QrCodeMessage{}
	err := c.ShouldBindJSON(&qrcode)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	log.Println("In CheckInEvent: ")
	log.Println(qrcode)

	dataQr := services.QrCodeImpl{}
	err = json.Unmarshal([]byte(qrcode.QrCode), &dataQr)

	if err != nil {
		log.Println("In CheckInEvent: error")
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	checkInData := services.QrCodeMessage{}
	_ = checkInData

	dataCheckIn := services.EventCheckIn{
		EventId:      dataQr.EventId,
		UserId:       qrcode.UserId,
		CheckIn:      true,
		CheckOut:     false,
		CheckInTime:  conTimeStrToInt64(qrcode.Timestamp),
		CheckOutTime: 0,
		CheckInPlace: dataQr.ClinicNo,
	}

	log.Println(dataCheckIn)

	resCheckIn, err := h.lineService.CheckInEvent(&dataCheckIn)
	if err != nil {
		log.Println("In CheckInEvent: error")
		log.Println(err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	_ = resCheckIn

	c.JSON(200, gin.H{"isCheckIn": true})
}
func (h *LineWebhookHandler) PostMyEvents(c *gin.Context) {

	userId := services.EventCheckIn{}
	err := c.ShouldBindJSON(&userId)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if len(userId.UserId) == 0 {
		c.JSON(400, gin.H{"error": "userId is required"})
		return
	}

	myEvents, err := h.lineService.MyEvents(userId.UserId)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"myEvents": myEvents})

}
func conTimeStrToInt64(checkInTimeStr string) int64 {
	// Assuming the check-in time is provided as a string in a specific datetime format.
	const layout = "2006-01-02T15:04:05Z" // Example datetime format (adjust to your needs)

	checkInTime, err := time.Parse(layout, checkInTimeStr)
	if err != nil {
		log.Println(err.Error())
	}
	return checkInTime.Unix()

}
