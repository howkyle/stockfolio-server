package portfolio

import (
	"errors"
	"log"

	"github.com/howkyle/stockfolio-server/cust_error"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func (r repository) Get(userid uint) (Portfolio, error) {
	var p Portfolio
	res := r.db.Where(&Portfolio{UserID: userid}).First(&p)
	if res.Error != nil {
		log.Println(res.Error)

		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return Portfolio{}, cust_error.NotFound
		}
		return Portfolio{}, res.Error
	}
	return p, nil
}

func NewRepository(database *gorm.DB) Repo {
	return &repository{database}
}
