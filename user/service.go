//implements user service
package user

import (
	"fmt"
	"log"

	"github.com/howkyle/stockfolio-server/auth"
	"github.com/howkyle/stockfolio-server/portfolio"
)

type service struct {
	repository  Repo
	authManager auth.AuthManager
}

//creates a new instance of the user service
func CreateService(r Repo, auth auth.AuthManager) service {
	s := service{repository: r, authManager: auth}

	return s
}

func (s service) Signup(us UserSignup) error {
	cred := s.authManager.NewCredentials(us.UserName, us.Password)
	hash, err := cred.Hash()
	if err != nil {
		log.Println(err)
		return fmt.Errorf("unable to signup user: %v", err)
	}
	u := User{Email: us.Email, Username: us.UserName, Password: hash, Portfolio: portfolio.Portfolio{Title: fmt.Sprintf("%v's Portfolio", us.UserName)}}
	s.repository.Create(u)
	return nil
}

//uses authenticator to authenitcate user on login
func (s service) Login(username, password string) (string, error) {
	u, err := s.repository.Retrieve(username)
	if err != nil {
		return "", err
	}
	credentials := s.authManager.NewCredentials(u.ID, u.Password)
	auth, err := s.authManager.Authenticate(credentials, password)
	if err != nil {
		log.Printf("login failed: %v", err)
		return "", fmt.Errorf("login failed: %v", err)
	}
	return auth.Get(), nil
}
