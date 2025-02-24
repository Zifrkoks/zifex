package repo

import (
	. "zifex_trade_service/internal/services/models"

	"gorm.io/gorm"
)

type CryptoRepository struct {
	db *gorm.DB
}

func NewCryptoRepository() *TradeRepository {
	return nil
}
func (CryptoRepository) GetAll() (crypto []Crypto, err error) {
	return nil, nil
}

func (CryptoRepository) GetAllPaging(page uint, count uint) (crypto []Crypto, err error) {
	return nil, nil
}

func (CryptoRepository) Get(id uint) (crypto *Crypto, err error) {
	return nil, nil
}

func (CryptoRepository) GetBySymbol(symbol string) (crypto *Crypto, err error) {
	return nil, nil
}
func (CryptoRepository) GetPoint(symbol string) (point uint8, err error) {
	return 0, nil
}
func (CryptoRepository) Create(Crypto *Crypto) (crypto *Crypto, err error) {
	return nil, nil
}
func (CryptoRepository) Update(Crypto *Crypto) (crypto *Crypto, err error) {
	return nil, nil
}

func (CryptoRepository) Delete(Crypto *Crypto) (crypto *Crypto, err error) {
	return nil, nil
}

func (CryptoRepository) CheckNames(names ...string) error {
	return nil
}
