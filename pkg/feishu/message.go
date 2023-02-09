package feishu

const (
	MsgTypeText        MessageType = "text"
	MsgTypeInteractive MessageType = "interactive"
)

type MessageType string

type Message interface {
	GetMsgType() MessageType
}

//
type TextMessage struct {
	MsgType    MessageType        `json:"msg_type"`
	MsgContent TextMessageContent `json:"content"`
}

//
type TextMessageContent struct {
	Text string `json:"text"`
}

//
func (m TextMessage) GetMsgType() MessageType {
	return MsgTypeText
}

// IntractiveMessage
type InteractiveMessage struct {
	MsgType MessageType            `json:"msg_type"`
	MsgCard InteractiveMessageCard `json:"card"`
}

func (m InteractiveMessage) GetMsgType() MessageType {
	return MsgTypeInteractive
}

type InteractiveMessageCard struct {
	Header   InteractiveMessageCardHeader    `json:"header"`
	Elements []InteractiveMessageCardElement `json:"elements"`
}

type InteractiveMessageCardHeader struct {
	Template string                            `json:"template"`
	Title    InteractiveMessageCardHeaderTitle `json:"title"`
}

type InteractiveMessageCardHeaderTitle struct {
	Content string `json:"content"`
	Tag     string `json:"tag"`
}

type InteractiveMessageCardElement struct {
	Tag     string                                `json:"tag"`
	Text    InteractiveMessageCardElementText     `json:"text,omitempty"`
	Content string                                `json:"content,omitempty"`
	Actions []InteractiveMessageCardElementAction `json:"actions,omitempty"`
}

type InteractiveMessageCardElementText struct {
	Content string `json:"content"`
	Tag     string `json:"tag"`
}

type InteractiveMessageCardElementAction struct {
	Tag   string                            `json:"tag"`
	Text  InteractiveMessageCardElementText `json:"text"`
	Url   string                            `json:"url"`
	Type  string                            `json:"type"`
	Value map[string]string                 `json:"value"`
}
