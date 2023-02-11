package message

import (
	"fmt"

	"github.com/wooos/alerter/internal/config"
	"github.com/wooos/alerter/internal/pkg/metrics"
	"github.com/wooos/alerter/internal/pkg/request"
	"github.com/wooos/alerter/pkg/dingtalk"
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

		msg.Card.Elements = append(msg.Card.Elements, feishu.CardElement{
			Tag:     "markdown",
			Content: "**Labels**",
		})

		for k, v := range alert.Labels {
			if utils.StrInArray(k, IgnoreLabelKeys) {
				continue
			}

			msg.Card.Elements = append(msg.Card.Elements, feishu.CardElement{
				Tag:     "markdown",
				Content: fmt.Sprintf("  - %s: %s", k, v),
			})
		}

		msg.Card.Elements = append(msg.Card.Elements, feishu.CardElement{
			Tag:     "markdown",
			Content: "**Annotations**",
		})

		for k, v := range alert.Annotations {
			msg.Card.Elements = append(msg.Card.Elements, feishu.CardElement{
				Tag:     "markdown",
				Content: fmt.Sprintf("  - %s: %s", k, v),
			})
		}

		msg.Card.Elements = append(msg.Card.Elements, feishu.CardElement{
			Tag: "hr",
		})

		switch alert.Status {
		case "firing":
			msg.Card.Header = feishu.CardHeader{
				Template: "red",
				Title: feishu.CardHeaderTitle{
					Content: fmt.Sprintf("[告警触发] %s", alert.Labels["alertname"]),
					Tag:     "plain_text",
				},
			}

			msg.Card.Elements = append(msg.Card.Elements, feishu.CardElement{
				Tag: "action",
				Actions: []feishu.CardElementAction{
					{
						Tag:  "button",
						Type: "danger",
						Text: feishu.CardElementText{
							Tag:     "plain_text",
							Content: "更多",
						},
						Url: alert.GeneratorURl,
					},
				},
			})
		case "resolved":
			msg.Card.Header = feishu.CardHeader{
				Template: "green",
				Title: feishu.CardHeaderTitle{
					Content: fmt.Sprintf("[告警恢复] %s", alert.Labels["alertname"]),
					Tag:     "plain_text",
				},
			}

			msg.Card.Elements = append(msg.Card.Elements, feishu.CardElement{
				Tag: "action",
				Actions: []feishu.CardElementAction{
					{
						Tag:  "button",
						Type: "green",
						Text: feishu.CardElementText{
							Tag:     "plain_text",
							Content: "更多",
						},
						Url: alert.GeneratorURl,
					},
				},
			})
		}

		client := feishu.NewClient(config.Conf.Feishu.Secret)
		if err := client.SendMessage(msg); err != nil {
			metrics.AlerterSendMessageTotal.WithLabelValues("feishu", "send_failed").Inc()
		} else {
			metrics.AlerterSendMessageTotal.WithLabelValues("feishu", "send_success").Inc()
		}
	}

	if config.Conf.Dingtalk.Enabled {
		msg := dingtalk.ActionCardMessage{
			MsgType: dingtalk.MsgTypeActionCard,
			ActionCard: dingtalk.ActionCard{
				Title:          "",
				Text:           "",
				BtnOrientation: "",
				SingleTitle:    "",
				SingleURL:      "",
				Btns:           []dingtalk.ActionCardBtn{},
			},
		}

		var msgText string
		switch alert.Status {
		case "firing":
			msg.ActionCard.Title = fmt.Sprintf("【告警通知】%s", alert.Labels["alertname"])
			msgText = fmt.Sprintf("### 【告警通知】%s", alert.Labels["alertname"])
		case "resolved":
			msg.ActionCard.Title = fmt.Sprintf("【告警恢复】%s", alert.Labels["alertname"])
			msgText = fmt.Sprintf("### 【告警恢复】%s", alert.Labels["alertname"])
		}

		msgText = fmt.Sprintf("%s\n\n%s", msgText, "**Labels**")

		for k, v := range alert.Labels {
			if utils.StrInArray(k, IgnoreLabelKeys) {
				continue
			}

			msgText = fmt.Sprintf("%s\n\n> %s: %s", msgText, k, v)
		}

		msgText = fmt.Sprintf("%s\n\n%s", msgText, "**Annotations**")

		for k, v := range alert.Annotations {
			msgText = fmt.Sprintf("%s\n\n> %s: %s", msgText, k, v)
		}

		msg.ActionCard.Text = msgText
		msg.ActionCard.SingleURL = fmt.Sprintf("dingtalk://dingtalkclient/page/link?url=%s&pc_slide=false", alert.GeneratorURl)

		msg.ActionCard.SingleTitle = "更多"

		client := dingtalk.NewClient(config.Conf.Dingtalk.Token, config.Conf.Dingtalk.Secret)
		if err := client.SendMessage(msg); err != nil {
			metrics.AlerterSendMessageTotal.WithLabelValues("dingtalk", "send_failed").Inc()
		} else {
			metrics.AlerterSendMessageTotal.WithLabelValues("dingtalk", "send_success").Inc()
		}
	}
}
