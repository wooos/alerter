package feishu

const (
	MsgTypeText        MsgType = "text"
	MsgTypeInteractive MsgType = "interactive"
)

type MsgType string

type Message interface {
	GetMsgType() MsgType
}

type TextMessage struct {
	MsgType MsgType `json:"msg_type"`
	Content Content `json:"content"`
}

type Content struct {
	// Text message content
	Text string `json:"text"`
}

func (m TextMessage) GetMsgType() MsgType {
	return MsgTypeText
}

// IntractiveMessage
type InteractiveMessage struct {
	MsgType MsgType `json:"msg_type"`
	Card    Card    `json:"card"`
}

type Card struct {
	Header   CardHeader    `json:"header"`
	Elements []CardElement `json:"elements"`
}

type CardHeader struct {
	Template string          `json:"template"`
	Title    CardHeaderTitle `json:"title"`
}

type CardHeaderTitle struct {
	Content string `json:"content"`
	Tag     string `json:"tag"`
}

type CardElement struct {
	Tag     string              `json:"tag"`
	Text    CardElementText     `json:"text,omitempty"`
	Content string              `json:"content,omitempty"`
	Actions []CardElementAction `json:"actions,omitempty"`
}

type CardElementText struct {
	Content string `json:"content"`
	Tag     string `json:"tag"`
}

type CardElementAction struct {
	Tag   string            `json:"tag"`
	Text  CardElementText   `json:"text"`
	Url   string            `json:"url"`
	Type  string            `json:"type"`
	Value map[string]string `json:"value"`
}

func (m InteractiveMessage) GetMsgType() MsgType {
	return MsgTypeInteractive
}
