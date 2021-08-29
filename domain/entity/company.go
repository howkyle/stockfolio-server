package entity

import "github.com/howkyle/stockfolio-server/domain/types"

type Company struct {
	Name   string
	Symbol string
	Market int32
	Shares int
	Price  types.Dollars
	Earnings
	FinacialPosition
}

type FinacialPosition struct {
	CurrentAssets      types.Dollars
	CurrentLiabilities types.Dollars
	LongAssets         types.Dollars
	LongLiabilities    types.Dollars
}

type Earnings struct {
	Revenue     types.Dollars
	Expenditure types.Dollars
}
