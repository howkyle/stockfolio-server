package portfolio

import (
	"errors"
	"log"

	"gorm.io/gorm"
)

var NotFound = errors.New("portfolio not found")

type repository struct {
	db *gorm.DB
}

func (r repository) Get(userid uint) (Portfolio, error) {
	var p Portfolio
	res := r.db.Where(&Portfolio{UserID: userid}).First(&p)
	if res.Error != nil {
		log.Println(res.Error)

		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return Portfolio{}, NotFound
		}
		return Portfolio{}, res.Error
	}
	return p, nil
}

func NewRepository(database *gorm.DB) Repo {
	return repository{database}
}
