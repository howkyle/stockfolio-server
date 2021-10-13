package user

import (
	"fmt"
	"log"

	"github.com/howkyle/authman"
	"github.com/howkyle/stockfolio-server/portfolio"
)

//todo redo error handling

type service struct {
	repo    Repo
	authman authman.AuthManager
}

func (s service) Register(u User) (interface{}, error) {
	hashedPass, err := authman.NewUserPassCredentials(u.Username, u.Password).Hash()
	if err != nil {
		// http.Error(w, "unable to hash credentials", http.StatusInternalServerError)
		// return
	}

	// ut := User{Username: body.UserName, Password: hashedPass, Portfolio: portfolio.Portfolio{Title: fmt.Sprintf("%v's Portfolio", body.UserName)}}
	u.Password = hashedPass
	id, err := s.repo.Create(u)
	if err != nil {
		// log.Println(err)
		// http.Error(w, "unable to create user", http.StatusInternalServerError)
	}
	return id, nil
}

func (s service) Signin(u User) (interface{}, error) {
	user, err := s.repo.Retrieve(u)
	if err != nil {
		// log.Println(err)
		// http.Error(w, "failed login", http.StatusUnauthorized)
		// return
	}

	cred := authman.NewUserPassCredentials(fmt.Sprintf("%v", user.ID), user.Password)
	auth, err := s.authman.Authenticate(cred, u.Password)
	if err != nil {
		log.Println(err)
		// http.Error(w, "failed login", http.StatusUnauthorized)
		// return
	}

	return auth.AsCookie(), nil
}

func NewService(repo Repo, authman authman.AuthManager) Service {
	return service{repo: repo, authman: authman}
}

type UserRegistration struct {
	Username string
	Email    string
	Password string
}

func (u UserRegistration) User() User {
	return User{Username: u.Username, Password: u.Password, Email: u.Email, Portfolio: portfolio.Portfolio{Title: fmt.Sprintf("%v's Portfolio", u.Username)}}
}
