package dingtalk

// MessageType message type
type MessageType string

//
type Message interface {
	GetMsgType() MessageType
}

// MessageAt
type MessageAt struct {
	// AtMobiles list of mobile phone numbers of people to be reminded
	AtMobiles []string `json:"atMobiles"`
	// AtUserIds list of user ids of people to be reminded
	AtUserIds []string `json:"atUserIds"`
	// IsAtAll
	IsAtAll bool `json:"isAtAll"`
}

type MessageText struct {
	// Content message content
	Content string `json:"content"`
}

type DingtalkMessageText struct {
	MsgType MessageType `json:"msgtype"`
	At      MessageAt   `json:"at"`
	Text    MessageText `json:"text"`
}

func (m DingtalkMessageText) GetMsgType() MessageType {
	return "text"
}

type MessageLink struct {
	// Text message content
	Text string `json:"text"`
	// Title message title
	Title string `json:"title"`
	// PicUrl picture url
	PicUrl string `json:"picUrl"`
	// MessageUrl jump url of the click message
	MessageUrl string `json:"messageUrl"`
}

type DingtalkMessageLink struct {
	MsgType MessageType `json:"msgtype"`
	Link    MessageLink `json:"link"`
}

func (m DingtalkMessageLink) GetMsgType() MessageType {
	return "link"
}

type MessageMarkdown struct {
	// Title the display content revealed in the first screen conversation
	Title string `json:"title"`
	// Text message content that format is markdown
	Text string `json:"text"`
}

type DingtalkMessageMarkdown struct {
	MsgType  MessageType     `json:"msgtype"`
	Markdown MessageMarkdown `json:"markdown"`
	At       MessageAt       `json:"at"`
}

func (m DingtalkMessageMarkdown) GetMsgType() MessageType {
	return "markdown"
}

type MessageActionCardBtn struct {
	// Title title of button
	Title string `json:"title"`
	// ActionURL jump url of the click message
	ActionURL string `json:"actionURL"`
}

type MessageActionCard struct {
	// Title the display content revealed in the first screen conversation
	Title string `json:"title"`
	// Text message content that format is markdown
	Text string `json:"text"`
	// BtnOrientation button orientation, 0 is vertical, 1 is horizontal
	BtnOrientation string `json:"btnOrientation"`
	// SingleTitle title of a single button, if set and set singleURL, btns is invalid
	SingleTitle string `json:"singleTitle"`
	// SingleURL jump url of the click message
	SingleURL string `json:"singleURL"`
	// Btns buttons array
	Btns []MessageActionCardBtn `json:"btns"`
}

type DingtalkMessageActionCard struct {
	MsgType    MessageType       `json:"msgtype"`
	ActionCard MessageActionCard `json:"actionCard"`
}

// NewDingtalkMessageText return dingtalk message that type is text
func NewDingtalkMessageText(message string, atAll bool, mobiles, uids []string) DingtalkMessageText {
	return DingtalkMessageText{
		MsgType: "text",
		Text: MessageText{
			Content: message,
		},
		At: MessageAt{
			IsAtAll:   atAll,
			AtMobiles: mobiles,
			AtUserIds: uids,
		},
	}
}

// NewDingtalkMessageLink return dingtalk message that type is link
func NewDingtalkMessageLink(title, message, messageUrl, picUrl string) DingtalkMessageLink {
	return DingtalkMessageLink{
		MsgType: "link",
		Link: MessageLink{
			Text:       message,
			Title:      title,
			MessageUrl: messageUrl,
			PicUrl:     picUrl,
		},
	}
}

// NewDingtalkMessageMarkdown return dingtalk message that type is markdown
func NewDingtalkMessageMarkdown(title, message string, atAll bool, mobiles, uids []string) DingtalkMessageMarkdown {
	return DingtalkMessageMarkdown{
		MsgType: "markdown",
		Markdown: MessageMarkdown{
			Title: title,
			Text:  message,
		},
		At: MessageAt{
			IsAtAll:   atAll,
			AtMobiles: mobiles,
			AtUserIds: uids,
		},
	}
}

// NewDingtalkMessageActionCard return dingtalk message that type is actionCard
func NewDingtalkMessageActionCard(title, message, btnOrientation, singleTitle, singleURL string) DingtalkMessageActionCard {
	return DingtalkMessageActionCard{
		MsgType: "actionCard",
		ActionCard: MessageActionCard{
			Title:          title,
			Text:           message,
			BtnOrientation: btnOrientation,
			SingleTitle:    singleTitle,
			SingleURL:      singleURL,
		},
	}
}
