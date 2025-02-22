package models

type (
	Crypto struct {
		ID                 uint
		Symbol             string
		LastPricesToCrypto map[string]uint
		Count              uint64
		CorrToReal         int8
	}
)
