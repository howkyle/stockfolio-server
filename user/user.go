package user

import (
	"github.com/howkyle/stockfolio-server/portfolio"
	"github.com/howkyle/uman"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email     string
	Username  string
	Password  string
	Portfolio portfolio.Portfolio
}

func (u User) GetID() interface{} {
	return u.ID
}
func (u User) GetUsername() string {
	return u.Username
}
func (u User) GetEmail() string {
	return u.Email
}

func (u User) GetPassword() string {
	return u.Password
}

type UserSignup struct {
	UserName string
	Email    string
	Password string
}

type Repo interface {
	Create(user uman.User) (interface{}, error)
	Retrieve(u interface{}) (uman.User, error)
	Delete(id interface{}) error
}
