package jwt_auth

import (
	"fmt"
	"log"

	"github.com/howkyle/stockfolio-server/auth"
)

type userPassCredentials struct {
	Username, Password string
}

type jwtAuth struct {
	access_token string
}

type jwtAuthMan struct {
}

func (j jwtAuth) Get() string {
	return j.access_token
}

func (a jwtAuthMan) Authenticate(actual, expected auth.Credentials) (auth.Auth, error) {
	if !compare(actual, expected) {
		log.Println("invalid credentials")
		return jwtAuth{}, fmt.Errorf("credentials not equal")
	}

	return jwtAuth{"testtoken"}, nil
}

func (a jwtAuthMan) NewCredentials(username, password string) auth.Credentials {
	u := userPassCredentials{Username: username, Password: password}
	return u
}

func (a jwtAuthMan) CheckAuth() error {
	return nil
}

func NewJWTAuth() jwtAuthMan {
	return jwtAuthMan{}
}

func compare(a auth.Credentials, b auth.Credentials) bool {
	return a == b
}
