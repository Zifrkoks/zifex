package models

import (
	"crypto/hmac"
	"crypto/sha256"
	b64 "encoding/base64"
	"encoding/json"
	"errors"
	"strings"
)

type JwtTokenBuilder struct {
	header  []byte
	payload []byte
	secret  []byte
}

func NewTokenBuilder() *JwtTokenBuilder {
	return &JwtTokenBuilder{}
}

func (b *JwtTokenBuilder) SetSecret(key string) {
	b.secret = []byte(key)
	return
}
func (b *JwtTokenBuilder) AddToHeader(key string, value string) (err error) {
	var data map[string]string
	if err = json.Unmarshal(b.header, &data); err != nil {
		return errors.New("error in unmarshal header")
	}
	data[key] = value
	if b.header, err = json.Marshal(data); err != nil {
		return errors.New("error in marshal header")
	}
	return nil
}
func (b *JwtTokenBuilder) AddToPayload(key string, value string) (err error) {
	var data map[string]string
	if err = json.Unmarshal(b.payload, &data); err != nil {
		return errors.New("error in unmarshal payload")
	}
	data[key] = value
	if b.payload, err = json.Marshal(data); err != nil {
		return errors.New("error in marshal payload")
	}
	return nil
}

func (b *JwtTokenBuilder) Build() (token string, err error) {
	mac := hmac.New(sha256.New, []byte(b.secret))
	data := b64.StdEncoding.EncodeToString([]byte(b.header)) + "." + b64.StdEncoding.EncodeToString([]byte(b.payload))
	_, err = mac.Write([]byte(data))
	if err != nil {
		return
	}
	token = data + "." + string(mac.Sum(nil))
	return
}

func ValidateToken(token string, key string) error {
	mac := hmac.New(sha256.New, []byte(key))
	str := strings.Split(token, ".")
	if len(str) != 3 {
		return errors.New("token format invalid")
	}
	mac.Write([]byte(str[0] + "." + str[1]))
	secret := mac.Sum(nil)
	if string(secret[:]) == str[2] {
		return nil
	} else {
		return errors.New("invalid signature")
	}

}

type RefreshTokenBuilder struct {
}
