package portfolio

import (
	"gorm.io/gorm"
)

type Dollars float32

type Portfolio struct {
	Title  string
	UserID uint
	gorm.Model
	Companies []Company
}

type Service interface {
	//retrieves a users portfolio
	Portfolio(userid uint) (Portfolio, error)
	//adds a new company to the portfolio and returns the company id
	AddCompany(c Company) (uint, error)
	//retrieves a company in the users porfolio
	Company(cid uint) (Company, error)
	//adds a financial report for a company in the portfolio
	AddReport(cid uint, r FinancialReport) (uint, error)
	//retrieves a financial report associated with a company in the portfolio
	GetReport(rid uint) (FinancialReport, error)
}

type Repo interface {
	//retrieves a portfolio associated with a user
	Get(userid uint) (Portfolio, error)
	//inserts a new company linked to a portfolio into the table
	AddCompany(c Company) (uint, error)
	//takes a company id and removes the company from the db
	DeleteCompany(cid uint) error
	//retrieves a company from the db by id
	Company(cid uint) (Company, error)
	//inserts a new report associated with a company
	AddReport(fr FinancialReport) (uint, error)
	//retrieves report associated with a company
	GetReport(rid uint) (FinancialReport, error)
}
