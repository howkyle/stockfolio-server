package user

import (
	"fmt"
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

func (r repository) Retrieve(username string) (*User, error) {
	log.Printf("retrieving user '%s'", username)
	user := User{}
	res := r.db.Where(&User{Username: username}).First(&user)
	if res.Error != nil {
		log.Printf("unable to retrieve user '%s': %v", username, res.Error)
		return nil, fmt.Errorf("unable to retrieve user:%v", res.Error)
	}

	return &user, nil
}

func (r repository) Delete(id string) {

}
