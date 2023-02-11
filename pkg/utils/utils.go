package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

// StrInArray return ture if
func StrInArray(s string, a []string) bool {
	for _, v := range a {
		if s == v {
			return true
		}
	}

	return false
}

func HMacSHA256ToBase64(secret string, data string) (string, error) {
	mac := hmac.New(sha256.New, []byte(secret))
	_, err := mac.Write([]byte(data))
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(mac.Sum(nil)), nil
}
