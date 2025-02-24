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
	for _, trade := range trades {
		arr := tradesByPrices[uint(trade.GetPrice())]
		if arr != nil {
			arr = append(arr, trade)
		} else {
			arr = []Trade{trade}
		}
		tradesByPrices[uint(trade.GetPrice())] = arr
	}
	return tradesByPrices, nil
}

func (service TradeService) FindActiveTrades(bought string, sold string) (tradesByPrices map[uint64]uint64, err error) {
	trades, err := service.trades.GetActiveByBuyAndSell(bought, sold)
	if err != nil {
		return
	}
	tradesByPrices = map[uint64]uint64{}
	for _, trade := range trades {
		price := trade.GetPrice()
		tradesByPrices[price] += trade.OnSaleCount
	}
	return tradesByPrices, nil
}

func (service TradeService) CreateTrade(trade Trade) (err error) {
	creator, err := service.users.Get(trade.Castomer)
	if err != nil {
		return err
	}
	oppositePrice := trade.GetReversePrice()
	oppositeTrades, err := service.trades.GetActiveByBuySellPrice(trade.Sell, trade.Buy, oppositePrice)
	if err != nil {
		return err
	}
	if oppositeTrades == nil {
		service.createTrade(&trade, creator)
		return
	}

	for _, opTrade := range oppositeTrades {
		castomer, err := service.users.Get(opTrade.Castomer)
		if err != nil {
			continue
		}
		if opTrade.OnSaleCount < trade.BuyCount {
			service.closeTradePartly(&trade, creator, opTrade)
			service.closeTrade(&opTrade, castomer)
			continue
		}
		if opTrade.OnSaleCount > trade.BuyCount {
			service.closeTradePartly(&opTrade, castomer, trade)
			service.closeTrade(&trade, creator)
			continue
		}
		if opTrade.OnSaleCount == trade.BuyCount {
			service.closeTrade(&trade, creator)
			service.closeTrade(&opTrade, castomer)
			continue
		}

	}

	return err
}

func (service TradeService) createTrade(tr *Trade, u *User) error {
	if tr.Castomer != u.ID {
		return errors.New("castomer id and user id not equal")
	}
	if err := service.cryptos.CheckNames(tr.Buy, tr.Sell); err != nil {
		return err
	}
	totalCost := tr.OnSaleCount + (tr.OnSaleCount * uint64(u.TariffProcent) / 10000)

	if val, ok := u.CryptoWallets[tr.Sell]; !ok || totalCost > val {
		return errors.New("user has no so mach money")
	}

	u.CryptoWallets[tr.Sell] -= totalCost
	u.FreezeCrypto[tr.Sell] += tr.OnSaleCount
	return nil
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
	commision := uint64(math.Ceil(float64((tr2.OnSaleCount * uint64(u1.TariffProcent))) / float64(10000)))
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
