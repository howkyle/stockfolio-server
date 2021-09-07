package user

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

func Login(id string) (*User, error) {
	return nil, nil
}
