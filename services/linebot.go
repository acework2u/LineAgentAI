package services

type LineBotService interface {
	SendTextMessage(text string) error
	ReplyMessage(replyToken string, text string) error
}

type SourceHook struct {
	Type   string `json:"type"`
	Source struct {
		Type   string `json:"type"`
		UserID string `json:"userId"`
	} `json:"source"`
}
type LineEventResponse struct {
	Destination string `json:"destination"`
	Events      []struct {
		ReplyToken string `json:"replyToken"`
		Type       string `json:"type"`
		Mode       string `json:"mode"`
		Timestamp  int64  `json:"timestamp"`
		Source     struct {
			Type    string `json:"type"`
			GroupID string `json:"groupId"`
			UserID  string `json:"userId"`
		} `json:"source"`
		WebhookEventID  string `json:"webhookEventId"`
		DeliveryContext struct {
			IsRedelivery bool `json:"isRedelivery"`
		} `json:"deliveryContext"`
		Message struct {
			ID         string `json:"id"`
			Type       string `json:"type"`
			QuoteToken string `json:"quoteToken"`
			Text       string `json:"text"`
			Emojis     []struct {
				Index     int    `json:"index"`
				Length    int    `json:"length"`
				ProductID string `json:"productId"`
				EmojiID   string `json:"emojiId"`
			} `json:"emojis"`
			Mention struct {
				Mentionees []struct {
					Index  int    `json:"index"`
					Length int    `json:"length"`
					Type   string `json:"type"`
					UserID string `json:"userId,omitempty"`
					IsSelf bool   `json:"isSelf,omitempty"`
				} `json:"mentionees"`
			} `json:"mention"`
		} `json:"message"`
	} `json:"events"`
}
