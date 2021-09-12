//implements Auth Interface using jwt
package jwt_auth

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/howkyle/stockfolio-server/auth"
	"golang.org/x/crypto/bcrypt"
)

type userPassCredentials struct {
	username, password string
}

func (u userPassCredentials) Principal() string {
	return u.username
}

func (u userPassCredentials) Hash() (string, error) {
	hash, err := hashPass(u.password)
	if err != nil {
		return "", fmt.Errorf("unable to hash password: %v", err)
	}
	return string(hash), nil
}

func (u userPassCredentials) Password() string {
	return u.password
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

func (a jwtAuthMan) Authenticate(u auth.Credentials, password string) (auth.Auth, error) {
	err := compare([]byte(u.Password()), []byte(password))
	if err != nil {
		log.Println("invalid credentials")
		return jwtAuth{}, fmt.Errorf("credentials not equal: %v", err)
	}
	log.Printf("authenticated '%s'", u.Principal())
	token, err := createToken(u.Principal(), a.secret)
	if err != nil {
		log.Printf("token creation failed: %v", err)
		return nil, fmt.Errorf("token creation failed: %v", err)
	}
	return jwtAuth{token}, nil
}

//takes username and password and returns credential struct
func (a jwtAuthMan) NewCredentials(username, password string) auth.Credentials {

	u := userPassCredentials{username: username, password: password}
	return u
}

func (a jwtAuthMan) CheckAuth() error {
	return nil
}

//creates a new instance of the jwt auth manager
func NewJWTAuth(secret string) jwtAuthMan {
	return jwtAuthMan{secret: secret}
}

//compares passwords
func compare(a, b []byte) error {
	log.Printf("comparing %v and %v", a, b)
	err := bcrypt.CompareHashAndPassword(a, b)
	if err != nil {
		return fmt.Errorf("password comparision failed: %v", err)
	}
	return nil
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

//hashes password using bcrypt
func hashPass(password string) ([]byte, error) {
	hp, err := bcrypt.GenerateFromPassword([]byte(password), 15)
	if err != nil {
		log.Printf("hash failed: %v", err)
		return nil, fmt.Errorf("unable to hash password: %v", err)
	}
	return hp, nil
}
