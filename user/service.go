package user

import "fmt"

type service struct {
	repository Repo
}

func CreateService(r Repo) service {
	s := service{repository: r}

	return s
}

func (s service) Signup(us UserSignup) {
	u := User{Email: us.Email, Username: us.UserName, Password: us.Password}

	s.repository.Create(u)
}

func (s service) Login(username, password string) error {
	u, err := s.repository.Retrieve(username)
	if err != nil {
		return err
	}
	fmt.Println(u)
	return nil
}
