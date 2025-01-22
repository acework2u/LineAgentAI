package services

type LineBotService interface {
	SendTextMessage(text string) error
	ReplyMessage(replyToken string, text string) error
	RegisterMember(member *Member) error
	GetLineProfile(userId string) (*UserInfo, error)
	UpdateMemberProfile(userId string, member *Member) error
	CheckMemberRegister(userId string) (bool, error)
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
type Member struct {
	Name         string `json:"name" binding:"required" validate:"required,min=3,max=20"`
	LastName     string `json:"lastName" binding:"required" validate:"required,min=3,max=20"`
	PinCode      int    `json:"pinCode"`
	Email        string `json:"email" binding:"required" validate:"required,email"`
	Phone        string `json:"phone" binding:"required" validate:"required,numeric,min=10,max=10"`
	BirthDate    int64  `json:"birthDate" `
	Med          string `json:"med" binding:"required"`
	Organization string `json:"organization" binding:"required"`
	Position     string `json:"position" binding:"required"`
	Course       string `json:"course" binding:"required"`
	LineId       string `json:"lineId" binding:"required"`
	Line         string `json:"line"`
	Facebook     string `json:"facebook"`
	Instagram    string `json:"instagram"`
	FoodAllergy  string `json:"foodAllergy"`
	Religion     string `json:"religion"`
	RegisterDate int64  `json:"registerDate,omitempty"`
	UpdatedDate  int64  `json:"updatedDate,omitempty"`
	Status       bool   `json:"status"`
}
type UserInfo struct {
	UserID     string `json:"userId"`
	Name       string `json:"name"`
	LastName   string `json:"lastName"`
	PictureURL string `json:"pictureUrl"`
	LineID     string `json:"lineId"`
	Status     bool   `json:"status"`
	Member     Member `json:"member"`
}

type UserMinInfo struct {
	UserID string `json:"userId"`
}
