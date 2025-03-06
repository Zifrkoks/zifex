package repo

import (
	. "zifex_trade_service/internal/services/models"

	"gorm.io/gorm"
)

type TradeRepository struct {
	db *gorm.DB
}

func (TradeRepository) GetAllForUser(username string) []Trade {
	return nil
}
func (TradeRepository) GetActiveByBuyAndSell(bought string, sold string) (trades []Trade, err error) {
	return
}
func (TradeRepository) GetActiveByBuySellPrice(bought string, sold string, price uint64) (trades []Trade, err error) {
	return
}
func (TradeRepository) GetAllForUserPaging(page uint, count uint, username string) (trades []Trade, err error) {
	return
}
func (TradeRepository) Get(id uint) (*Trade, error) {
	return nil, nil
}

func (TradeRepository) Create(trade *Trade) error {
	return nil
}
func (TradeRepository) SaveWithUser(trade *Trade, user *User) error {
	return nil
}

func NewTradeRepository() *TradeRepository {
	return nil
}
