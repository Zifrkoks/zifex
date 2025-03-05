package services

import (
	. "zifex_trade_service/internal/repo"
	. "zifex_trade_service/internal/services/models"
)

type UserService struct {
	users   *UserRepository
	cryptos *CryptoRepository
	trades  *TradeRepository
	tariffs *TariffRepository
}

func (service UserService) CreateUser(u User) error {
	tariff, err := service.tariffs.GetStandartTariff()
	u.TariffProcent = tariff.Comission
	u.SecLevel = 0
	_, err = service.users.Create(&u)
	return err
}

func (service UserService) updateTariff(uId uint, TarId uint) error {
	tariff, err := service.tariffs.Get(TarId)
	if err != nil {
		return err
	}
	user, err1 := service.users.Get(uId)
	if err1 != nil {
		return err
	}
	user.TariffProcent = tariff.Comission
	err = service.users.Update(user)
	if err != nil {
		return err
	}
	return nil
}
