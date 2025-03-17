package service

import (
	"crypto/sha256"
	"errors"
	. "zifex_auth_service/internal/data"
	. "zifex_auth_service/internal/service/models"
	. "zifex_auth_service/pkg/models"

	"github.com/spf13/viper"
)

type AuthService struct {
	users *UserRepository
}

func (service AuthService) Reg(login string, password string) error {
	return nil
}

func (service AuthService) Login(login string, password string) (token string, err error) {
	var passCh chan string
	var user *User
	go func() {
		user, err = service.users.GetByUsername(login)
		passCh <- user.Password
	}()
	hash := sha256.New()
	hash.Write([]byte(password))
	password = string(hash.Sum(nil))
	userPass := <-passCh
	if err != nil {
		return "", errors.New("user not found")
	}

	if password != userPass {
		return "", errors.New("password wrong")
	}
	builder := NewTokenBuilder()
	builder.AddToHeader("type", "jwt")
	builder.AddToHeader("alg", "HS256")
	builder.AddToHeader("issuer", "zifex")
	for key, value := range user.Permissions {
		builder.AddToPayload(key, value)
	}
	builder.AddToHeader("exp", viper.GetString("jwt.exp"))
	builder.SetSecret(viper.GetString("jwt.secret"))
	token, err = builder.Build()
	return
}

func (service AuthService) ChangePassword(newP string) error {
	return nil
}

func (service AuthService) ValidateToken(token string) error {
	return nil
}

func (service AuthService) AddPermisions(map[string]string) error {
	return nil
}
