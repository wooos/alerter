package dingtalk

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/wooos/alerter/pkg/utils"
)

const (
	OpenApi string = "https://oapi.dingtalk.com/robot/send"
)

type Client struct {
	// AccessToken access token
	AccessToken string
	// Secret sign secret
	Secret string
}

type Response struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// NewClient return dingtalk client
func NewClient(access_token, secret string) Client {
	return Client{
		AccessToken: access_token,
		Secret:      secret,
	}
}

// SendMessage
func (c Client) SendMessage(msg Message) error {
	requestBody, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	var requestUrl string
	if c.Secret != "" {
		now := time.Now().UnixMilli()
		strToSign := fmt.Sprintf("%d\n%s", now, c.Secret)
		signStr, err := utils.HMacSHA256ToBase64(c.Secret, strToSign)
		if err != nil {
			return err
		}
		requestUrl = fmt.Sprintf("%s?access_token=%s&timestamp=%d&sign=%s", OpenApi, c.AccessToken, now, signStr)
	} else {
		requestUrl = fmt.Sprintf("%s?access_token=%s", OpenApi, c.AccessToken)
	}

	response, err := http.Post(requestUrl, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	var responseData Response
	if err := json.Unmarshal(data, &responseData); err != nil {
		return err
	}

	if responseData.ErrCode != 0 {
		return errors.New(responseData.ErrMsg)
	}

	return nil
}
