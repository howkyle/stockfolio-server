package user

import (
	"github.com/howkyle/stockfolio-server/portfolio"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email     string
	Username  string
	Password  string
	Portfolio portfolio.Portfolio
}

type UserSignup struct {
	UserName string
	Email    string
	Password string
}

type Service interface {
	Register(u User) (interface{}, error)
	Signin(u User) (interface{}, error)
}
type Repo interface {
	Create(user User) (interface{}, error)
	Retrieve(u interface{}) (User, error)
	Delete(id interface{}) error
}
