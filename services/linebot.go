package services

type LineBotService interface {
	SendTextMessage(text string) error
	ReplyMessage(replyToken string, text string) error
	SendFlexMessage(replyToken string) error
	SendImageMessage(replyToken string, imageUrl string) error
	SendFlexCarouselMessage(replyToken string) error
	SendFlexJsonMessage(replyToken string, json string) error
	SendQuickReplyMessage(replyToken string) error
	SendDateTimePickerMessage(replyToken string) error
	RegisterMember(member *Member) error
	GetLineProfile(userId string) (*UserInfo, error)
	UpdateMemberProfile(userId string, member *Member) error
	CheckMemberRegister(userId string) (bool, error)
	EventJoin(event *MemberJoinEvent) error
	CheckEventJoin(eventId string, userId string) (bool, error)
	GetEventJoin(eventId string, userId string) (*MemberJoinEvent, error)
	CheckInEvent(eventCheckIn *EventCheckIn) (bool, error)
}
type SourceHook struct {
	Type   string `json:"type"`
	Source struct {
		Type   string `json:"type"`
		UserID string `json:"userId"`
	} `json:"source"`
}
type Member struct {
	Title        string `json:"title"`
	Name         string `json:"name" binding:"required" validate:"required,min=3,max=20"`
	LastName     string `json:"lastName" binding:"required" validate:"required,min=3,max=20"`
	PinCode      int    `json:"pinCode"`
	Email        string `json:"email" binding:"required" validate:"required,email"`
	Phone        string `json:"phone" binding:"required" validate:"required,numeric,min=10,max=10"`
	BirthDate    int64  `json:"birthDate" `
	Med          string `json:"med" binding:"required"`
	MedExtraInfo string `json:"medExtraInfo"`
	Organization string `json:"organization" binding:"required"`
	Position     string `json:"position" binding:"required"`
	Course       string `json:"course" binding:"required"`
	LineId       string `json:"lineId" binding:"required"`
	LineName     string `json:"lineName"`
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

type FlexMessageTemplate struct {
	Type     string       `json:"type"`
	AltText  string       `json:"altText"`
	Contents FlexContents `json:"contents"`
}
type FlexContents struct {
	Type string       `json:"type"`
	Body LineFlexBody `json:"body"`
}

type LineFlexBody struct {
	Type     string                `json:"type"`
	Layout   string                `json:"layout"`
	Contents []LineFlexBodyContent `json:"contents"`
}
type LineFlexBodyContent struct {
	Type string `json:"type,omitempty"`
	Text string `json:"text,omitempty"`
}
type FlexMessageResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
type FlexMessageRequest struct {
	To       string                `json:"to"`
	Messages []FlexMessageTemplate `json:"messages"`
}
type ImageMessageTemplate struct {
	Type        string `json:"type"`
	URL         string `json:"url"`
	Size        string `json:"size"`
	AspectRatio string `json:"aspectRatio"`
}
type ReplyMessage struct {
	ReplyToken string `json:"replyToken"`
	Messages   []Text `json:"messages"`
}
type Text struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type MemberJoinEvent struct {
	EventId        string `json:"eventId" binding:"required" validate:"required,min=3,max=20"`
	UserId         string `json:"userId" binding:"required"`
	JoinTime       string `json:"joinTime,omitempty"`
	Name           string `json:"name" binding:"required" validate:"required,min=3,max=20"`
	LastName       string `json:"lastName" binding:"required" validate:"required,min=3,max=20"`
	Organization   string `json:"organization" binding:"required"`
	Position       string `json:"position" binding:"required"`
	Course         string `json:"course" binding:"required"`
	LineId         string `json:"lineId" binding:"required"`
	LineName       string `json:"lineName" binding:"required"`
	Tel            string `json:"tel" binding:"required" validate:"required,numeric,min=10,max=10"`
	ReferenceName  string `json:"referenceName" binding:"required" validate:"required,min=3,max=20"`
	ReferencePhone string `json:"referencePhone" binding:"required" validate:"required,numeric,min=10,max=10"`
	Clinic         string `json:"clinic" binding:"required"`
}

type QrCodeMessage struct {
	QrCode    string `json:"qrCode"`
	Timestamp string `json:"timestamp"`
	UserId    string `json:"userId"`
}

type QrCodeImpl struct {
	ClinicNo string `json:"clinicNo"`
	EventId  string `json:"eventId"`
}
type EventCheckIn struct {
	EventId      string `json:"eventId"`
	UserId       string `json:"userId"`
	CheckIn      bool   `json:"checkIn"`
	CheckOut     bool   `json:"checkOut"`
	CheckInTime  int64  `json:"checkInTime,omitempty"`
	CheckOutTime int64  `json:"checkOutTime,omitempty"`
	CheckInPlace string `json:"checkInPlace"`
}
