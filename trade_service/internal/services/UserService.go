package services

import (
	. "zifex_trade_service/internal/repo"
	. "zifex_trade_service/internal/services/models"

	"github.com/spf13/viper"
)

type UserService struct {
	users   *UserRepository
	cryptos *CryptoRepository
	trades  *TradeRepository
}

func (service UserService) CreateUser(u User) error {
	u.TariffProcent = uint8(viper.GetUint("service.StandartTarriff"))
	u.SecLevel = 0
	_, err := service.users.Create(&u)
	return err
}

func (service UserService) UpdateTariff(username string, procent uint8) error {

	user, err := service.users.GetByUsername(username)
	if err != nil {
		return err
	}
	user.TariffProcent = procent
	err = service.users.Update(user)
	if err != nil {
		return err
	}
	return nil
}

func (service UserService) UpdateSecLvl(username string, lvl uint8) error {

	user, err := service.users.GetByUsername(username)
	if err != nil {
		return err
	}
	user.SecLevel = lvl
	err = service.users.Update(user)
	if err != nil {
		return err
	}
	return nil
}

func (service UserService) GetUserTrades(username string, count uint, page uint) ([]Trade, error) {
	return service.trades.GetAllForUserPaging(page, count, username)
}
