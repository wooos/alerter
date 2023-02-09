package feishu

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
	OpenApi string = "https://open.feishu.cn/open-apis/bot/v2/hook/"
)

type Client struct {
	Secret string
}

func NewClient(secret string) Client {
	return Client{
		Secret: secret,
	}
}

func (c Client) SendMessage(msg Message) {
	requestBody, err := json.Marshal(msg)
	if err != nil {
		log.Println(err)
	}

	requestUrl := fmt.Sprintf("%s/%s", OpenApi, c.Secret)
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
