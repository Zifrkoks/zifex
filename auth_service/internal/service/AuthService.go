package service

import (
	. "zifex_auth_service/internal/data"
	. "zifex_auth_service/internal/service/models"
)

type AuthService struct {
	users *UserRepository
}
