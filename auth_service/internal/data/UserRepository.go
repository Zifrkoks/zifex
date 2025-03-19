package data

import (
	. "zifex_auth_service/internal/service/models"
)

type UserRepository struct {
}

func (r *UserRepository) Create(user *User) error {
	panic("unimplemented")
}

func (r *UserRepository) Update(user *User) error {
	panic("unimplemented")
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

/*
you need to get error from errChannel for make sure that work with user is closed
*/
func (r UserRepository) GetByUsernamePromise(username string) (user *User, errChannel chan error) {
	return nil, make(chan error)
}

func (r UserRepository) CreateUserRefreshTokenPromise(username string) (refresh *RefreshToken, errChannel chan error) {
	return nil, make(chan error)
}
