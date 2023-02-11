package dingtalk

const (
	MsgTypeText       MsgType = "text"
	MsgTypeLink       MsgType = "link"
	MsgTypeMarkdown   MsgType = "markdown"
	MsgTypeActionCard MsgType = "actionCard"
)

// MsgType message type
type MsgType string

//
type Message interface {
	GetMsgType() MsgType
}

type MessageText struct {
	// Content message content
	Content string `json:"content"`
}

// TextMessage
type TextMessage struct {
	MsgType MsgType `json:"msgtype"`
	Text    Text    `json:"content"`
	At      At      `json:"at"`
}

type Text struct {
	// Content message content
	Content string `json:"content"`
}

type At struct {
	// AtMobiles list of mobile phone numbers of people to be reminded
	AtMobiles []string `json:"atMobiles,omitempty"`
	// AtUserIds list of user ids of people to be reminded
	AtUserIds []string `json:"atUserIds,omitempty"`
	// IsAtAll
	IsAtAll bool `json:"isAtAll,omitempty"`
}

func (m TextMessage) GetMsgType() MsgType {
	return MsgTypeText
}

type LinkMessage struct {
	MsgType MsgType `json:"msgtype"`
	Link    Link    `json:"link"`
}

type Link struct {
	// Text message content
	Text string `json:"text"`
	// Title message title
	Title string `json:"title"`
	// PicUrl picture url
	PicUrl string `json:"picUrl"`
	// MessageUrl jump url of the click message
	MessageUrl string `json:"messageUrl"`
}

func (m LinkMessage) GetMsgType() MsgType {
	return MsgTypeLink
}

type MarkdownMessage struct {
	MsgType  MsgType  `json:"msgtype"`
	Markdown Markdown `json:"link"`
	At       At       `json:"at"`
}

type Markdown struct {
	// Title the display content revealed in the first screen conversation
	Title string `json:"title"`
	// Text message content that format is markdown
	Text string `json:"text"`
}

func (m MarkdownMessage) GetMsgType() MsgType {
	return MsgTypeMarkdown
}

type ActionCardMessage struct {
	MsgType    MsgType    `json:"msgtype"`
	ActionCard ActionCard `json:"actionCard"`
}

type ActionCard struct {
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
	Btns []ActionCardBtn `json:"btns,omitempty"`
}

type ActionCardBtn struct {
	// Title title of button
	Title string `json:"title"`
	// ActionURL jump url of the click message
	ActionURL string `json:"actionURL"`
}

func (m ActionCardMessage) GetMsgType() MsgType {
	return MsgTypeActionCard
}
