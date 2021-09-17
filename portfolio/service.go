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
		return Portfolio{}, fmt.Errorf("unable to retrieve portfolio: %w", err)
	}
	return p, nil
}

func CreateService(r Repo) service {
	return service{repo: r}
}
