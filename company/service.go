package company

import (
	"fmt"
	"log"
	"time"
)

type AddCompany struct {
	PortfolioID uint
	Name        string
	Symbol      string
}

func (a AddCompany) Company() Company {
	return Company{PortfolioID: a.PortfolioID,
		Name: a.Name, Symbol: a.Symbol}
}

type AddReport struct {
	CompanyID          uint
	Shares             int
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
	return FinancialReport{
		CompanyID: a.CompanyID,
		Shares:    a.Shares,
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
}

type service struct {
	repo Repo
}

func (s service) AddCompany(c Company) (uint, error) {
	cid, err := s.repo.AddCompany(c)
	if err != nil {
		log.Println(err)
		return 0, fmt.Errorf("unable to add company: %v", err)
	}
	return cid, nil
}

func (s service) Company(cid uint) (Company, error) {
	c, err := s.repo.Company(cid)
	if err != nil {
		log.Println(err)
		return Company{}, err
	}
	return c, nil
}

func (s service) CompaniesByPortfolio(pid uint) ([]Company, error) {
	c, err := s.repo.Companies(pid)
	if err != nil {
		log.Printf("company service failed: %v", err)
		return nil, err
	}
	return c, nil
}

func (s service) AddReport(r FinancialReport) (uint, error) {

	rid, err := s.repo.AddReport(r)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return rid, nil
}

func (s service) Report(rid uint) (FinancialReport, error) {
	fr, err := s.repo.Report(rid)
	if err != nil {
		log.Println(err)
		return FinancialReport{}, err
	}
	return fr, nil
}

func (s service) ReportsByCompany(cid uint) ([]FinancialReport, error) {
	fr, err := s.repo.ReportsByCompany(cid)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return fr, nil
}

func CreateService(r Repo) service {
	return service{repo: r}
}
