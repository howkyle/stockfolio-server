package portfolio

import (
	"github.com/howkyle/stockfolio-server/company"
	"gorm.io/gorm"
)

type Portfolio struct {
	Title  string
	UserID uint
	gorm.Model
	Companies []company.Company
}

type Service interface {
	//retrieves a users portfolio
	Portfolio(userid uint) (Portfolio, error)
}

type Repo interface {
	//retrieves a portfolio associated with a user
	Get(userid uint) (Portfolio, error)
}
