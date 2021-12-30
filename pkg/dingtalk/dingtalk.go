package dingtalk

const (
	WebhookApi = "https://oapi.dingtalk.com/robot/send"
)

type Client struct {
	// AccessToken access token
	AccessToken string
	// Secret sign secret
	Secret string
}

func NewClient(access_token, secret string) Client {
	return Client{
		AccessToken: access_token,
		Secret:      secret,
	}
}
