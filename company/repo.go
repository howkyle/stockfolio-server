package company

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func (r repository) AddCompany(c Company) (uint, error) {
	res := r.db.Create(&c)
	if res.Error != nil {
		log.Println(res.Error)
		return 0, fmt.Errorf("unable to create company: %v", res.Error)
	}
	return c.ID, nil
}

func (r repository) DeleteCompany(cid uint) error {
	res := r.db.Delete(&Company{}, cid)
	if res.Error != nil {
		log.Println(res.Error)
		return fmt.Errorf("unable to delete company: %v", res.Error)
	}
	return nil
}

//retrieves a company from the db using the company id
func (r repository) Company(cid uint) (Company, error) {
	var c Company
	res := r.db.First(&c, cid)
	if res.Error != nil {
		log.Println(res.Error)
		return Company{}, fmt.Errorf("unable to retrieve company: %w", res.Error)
	}
	return c, nil
}

//retrieves a slice of companies from the db using the portfolio id
func (r repository) Companies(pid uint) ([]Company, error) {
	var c []Company
	res := r.db.Where(Company{PortfolioID: pid}).Find(&c)
	if res.Error != nil {
		log.Println(res.Error)
		return nil, fmt.Errorf("unable to retrieve companies: %w", res.Error)

	}
	return c, nil
}

//inserts a new financial report to the db and return the id
func (r repository) AddReport(fr FinancialReport) (uint, error) {
	res := r.db.Create(&fr)
	if res.Error != nil {
		log.Println(res.Error)
		return 0, fmt.Errorf("unable to create report: %v", res.Error)
	}
	log.Printf("report added, id: %v", fr.ID)
	return fr.ID, nil
}

//retrieves a financial report from the db using the report id
func (r repository) Report(rid uint) (FinancialReport, error) {
	var fr FinancialReport
	res := r.db.First(&fr, rid)
	if res.Error != nil {
		log.Println((res.Error))
		return FinancialReport{}, res.Error
	}
	return fr, nil
}

//retrieves a slice financial report from the db using the company id
func (r repository) ReportsByCompany(cid uint) ([]FinancialReport, error) {
	var fr []FinancialReport
	res := r.db.Where(FinancialReport{CompanyID: cid}).Find(&fr)
	if res.Error != nil {
		log.Println((res.Error))
		return nil, res.Error
	}
	return fr, nil
}

//returns a new company repository
func NewRepository(database *gorm.DB) Repo {
	return &repository{database}
}
