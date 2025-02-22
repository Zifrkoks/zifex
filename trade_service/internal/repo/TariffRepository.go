package repo

import (
	. "zifex_trade_service/internal/services/models"

	"gorm.io/gorm"
)

type TariffRepository struct {
	db *gorm.DB
}

func (TariffRepository) GetAll() (tar []Tariff, err error) {
	return
}

func (TariffRepository) GetAllPaging(page uint, count uint) (tar []Tariff, err error) {
	return
}

func (TariffRepository) Get(id uint) (tar *Tariff, err error) {
	return
}
func (TariffRepository) GetCommision(id uint) (tar *Tariff, err error) {
	return
}

func (TariffRepository) Create(Tariff *Tariff) (tar *Tariff, err error) {
	return
}
func (TariffRepository) Update(Tariff *Tariff) (tar *Tariff, err error) {
	return
}

func (TariffRepository) Delete(Tariff *Tariff) (tar *Tariff, err error) {
	return
}

func NewTariffRepository() *TradeRepository {
	return nil
}
