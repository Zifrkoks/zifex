package repo

import (
	. "zifex_trade_service/internal/services/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func (UserRepository) GetAll() []User {
	return nil
}

func (UserRepository) GetAllPaging(page uint, count uint) []User {
	return nil
}

func (UserRepository) Get(id uint) (*User, error) {
	return nil, nil
}

func (UserRepository) Create(user *User) (u *User, err error) {
	return
}
func (UserRepository) Update(user *User) error {
	return nil
}

func (UserRepository) Delete(user *User) *User {
	return nil
}

func NewUserRepository() *UserRepository {
	return nil
}
