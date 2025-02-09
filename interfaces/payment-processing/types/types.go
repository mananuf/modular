package types

type CardType string

type Payment interface {
	Process()
	Refund()
}

type Details struct {
	Expiry     string
	Balance    uint
}

type CreditCard struct {
	Card       CardType
	CCV        uint8
	CardNumber uint
	Details
}

type Paypal struct {
	Email    uint
	IsActive bool
	Details
}

type Crypto struct{
	WalletAddress string
	Balance uint
}
