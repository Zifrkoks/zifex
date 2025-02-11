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
func (TradeRepository) GetActiveByBuyAndSell(bought string, sold string) (trades *[]Trade, err error) {
	return
}
func (TradeRepository) GetActiveByBuySellPrice(bought string, sold string, price uint64) (trades *[]Trade, err error) {
	return
}
func (TradeRepository) GetAllPaging(page uint, count uint) []Trade {
	return nil
}

func (TradeRepository) Get(id uint) *Trade {
	return nil
}

func (TradeRepository) Create(trade *Trade) *Trade {
	return nil
}
func (TradeRepository) Update(trade *Trade) *Trade {
	return nil
}

func (TradeRepository) Delete(Trade *Trade) *Trade {
	return nil
}

func (TradeRepository) Where(trade *Trade) *[]Trade {
	return nil
}

func NewTradeRepository() *TradeRepository {
	return nil
}
