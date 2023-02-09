package message

import (
	"fmt"

	"github.com/wooos/alerter/internal/config"
	"github.com/wooos/alerter/internal/pkg/request"
	"github.com/wooos/alerter/pkg/feishu"
	"github.com/wooos/alerter/pkg/utils"
)

var (
	IgnoreLabelKeys = []string{
		"alertname",
		"endpoint",
		"uid",
	}
)

func SendMessage(alert request.AlertRequestAlert) {
	if config.Conf.Feishu.Enabled {
		msg := feishu.InteractiveMessage{
			MsgType: feishu.MsgTypeInteractive,
		}

		msg.MsgCard.Elements = append(msg.MsgCard.Elements, feishu.InteractiveMessageCardElement{
			Tag:     "markdown",
			Content: "**Labels**",
		})

		for k, v := range alert.Labels {
			if utils.StrInArray(k, IgnoreLabelKeys) {
				continue
			}

			msg.MsgCard.Elements = append(msg.MsgCard.Elements, feishu.InteractiveMessageCardElement{
				Tag:     "markdown",
				Content: fmt.Sprintf("  - %s: %s", k, v),
			})
		}

		msg.MsgCard.Elements = append(msg.MsgCard.Elements, feishu.InteractiveMessageCardElement{
			Tag:     "markdown",
			Content: "**Annotations**",
		})

		for k, v := range alert.Annotations {
			msg.MsgCard.Elements = append(msg.MsgCard.Elements, feishu.InteractiveMessageCardElement{
				Tag:     "markdown",
				Content: fmt.Sprintf("  - %s: %s", k, v),
			})
		}

		msg.MsgCard.Elements = append(msg.MsgCard.Elements, feishu.InteractiveMessageCardElement{
			Tag: "hr",
		})

		switch alert.Status {
		case "firing":
			msg.MsgCard.Header = feishu.InteractiveMessageCardHeader{
				Template: "red",
				Title: feishu.InteractiveMessageCardHeaderTitle{
					Content: fmt.Sprintf("[告警触发] %s", alert.Labels["alertname"]),
					Tag:     "plain_text",
				},
			}

			msg.MsgCard.Elements = append(msg.MsgCard.Elements, feishu.InteractiveMessageCardElement{
				Tag: "action",
				Actions: []feishu.InteractiveMessageCardElementAction{
					{
						Tag:  "button",
						Type: "danger",
						Text: feishu.InteractiveMessageCardElementText{
							Tag:     "plain_text",
							Content: "更多",
						},
						Url: alert.GeneratorURl,
					},
				},
			})
		case "resolved":
			msg.MsgCard.Header = feishu.InteractiveMessageCardHeader{
				Template: "green",
				Title: feishu.InteractiveMessageCardHeaderTitle{
					Content: fmt.Sprintf("[告警恢复] %s", alert.Labels["alertname"]),
					Tag:     "plain_text",
				},
			}

			msg.MsgCard.Elements = append(msg.MsgCard.Elements, feishu.InteractiveMessageCardElement{
				Tag: "action",
				Actions: []feishu.InteractiveMessageCardElementAction{
					{
						Tag:  "button",
						Type: "green",
						Text: feishu.InteractiveMessageCardElementText{
							Tag:     "plain_text",
							Content: "更多",
						},
						Url: alert.GeneratorURl,
					},
				},
			})
		}

		client := feishu.NewClient(config.Conf.Feishu.Secret)
		client.SendMessage(msg)
	}

}
