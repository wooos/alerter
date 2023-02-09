package dingtalk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
	OpenApi         string      = "https://oapi.dingtalk.com/robot/send"
	MarkdownMessage MessageType = "markdown"
)

type Client struct {
	// AccessToken access token
	AccessToken string
	// Secret sign secret
	Secret string
}

// NewClient return dingtalk client
func NewClient(access_token, secret string) Client {
	return Client{
		AccessToken: access_token,
		Secret:      secret,
	}
}

// SendMessage
func (c Client) SendMessage(msg Message) {
	requestBody, err := json.Marshal(msg)
	if err != nil {
		log.Println(err)
	}

	requestUrl := fmt.Sprintf("%s?access_token=%s", OpenApi, c.AccessToken)
	response, err := http.Post(requestUrl, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Println(err)
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
	}

	log.Println(string(data))
}
