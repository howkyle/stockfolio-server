package user

import (
	"log"

	"gorm.io/gorm"
)

func NewRepository(database *gorm.DB) Repo {
	return &repository{database}
}

type repository struct {
	db *gorm.DB
}

func (r repository) Create(user User) {
	res := r.db.Create(&user)
	if res.Error != nil {
		log.Printf("unable to add user:%v", res.Error)
		return
	}

	log.Printf("user added")

}

func (r repository) Retrieve(id string) {

}

func (r repository) Delete(id string) {

}
