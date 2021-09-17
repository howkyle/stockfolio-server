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

func (r repository) Company(cid uint) (Company, error) {
	var c Company
	res := r.db.First(&c, cid)
	if res.Error != nil {
		log.Println(res.Error)
		return Company{}, fmt.Errorf("unable to retrieve company: %v", res.Error)
	}
	return c, nil
}

func (r repository) AddReport(fr FinancialReport) (uint, error) {
	res := r.db.Create(&fr)
	if res.Error != nil {
		log.Println(res.Error)
		return 0, fmt.Errorf("unable to create report: %v", res.Error)
	}
	log.Printf("report added, id: %v", fr.ID)
	return fr.ID, nil
}

func (r repository) GetReport(rid uint) (FinancialReport, error) {
	var fr FinancialReport
	res := r.db.First(&fr, rid)
	if res.Error != nil {
		log.Println((res.Error))
		return FinancialReport{}, res.Error
	}
	return fr, nil
}

func NewRepository(database *gorm.DB) Repo {
	return &repository{database}
}
