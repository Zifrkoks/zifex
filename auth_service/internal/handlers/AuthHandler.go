package handlers

import (
	sha256 "crypto/sha256"
	"encoding/base64"
	"strings"
	. "zifex_auth_service/internal/service/models"
)

func HashSHA256(value, salt string) string {
	var s = append([]byte(value), []byte(salt)...)
	bytes := sha256.Sum256(s)
	return string(bytes[:])
}
func ValidateToken(token string) bool {
	str := strings.Split(token, ".")
	if len(str) != 3 {
		return false
	}
	header, err1 := base64.StdEncoding.DecodeString(str[0])
	payload, err2 := base64.StdEncoding.DecodeString(str[1])
	return true
}
func CreateToken(user *User) string {

}
