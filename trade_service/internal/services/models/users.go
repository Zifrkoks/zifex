package models

type (
	User struct {
		ID              uint
		CryptoWallets   map[string]uint64 //[symbol]count
		FreezeCrypto    map[string]uint64
		FreezeCommision map[uint]uint64 //[trade id]countOfSellingCrypto
		FavoriteCripto  []string
		Username        string
		TariffProcent   uint8 //1 equal 0.001 procent
		SecLevel        byte
		Trades          []uint
	}

	Tariff struct {
		ID         uint
		Comission  uint
		Permission uint
	}
)
