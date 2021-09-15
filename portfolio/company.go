package portfolio

import (
	"time"

	"gorm.io/gorm"
)

type Company struct {
	gorm.Model
	PortfolioID      uint
	Name             string
	Symbol           string
	Market           int32
	Shares           int
	FinancialReports []FinancialReport
}

type FinancialReport struct {
	gorm.Model
	CompanyID uint
	Year      time.Time
	Quarter   int
	Price     Dollars
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
	Income      Dollars
	Expenditure Dollars
}
