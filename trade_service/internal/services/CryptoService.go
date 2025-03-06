package services

import (
	. "zifex_trade_service/internal/repo"
	. "zifex_trade_service/internal/services/models"
)

type CryptoService struct {
	users   *UserRepository
	cryptos *CryptoRepository
}

func (service CryptoService) CreateCrypto(crypto Crypto) {

}
