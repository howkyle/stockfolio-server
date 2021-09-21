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

type ReportAnalysis struct {
	Shares             int
	Price              company.Dollars
	CurrentAssets      company.Dollars
	CurrentLiabilities company.Dollars
	LongAssets         company.Dollars
	LongLiabilities    company.Dollars
	Income             company.Dollars
	Expenditure        company.Dollars
}

func (r ReportAnalysis) ToFinancialReport() company.FinancialReport {
	return company.FinancialReport{
		Shares: r.Shares,
		Price:  r.Price,
		Earnings: company.Earnings{
			Income:      r.Income,
			Expenditure: r.Expenditure,
		},
		FinacialPosition: company.FinacialPosition{
			CurrentAssets:      r.CurrentAssets,
			CurrentLiabilities: r.CurrentLiabilities,
			LongAssets:         r.LongAssets,
			LongLiabilities:    r.LongLiabilities,
		},
	}
}

//represents the result of an analysis
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
	//takes a financial report as input and runs analysis to return
	//result or error
	Analyze(c company.FinancialReport) (*Result, error)
}
