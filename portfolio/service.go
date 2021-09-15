package portfolio

import (
	"fmt"
	"log"
)

type service struct {
	repo Repo
}

func (s service) Portfolio(userid uint) (Portfolio, error) {
	p, err := s.repo.Get(userid)
	if err != nil {
		log.Println(err)
		return Portfolio{}, fmt.Errorf("unable to retrieve users portfolio: %v", err)
	}
	return p, nil
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

func (s service) AddReport(cid uint, r FinancialReport) (uint, error) {
	r.CompanyID = cid
	rid, err := s.repo.AddReport(r)
	if err != nil {
		return 0, err
	}
	return rid, nil
}

func (s service) GetReport(rid uint) (FinancialReport, error) {
	fr, err := s.repo.GetReport(rid)
	if err != nil {
		log.Println(err)
		return FinancialReport{}, err
	}
	return fr, nil
}

func CreateService(r Repo) service {
	return service{repo: r}
}
