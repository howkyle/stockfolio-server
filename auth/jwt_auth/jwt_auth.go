//implements Auth Interface using jwt
package jwt_auth

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/howkyle/stockfolio-server/auth"
)

type userPassCredentials struct {
	Username, Password string
}

func (u userPassCredentials) Principal() string {
	return u.Username
}

type jwtAuth struct {
	access_token string
}

type jwtAuthMan struct {
	secret string
}

func (j jwtAuth) Get() string {
	return j.access_token
}

func (a jwtAuthMan) Authenticate(actual, expected auth.Credentials) (auth.Auth, error) {
	if !compare(actual, expected) {
		log.Println("invalid credentials")
		return jwtAuth{}, fmt.Errorf("credentials not equal")
	}
	log.Printf("authenticated '%s'", expected.Principal())
	token, err := createToken(expected.Principal(), a.secret)
	if err != nil {
		log.Printf("token creation failed: %v", err)
		return nil, fmt.Errorf("token creation failed: %v", err)
	}
	return jwtAuth{token}, nil
}

func (a jwtAuthMan) NewCredentials(username, password string) auth.Credentials {
	u := userPassCredentials{Username: username, Password: password}
	return u
}

func (a jwtAuthMan) CheckAuth() error {
	return nil
}

//creates a new instance of the jwt auth manager
func NewJWTAuth(secret string) jwtAuthMan {
	return jwtAuthMan{secret: secret}
}

//compares credential structs
func compare(a auth.Credentials, b auth.Credentials) bool {
	return a == b
}

//creates jwt using subject and secret, returns signed string
func createToken(subject string, secret string) (string, error) {
	claims := jwt.StandardClaims{
		Subject:   subject,
		ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		Issuer:    "localhost",
		IssuedAt:  time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ts, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", fmt.Errorf("unable to sign token: %v", err)
	}
	return ts, nil
}
