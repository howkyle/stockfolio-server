package analysis

import (
	"fmt"
	"log"

	"github.com/howkyle/stockfolio-server/domain/entity"
	"github.com/howkyle/stockfolio-server/domain/types"
)

type Result struct {
	Pe            float32
	Pb            float32
	CurrentRatio  float32
	DeRatio       float32
	EarningsYield float32
	Roe           float32
	ProfitMargin  float32
	Eps           types.Dollars
	Bvs           types.Dollars
	NetIncome     types.Dollars
	TotalEquity   types.Dollars
}

func Analyze(c *entity.Company) (*Result, error) {
	netIncome := netIncome(c.Revenue, c.Expenditure)
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
	currentRatio, err := finRatio(c.CurrentAssets, c.CurrentLiabilities)

	if err != nil {
		log.Printf("unable to calculate current ratio: %v", err)
		return nil, err
	}

	deRatio, err := finRatio(c.CurrentLiabilities+c.LongLiabilities, totalEquity)
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
	profitMargin, err := finRatio(netIncome, c.Revenue)
	if err != nil {
		return nil, err

	}

	return &Result{pe, pb, currentRatio, deRatio, earningsYield, roe, profitMargin, eps, bookVal, netIncome, totalEquity}, nil
}

func PEMultiplePrice(market *Market, price types.Dollars, pe float32) (types.Dollars, error) {
	if pe == 0 {
		return -1, fmt.Errorf("cant divide by 0")
	}
	return (types.Dollars(market.peAvg) * price) / types.Dollars(pe), nil
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
func eps(netIncome types.Dollars, shares float32) (types.Dollars, error) {
	if shares == 0 {
		return -1, fmt.Errorf("cant divide by 0")
	}
	return netIncome / types.Dollars(shares), nil
}

//calculates the book value of each share
func bvs(equity types.Dollars, shares float32) (types.Dollars, error) {
	if shares == 0 {
		return -1, fmt.Errorf("cant divide by 0")
	}
	return equity / types.Dollars(shares), nil
}

func netIncome(revenue, expenditure types.Dollars) types.Dollars {
	return revenue - expenditure
}

//calculates the difference between the assets and  liabilties
func equity(assets, liabilities types.Dollars) types.Dollars {
	return assets - liabilities
}

//calculates a financial ratio
func finRatio(a, b types.Dollars) (float32, error) {
	if b == 0 {
		return -1, fmt.Errorf("can't divide by 0")
	}
	return float32(a / b), nil
}
