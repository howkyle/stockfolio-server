package user

import (
	"gorm.io/gorm"
)

// func newRepo() {
// 	db, err := gorm.Open( mysql.Open("server=localhost;port=3306;database=stockfolio;uid=dev;password=password"))
// 	if(err!=nil){
// 		panic("unable to connect to database")
// 	}
// 	db.AutoMigrate(User{})
// }

func NewRepository(database *gorm.DB) Repo {
	return &repository{database}
}

type repository struct {
	db *gorm.DB
}

func (r *repository) Create(user *User) {

}

func (r *repository) Retrieve(id string) {

}

func (r *repository) Delete(id string) {

}
