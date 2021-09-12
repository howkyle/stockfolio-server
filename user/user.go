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

type Repo interface {
	Create(user User)
	Retrieve(username string) (*User, error)
	Delete(id string)
}

type Service interface {
	Signup(u UserSignup) error
	Login(username, password string) (string, error)
}
