package analysis

import (
	"fmt"
	"log"

	"github.com/howkyle/stockfolio-server/portfolio"
)

type Company struct {
	Name   string
	Symbol string
	Market int32
	Shares int
	Price  portfolio.Dollars
	Earnings
	FinacialPosition
}

type FinacialPosition struct {
	CurrentAssets      []portfolio.Dollars
	CurrentLiabilities []portfolio.Dollars
	LongAssets         []portfolio.Dollars
	LongLiabilities    []portfolio.Dollars
}

type Earnings struct {
	Income      []portfolio.Dollars
	Expenditure []portfolio.Dollars
}
type Result struct {
	Pe            float32
	Pb            float32
	CurrentRatio  float32
	DeRatio       float32
	EarningsYield float32
	Roe           float32
	ProfitMargin  float32
	Eps           portfolio.Dollars
	Bvs           portfolio.Dollars
	NetIncome     portfolio.Dollars
	TotalEquity   portfolio.Dollars
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
	currentRatio, err := finRatio(sumSlice(c.CurrentAssets), sumSlice(c.CurrentLiabilities))

	if err != nil {
		log.Printf("unable to calculate current ratio: %v", err)
		return nil, err
	}

	deRatio, err := finRatio(sumSlice(c.CurrentLiabilities)+sumSlice(c.LongLiabilities), totalEquity)
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
	profitMargin, err := finRatio(netIncome, sumSlice(c.Income))
	if err != nil {
		return nil, err

	}

	return &Result{pe, pb, currentRatio, deRatio, earningsYield, roe, profitMargin, eps, bookVal, netIncome, totalEquity}, nil
}

func PEMultiplePrice(market *Market, price portfolio.Dollars, pe float32) (portfolio.Dollars, error) {
	if pe == 0 {
		return -1, fmt.Errorf("cant divide by 0")
	}
	return (portfolio.Dollars(market.peAvg) * price) / portfolio.Dollars(pe), nil
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
func eps(netIncome portfolio.Dollars, shares float32) (portfolio.Dollars, error) {
	if shares == 0 {
		return -1, fmt.Errorf("cant divide by 0")
	}
	return netIncome / portfolio.Dollars(shares), nil
}

//calculates the book value of each share
func bvs(equity portfolio.Dollars, shares float32) (portfolio.Dollars, error) {
	if shares == 0 {
		return -1, fmt.Errorf("cant divide by 0")
	}
	return equity / portfolio.Dollars(shares), nil
}

func netIncome(revenue, expenditure []portfolio.Dollars) portfolio.Dollars {
	return sumSlice(revenue) - sumSlice(expenditure)
}

//calculates the difference between the assets and  liabilties
func equity(assets, liabilities []portfolio.Dollars) portfolio.Dollars {
	sum_assets := sumSlice(assets)
	sum_liabilities := sumSlice(liabilities)

	return sum_assets - sum_liabilities
}

//takes a slice of dollars and returns the sum
func sumSlice(s []portfolio.Dollars) portfolio.Dollars {
	var sum portfolio.Dollars = 0
	for _, v := range s {
		sum += v
	}
	return sum
}

//calculates a financial ratio
func finRatio(a, b portfolio.Dollars) (float32, error) {
	if b == 0 {
		return -1, fmt.Errorf("can't divide by 0")
	}
	return float32(a / b), nil
}
