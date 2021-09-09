package user

import (
	"github.com/howkyle/stockfolio-server/auth"
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

func (s service) Signup(us UserSignup) {
	hashPass := hashPass(us.Password)
	u := User{Email: us.Email, Username: us.UserName, Password: hashPass}

	s.repository.Create(u)
}

//uses authenticator to authenitcate user on login
func (s service) Login(username, password string) (string, error) {
	u, err := s.repository.Retrieve(username)
	if err != nil {
		return "", err
	}
	expected := s.authManager.NewCredentials(u.Username, u.Password)
	actual := s.authManager.NewCredentials(username, password)

	auth, err := s.authManager.Authenticate(actual, expected)
	return auth.Get(), nil
}

func hashPass(p string) string {
	h := p
	return h
}
