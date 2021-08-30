package portfolio

type Dollars float32

type Company struct {
	Name   string
	Symbol string
	Market int32
	Shares int
	Price  Dollars
	Earnings
	FinacialPosition
}

type FinacialPosition struct {
	CurrentAssets      Dollars
	CurrentLiabilities Dollars
	LongAssets         Dollars
	LongLiabilities    Dollars
}

type Earnings struct {
	Revenue     Dollars
	Expenditure Dollars
}
