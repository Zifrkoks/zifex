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
		count := tradesByPrices[uint((*trades)[i].PricePerInt)] + (*trades)[i].SoldCount
		tradesByPrices[uint((*trades)[i].PricePerInt)] = count
	}
	return tradesByPrices, nil
}
func (service TradeService) CreateTrade(trade Trade) (result Trade, err error) {
	var oppositePrice uint64
	oppositePrice = 1 / trade.PricePerInt // пока додумать как перевести в интовую форму а не float, без потерь в точности
	oppositeTrades, err := service.trades.GetActiveByBuySellPrice(trade.Sell, trade.Buy, oppositePrice)
	if err != nil {
		return trade, err
	}

	tradeBuyCripto, err := service.cryptos.GetBySymbol(trade.Buy)
	if err != nil {
		return trade, err
	}
	tradeSellCripto, err := service.cryptos.GetBySymbol(trade.Sell)
	if err != nil {
		return trade, err
	}
	for i := 0; i < len(*oppositeTrades); i++ {

	}

}
