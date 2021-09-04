package user

import (
	"github.com/howkyle/stockfolio-server/portfolio"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string
	Password  string
	Portfolio []portfolio.Portfolio
}

type Repo interface {
	Create(user *User)
	Retrieve(id string)
	Delete(id string)
}

type Service interface {
	Signup()
	Login() (*User, error)
}

type service struct {
	repository Repo
}

func Signup(u *User) {

}

func Login(id string) (*User, error) {
	return nil, nil
}
