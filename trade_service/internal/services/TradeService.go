package services

import (
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
	for i := 0; i < len(*trades); i++ {
		arr := tradesByPrices[uint((*trades)[i].PricePerInt)]
		if arr != nil {
			arr = append(arr, (*trades)[i])
		} else {
			arr = []Trade{(*trades)[i]}
		}
		tradesByPrices[uint((*trades)[i].PricePerInt)] = arr
	}
	return tradesByPrices, nil
}

func (service TradeService) FindActiveTrades(bought string, sold string) (tradesByPrices map[uint]uint64, err error) {
	trades, err := service.trades.GetActiveByBuyAndSell(bought, sold)
	if err != nil {
		return
	}
	tradesByPrices = map[uint]uint64{}
	for i := 0; i < len(*trades); i++ {
		count := tradesByPrices[uint((*trades)[i].GetPrice())] + (*trades)[i].SoldCount
		tradesByPrices[uint((*trades)[i].GetPrice())] = count
	}
	return tradesByPrices, nil
}

func (service TradeService) CreateTrade(trade Trade) (result Trade, err error) {
	BuyPoint, err := service.cryptos.GetPoint(trade.Buy)
	if err != nil {
		return trade, err
	}
	SellPoint, err := service.cryptos.GetPoint(trade.Sell)
	if err != nil {
		return trade, err
	}
	oppositePrice := trade.GetReversePrice(BuyPoint, SellPoint)
	oppositeTrades, err := service.trades.GetActiveByBuySellPrice(trade.Sell, trade.Buy, oppositePrice)
	if err != nil {
		return trade, err
	}
	if oppositeTrades == nil {
		//создание трейда
	}

	for i := 0; i < len(*oppositeTrades); i++ {
		if (*oppositeTrades)[i].OnSaleCount < trade.BuyCount {
			trade.BuyCount
			(*oppositeTrades)[i].OnSaleCount = 0

			//закрытие обратного трейда
		}
		if (*oppositeTrades)[i].OnSaleCount > trade.BuyCount {
			(*oppositeTrades)[i].OnSaleCount -= trade.BuyCount
			(*oppositeTrades)[i].SoldCount += trade.BuyCount
			//создание  выполненого трейда
			//изменение статуса обратного трейда
		}
		if (*oppositeTrades)[i].OnSaleCount == trade.BuyCount {
			(*oppositeTrades)[i].OnSaleCount = (*oppositeTrades)[i].OnSaleCount - trade.BuyCount
			//создание  выполненого трейда
			//закрытие обратного трейда
		}

	}

	return trade, err
}

func (service TradeService) reversePrice(price uint64) uint64 {
	return 0
}
