package user

import (
	"errors"
	"fmt"
	"log"

	"gorm.io/gorm"
)

var InvalidUser = errors.New("invalid user struct")
var CreationFailure = errors.New("unable to add user to repository")
var NotFound = errors.New("user not found")

func NewRepository(database *gorm.DB) Repo {
	return repository{database}
}

type repository struct {
	db *gorm.DB
}

func (r repository) Create(user User) (interface{}, error) {

	res := r.db.Create(&user)
	if res.Error != nil {
		log.Printf("unable to add user:%v", res.Error)
		return nil, CreationFailure
	}
	return user.ID, nil
}

//takes a struct with the criterial to to search
func (r repository) Retrieve(u interface{}) (User, error) {
	user := User{}
	res := r.db.Where(u).First(&user)
	if res.Error != nil {
		log.Println(res.Error)

		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return User{}, NotFound
		}
		return User{}, fmt.Errorf("unable to retrieve user:%w", res.Error)
	}
	return user, nil
}

func (r repository) Delete(id interface{}) error {
	return nil
}
