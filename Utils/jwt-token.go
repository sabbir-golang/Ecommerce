package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
)

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type Payload struct {
	Sub         int    `json:"sub"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"Last_name"`
	Email       string `json:"email"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

func Base64Encode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
func CreateJwt(secret string, data Payload) (string, error) {
	header := Header{
		Alg: "HS256",
		Typ: "JWT",
	}
	byteHeader, err := json.Marshal(header)
	if err != nil {
		return "", err
	}
	bytePayload, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	HeaderBase64 := Base64Encode(byteHeader)
	PayloadBase64 := Base64Encode(bytePayload)
	message := HeaderBase64 + "." + PayloadBase64

	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(message))
	signiture := h.Sum(nil)
	signitureBase64 := Base64Encode(signiture)

	jwt := message + "." + signitureBase64
	return jwt, nil
}
