package repo

import (
	. "zifex_trade_service/internal/services/models"

	"gorm.io/gorm"
)

type CryptoRepository struct {
	db *gorm.DB
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

func (CryptoRepository) Create(Crypto *Crypto) (crypto *Crypto, err error) {
	return nil, nil
}
func (CryptoRepository) Update(Crypto *Crypto) (crypto *Crypto, err error) {
	return nil, nil
}

func (CryptoRepository) Delete(Crypto *Crypto) (crypto *Crypto, err error) {
	return nil, nil
}

func NewCryptoRepository() *TradeRepository {
	return nil
}
