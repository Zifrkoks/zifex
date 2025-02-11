package repo

import (
	. "zifex_trade_service/internal/services/models"

	"gorm.io/gorm"
)

type TariffRepository struct {
	db *gorm.DB
}

func (TariffRepository) GetAll() []Tariff {
	return nil
}

func (TariffRepository) GetAllPaging(page uint, count uint) []Tariff {
	return nil
}

func (TariffRepository) Get(id uint) *Tariff {
	return nil
}

func (TariffRepository) Create(Tariff *Tariff) *Tariff {
	return nil
}
func (TariffRepository) Update(Tariff *Tariff) *Tariff {
	return nil
}

func (TariffRepository) Delete(Tariff *Tariff) *Tariff {
	return nil
}

func NewTariffRepository() *TradeRepository {
	return nil
}
