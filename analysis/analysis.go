package analysis

import (
	"github.com/howkyle/stockfolio-server/company"
)

type QuickAnalysis struct {
	Name               string
	Symbol             string
	Shares             int
	Price              company.Dollars
	CurrentAssets      company.DollarSlice
	CurrentLiabilities company.DollarSlice
	LongAssets         company.DollarSlice
	LongLiabilities    company.DollarSlice
	Income             company.DollarSlice
	Expenditure        company.DollarSlice
}

func (q QuickAnalysis) ToFinancialReport() company.FinancialReport {
	return company.FinancialReport{
		Shares: q.Shares,
		Price:  q.Price,
		Earnings: company.Earnings{
			Income:      q.Income.Total(),
			Expenditure: q.Expenditure.Total(),
		},
		FinacialPosition: company.FinacialPosition{
			CurrentAssets:      q.CurrentAssets.Total(),
			CurrentLiabilities: q.CurrentLiabilities.Total(),
			LongAssets:         q.LongAssets.Total(),
			LongLiabilities:    q.LongLiabilities.Total(),
		},
	}
}

type Result struct {
	Pe            float32
	Pb            float32
	CurrentRatio  float32
	DeRatio       float32
	EarningsYield float32
	Roe           float32
	ProfitMargin  float32
	Eps           company.Dollars
	Bvs           company.Dollars
	NetIncome     company.Dollars
	TotalEquity   company.Dollars
}

type Service interface {
	Analyze(c company.FinancialReport) (*Result, error)
}
