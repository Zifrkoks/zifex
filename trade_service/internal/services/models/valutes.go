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
)
