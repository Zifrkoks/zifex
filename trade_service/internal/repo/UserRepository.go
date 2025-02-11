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

func (UserRepository) Get(id uint) *User {
	return nil
}

func (UserRepository) Create(user *User) *User {
	return nil
}
func (UserRepository) Update(user *User) *User {
	return nil
}

func (UserRepository) Delete(user *User) *User {
	return nil
}

func NewUserRepository() *UserRepository {
	return nil
}
