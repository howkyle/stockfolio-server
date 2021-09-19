package company

import (
	"log"
	"time"

	"gorm.io/gorm"
)

type Dollars float32
type DollarSlice []Dollars

func (d DollarSlice) Total() Dollars {
	var sum Dollars = 0
	for _, v := range d {
		sum += v
	}
	return sum
}

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

//

type Repo interface {
	//inserts a new company linked to a portfolio into the table
	AddCompany(c Company) (uint, error)
	//takes a company id and removes the company from the db
	DeleteCompany(cid uint) error
	//retrieves a company from the db by id
	Company(cid uint) (Company, error)
	//retrieves a slice of companies from the db
	Companies(pid uint) ([]Company, error)
	//inserts a new report associated with a company
	AddReport(fr FinancialReport) (uint, error)
	//retrieves report associated with a company
	GetReport(rid uint) (FinancialReport, error)
}

type Service interface {
	//adds a new company to the portfolio and returns the company id
	AddCompany(c Company) (uint, error)
	//retrieves a company in the users porfolio
	Company(cid uint) (Company, error)
	//retrieves a slice of companies belonging to a portfolio
	CompaniesByPortfolio(pid uint) ([]Company, error)
	//adds a financial report for a company in the portfolio
	AddReport(r FinancialReport) (uint, error)
	//retrieves a financial report associated with a company in the portfolio
	GetReport(rid uint) (FinancialReport, error)
}

//UseCase

type AddCompany struct {
	PortfolioID uint
	Name        string
	Symbol      string
	Shares      int
}

func (a AddCompany) Company() Company {
	return Company{PortfolioID: a.PortfolioID,
		Name: a.Name, Symbol: a.Symbol, Shares: a.Shares}
}

type AddReport struct {
	gorm.Model
	CompanyID          uint
	Year               time.Time
	Quarter            int
	Price              Dollars
	CurrentAssets      DollarSlice
	CurrentLiabilities DollarSlice
	LongAssets         DollarSlice
	LongLiabilities    DollarSlice
	Income             DollarSlice
	Expenditure        DollarSlice
}

func (a AddReport) Report() FinancialReport {
	r := FinancialReport{
		CompanyID: a.CompanyID,
		Year:      a.Year,
		Quarter:   a.Quarter,
		Price:     a.Price,
		Earnings:  Earnings{Income: a.Income.Total(), Expenditure: a.Expenditure.Total()},
		FinacialPosition: FinacialPosition{
			CurrentAssets:      a.CurrentAssets.Total(),
			CurrentLiabilities: a.CurrentLiabilities.Total(),
			LongAssets:         a.LongAssets.Total(),
			LongLiabilities:    a.LongLiabilities.Total(),
		},
	}
	log.Println(r)
	return r
}
