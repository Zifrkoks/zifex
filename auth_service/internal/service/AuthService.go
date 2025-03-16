package service

import (
	. "zifex_auth_service/internal/data"
	. "zifex_auth_service/internal/service/models"
)

type AuthService struct {
	users *UserRepository
}

func (service AuthService) Reg(oldP string, newP string) error {
	return nil
}

func (service AuthService) Login(login string, password string) (token string, err error) {

	return "", nil
}

func (service AuthService) ChangePassword(oldP string, newP string) error {
	return nil
}
