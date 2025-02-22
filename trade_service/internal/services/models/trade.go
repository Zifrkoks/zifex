package models

type (
	Trade struct {
		ID             uint
		Castomer       uint   //user's id
		Status         Status //Created,PartlySuccess,Success,Cenceled
		Buy            string //symbol of bought coin
		Sell           string //symbol of sold coin
		BuyCount       uint64 //count of coin what user wanna buy
		OnSaleCount    uint64 //count of coin what user wanna sell
		TotalSaleCount uint64 //Total count for sale
	}
	Status uint8
)

const (
	Created Status = iota
	PartlySuccess
	Success
	Cenceled
	PartlyCenceled
)

func (t Trade) GetReversePrice() uint64 {
	return 0
}
func (t Trade) GetPrice() uint64 {
	return 0
}
