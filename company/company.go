package company

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Dollars float32

func (d Dollars) String() string {
	return fmt.Sprintf("$%g", d)
}

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
	FinancialReports []FinancialReport
}

type FinancialReport struct {
	gorm.Model
	CompanyID uint
	Shares    int
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
	//retrieves report using the report id
	Report(rid uint) (FinancialReport, error)
	//retrieves slice of reports using company id
	ReportsByCompany(cid uint) ([]FinancialReport, error)
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
	Report(rid uint) (FinancialReport, error)
	//retrieves a slice of financial reports belonging to a company
	ReportsByCompany(cid uint) ([]FinancialReport, error)
}
