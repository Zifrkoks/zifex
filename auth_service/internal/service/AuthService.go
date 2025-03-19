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

func (service AuthService) Reg(username string, password string) (err error) {

	user, errch := service.users.GetByUsernamePromise(username)
	bytes := sha256.Sum256([]byte(password))
	password = string(bytes[:])
	err = <-errch
	if user != nil {
		return errors.New("username unavailable")
	}
	user = &User{Username: username, Password: password}
	err = service.users.Create(user)
	if err != nil {
		return errors.New("db error")
	}
	return nil
}

func (service AuthService) Login(username string, password string) (token string, err error) {

	user, errch := service.users.GetByUsernamePromise(username)

	bytes := sha256.Sum256([]byte(password))
	password = string(bytes[:])
	err = <-errch
	if err != nil {
		return "", errors.New("user not found")
	}

	if password != user.Password {
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

func (service AuthService) ChangePassword(newP string, username string) (err error) {
	user, errch := service.users.GetByUsernamePromise(username)
	bytes := sha256.Sum256([]byte(newP))
	err = <-errch
	if err != nil {
		return errors.New("user repo error:" + err.Error())
	}
	user.Password = string(bytes[:])
	err = service.users.Update(user)
	if err != nil {
		return errors.New("db error:" + err.Error())
	}
	return nil
}

func (service AuthService) ValidateToken(token string) error {
	return ValidateToken(token, viper.GetString("jwt.secret"))
}

func (service AuthService) RefreshToken(oldRT string) (jwt string, newRT string, err error) {
	return
}
func (service AuthService) AddPermisions(map[string]string) error {
	return nil
}

func (service AuthService) CheckUsernameIsAvailble(username string) error {
	return nil
}
