package portfolio

import "gorm.io/gorm"

type Dollars float32

type Portfolio struct {
	UserID uint
	gorm.Model
	Companies []Company
}

type Company struct {
	gorm.Model
	PortfolioID uint
	Name        string
	Symbol      string
	Market      int32
	Shares      int
	Price       Dollars
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
