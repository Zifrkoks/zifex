package models

type (
	Crypto struct {
		ID                 uint
		Symbol             string
		LastPricesToCrypto map[string]uint
		Count              uint64
		MinUnitDegree      uint8
		CorrToReal         int8
	}

	Trade struct {
		ID          uint
		Castomer    uint
		Status      Status
		Buy         string
		Sell        string
		PricePerInt uint64
		SellCount   uint64
		SoldCount   uint64
	}

	Status uint8
)

const (
	Created Status = iota
	PartlySuccess
	Success
	Cenceled
)
