package repo

import (
	. "zifex_trade_service/internal/services/models"

	"gorm.io/gorm"
)

type TradeRepository struct {
	db *gorm.DB
}

func (TradeRepository) GetAll() []Trade {
	return nil
}
func (TradeRepository) GetActiveByBuyAndSell(bought string, sold string) (trades []Trade, err error) {
	return
}
func (TradeRepository) GetActiveByBuySellPrice(bought string, sold string, price uint64) (trades []Trade, err error) {
	return
}
func (TradeRepository) GetAllPaging(page uint, count uint) (trades *Trade, err error) {
	return
}

func (TradeRepository) Get(id uint) (*Trade, error) {
	return nil, nil
}

func (TradeRepository) Create(trade *Trade) (*Trade, error) {
	return nil, nil
}
func (TradeRepository) Update(trade *Trade) (*Trade, error) {
	return nil, nil
}

func (TradeRepository) Delete(Trade *Trade) (*Trade, error) {
	return nil, nil
}

func NewTradeRepository() *TradeRepository {
	return nil
}
