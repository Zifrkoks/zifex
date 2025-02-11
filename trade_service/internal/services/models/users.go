package models

type (
	User struct {
		ID             uint
		CryptoWallets  map[string]uint
		FavoriteCripto []string
		Username       string
		Tariff         uint
		SecLevel       byte
		Trades         []uint
	}

	Tariff struct {
		ID         uint
		Comission  uint
		Permission uint
	}
)
