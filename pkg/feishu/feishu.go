package feishu

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
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

func (c Client) SendMessage(msg Message) error {
	requestBody, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	requestUrl := fmt.Sprintf("%s/%s", OpenApi, c.Secret)
	response, err := http.Post(requestUrl, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	var responseData map[string]interface{}
	if err := json.Unmarshal(data, &responseData); err != nil {
		return err
	}

	if _, ok := responseData["StatusMessage"]; !ok {
		return errors.New(responseData["msg"].(string))
	}

	return nil
}
