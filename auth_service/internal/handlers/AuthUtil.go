package handlers

import (
	"crypto/hmac"
	sha256 "crypto/sha256"
	b64 "encoding/base64"
	"errors"
	"hash"
	"strings"

	"github.com/spf13/viper"
)

func HashSHA256(value, salt string) string {
	var s = append([]byte(value), []byte(salt)...)
	bytes := sha256.Sum256(s)
	return string(bytes[:])
}
func ValidateToken(token string) error {
	var mac hash.Hash
	mac = hmac.New(sha256.New, []byte(viper.GetString("SECRET_KEY")))
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
func CreateToken(header string, payload string) (token string, err error) {
	mac := hmac.New(sha256.New, []byte(viper.GetString("SECRET_KEY")))
	data := b64.StdEncoding.EncodeToString([]byte(header)) + "." + b64.StdEncoding.EncodeToString([]byte(payload))
	_, err = mac.Write([]byte(data))
	if err != nil {
		return "", err
	}
	secret := mac.Sum(nil)
	return (data + "." + string(secret)), nil
}
