package analysis

import (
	"fmt"
	"log"

	"github.com/howkyle/stockfolio-server/company"
)

type Company struct {
	Name   string
	Symbol string
	Market int32
	Shares int
	Price  company.Dollars
	Earnings
	FinacialPosition
}

type FinacialPosition struct {
	CurrentAssets      company.DollarSlice
	CurrentLiabilities company.DollarSlice
	LongAssets         company.DollarSlice
	LongLiabilities    company.DollarSlice
}

type Earnings struct {
	Income      company.DollarSlice
	Expenditure company.DollarSlice
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

func Analyze(c *Company) (*Result, error) {
	netIncome := netIncome(c.Income, c.Expenditure)
	eps, err := eps(netIncome, float32(c.Shares))
	if err != nil {
		return nil, err
	}
	longEquity := equity(c.LongAssets, c.LongLiabilities)
	pe, err := finRatio(c.Price, eps)
	if err != nil {
		return nil, err

	}
	currentEquity := equity(c.CurrentAssets, c.CurrentLiabilities)
	totalEquity := currentEquity + longEquity
	bookVal, err := bvs(totalEquity, float32(c.Shares))
	if err != nil {
		return nil, err

	}
	pb, err := finRatio(c.Price, bookVal)
	if err != nil {
		return nil, err

	}
	currentRatio, err := finRatio(c.CurrentAssets.Total(), c.CurrentLiabilities.Total())

	if err != nil {
		log.Printf("unable to calculate current ratio: %v", err)
		return nil, err
	}

	deRatio, err := finRatio(c.CurrentLiabilities.Total()+c.LongLiabilities.Total(), totalEquity)
	if err != nil {
		return nil, err

	}
	earningsYield, err := finRatio(eps, c.Price)
	if err != nil {
		return nil, err

	}
	roe, err := finRatio(eps, bookVal)
	if err != nil {
		return nil, err

	}
	profitMargin, err := finRatio(netIncome, c.Income.Total())
	if err != nil {
		return nil, err

	}

	return &Result{pe, pb, currentRatio, deRatio, earningsYield, roe, profitMargin, eps, bookVal, netIncome, totalEquity}, nil
}

func PEMultiplePrice(market *Market, price company.Dollars, pe float32) (company.Dollars, error) {
	if pe == 0 {
		return -1, fmt.Errorf("cant divide by 0")
	}
	return (company.Dollars(market.peAvg) * price) / company.Dollars(pe), nil
}

func underValued(market *Market, pe float32) bool {
	return pe < market.peAvg
}

//returns details about the mainmarket
func mainMarket() *Market {
	return &Market{"main", 21}
}

//returns details about the junior market
func juniorMarket() *Market {
	return &Market{"junior", 15}
}

//structs

type Market struct {
	name  string
	peAvg float32
}

//calculates earnings per share
func eps(netIncome company.Dollars, shares float32) (company.Dollars, error) {
	if shares == 0 {
		return -1, fmt.Errorf("cant divide by 0")
	}
	return netIncome / company.Dollars(shares), nil
}

//calculates the book value of each share
func bvs(equity company.Dollars, shares float32) (company.Dollars, error) {
	if shares == 0 {
		return -1, fmt.Errorf("cant divide by 0")
	}
	return equity / company.Dollars(shares), nil
}

func netIncome(revenue, expenditure company.DollarSlice) company.Dollars {
	return revenue.Total() - expenditure.Total()
}

//calculates the difference between the assets and  liabilties
func equity(assets, liabilities company.DollarSlice) company.Dollars {
	return assets.Total() - liabilities.Total()
}

//takes a slice of dollars and returns the sum
//marked for depracation
func sumSlice(s company.DollarSlice) company.Dollars {
	var sum company.Dollars = 0
	for _, v := range s {
		sum += v
	}
	return sum
}

//calculates a financial ratio
func finRatio(a, b company.Dollars) (float32, error) {
	if b == 0 {
		return -1, fmt.Errorf("can't divide by 0")
	}
	return float32(a / b), nil
}
