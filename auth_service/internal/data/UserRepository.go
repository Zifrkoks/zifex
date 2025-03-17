package data

import (
	. "zifex_auth_service/internal/service/models"
)

type UserRepository struct {
}

func (r UserRepository) GetById(id uint) (*User, error) {
	return nil, nil
}

func (r UserRepository) GetByUsername(username string) (*User, error) {
	return nil, nil
}

func (r UserRepository) AddPermisions(usernameOrId interface{}, permissions map[string]string) (bool, error) {
	return false, nil
}
