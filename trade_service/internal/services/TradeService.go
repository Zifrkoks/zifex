package services

import (
	"errors"
	"math"
	. "zifex_trade_service/internal/repo"
	. "zifex_trade_service/internal/services/models"
)

type TradeService struct {
	users   *UserRepository
	cryptos *CryptoRepository
	trades  *TradeRepository
	tariffs *TariffRepository
}

func (service TradeService) FindActiveTradesFull(bought string, sold string) (tradesByPrices map[uint][]Trade, err error) {
	trades, err := service.trades.GetActiveByBuyAndSell(bought, sold)
	if err != nil {
		return
	}
	tradesByPrices = map[uint]([]Trade){}
	for i := 0; i < len(trades); i++ {
		arr := tradesByPrices[uint(trades[i].GetPrice())]
		if arr != nil {
			arr = append(arr, trades[i])
		} else {
			arr = []Trade{trades[i]}
		}
		tradesByPrices[uint(trades[i].GetPrice())] = arr
	}
	return tradesByPrices, nil
}

func (service TradeService) FindActiveTrades(bought string, sold string) (tradesByPrices map[uint64]uint64, err error) {
	trades, err := service.trades.GetActiveByBuyAndSell(bought, sold)
	if err != nil {
		return
	}
	tradesByPrices = map[uint64]uint64{}
	for i := 0; i < len(trades); i++ {
		price := (trades)[i].GetPrice()
		tradesByPrices[price] += (trades)[i].OnSaleCount
	}
	return tradesByPrices, nil
}

func (service TradeService) CreateTrade(trade Trade) (result Trade, err error) {
	creator, err := service.users.Get(trade.Castomer)
	if err != nil {
		return trade, err
	}
	if err != nil {
		return trade, err
	}
	oppositePrice := trade.GetReversePrice()
	oppositeTrades, err := service.trades.GetActiveByBuySellPrice(trade.Sell, trade.Buy, oppositePrice)
	if err != nil {
		return trade, err
	}
	if oppositeTrades == nil {
		service.trades.Create(&trade)
		return
	}

	for i := 0; i < len(oppositeTrades); i++ {
		castomer, err := service.users.Get(oppositeTrades[i].Castomer)
		if err != nil {
			continue
		}
		if oppositeTrades[i].OnSaleCount < trade.BuyCount {
			service.closeTradePartly(&trade, creator, oppositeTrades[i])
			service.closeTrade(&oppositeTrades[i], castomer)
			continue
		}
		if oppositeTrades[i].OnSaleCount > trade.BuyCount {
			service.closeTradePartly(&oppositeTrades[i], castomer, trade)
			service.closeTrade(&trade, creator)
			continue
		}
		if oppositeTrades[i].OnSaleCount == trade.BuyCount {
			service.closeTrade(&trade, creator)
			service.closeTrade(&oppositeTrades[i], castomer)
			continue
		}

	}

	return trade, err
}

func (service TradeService) closeTrade(tr1 *Trade, u1 *User) error {
	if tr1.Castomer != u1.ID {
		return errors.New("user is not creator of trade")
	}
	if _, ok := u1.CryptoWallets[tr1.Buy]; ok {
		u1.CryptoWallets[tr1.Buy] += tr1.BuyCount
	} else {
		u1.CryptoWallets[tr1.Buy] = tr1.BuyCount
	}
	u1.FreezeCrypto[tr1.Sell] -= tr1.OnSaleCount
	if u1.FreezeCrypto[tr1.Sell] <= 0 {
		delete(u1.FreezeCrypto, tr1.Sell)
	}
	delete(u1.FreezeCommision, tr1.ID)
	tr1.BuyCount = 0
	tr1.OnSaleCount = 0
	tr1.Status = Success
	return nil
}

func (service TradeService) closeTradePartly(tr1 *Trade, u1 *User, tr2 Trade) error {
	if tr1.BuyCount < tr1.OnSaleCount {
		return errors.New("cant close partly because trade buy count less then trade that close")
	}
	commision := uint64(math.Ceil(float64((tr2.OnSaleCount * uint64(u1.TariffProcent))) / float64(100)))
	u1.FreezeCommision[tr1.ID] -= commision
	tr1.BuyCount -= tr2.OnSaleCount
	tr1.OnSaleCount -= tr2.BuyCount
	if _, ok := u1.CryptoWallets[tr1.Buy]; ok {
		u1.CryptoWallets[tr1.Buy] += tr2.OnSaleCount
	} else {
		u1.CryptoWallets[tr1.Buy] = tr2.OnSaleCount
	}
	u1.FreezeCrypto[tr1.Sell] -= tr2.BuyCount
	tr1.Status = PartlySuccess
	return nil
}
