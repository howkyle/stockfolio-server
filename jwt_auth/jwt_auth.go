//implements Auth Interface using jwt
package jwt_auth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/howkyle/stockfolio-server/auth"
	"golang.org/x/crypto/bcrypt"
)

type userPassCredentials struct {
	principal interface{}
	password  string
}

func (u userPassCredentials) Principal() interface{} {
	return u.principal
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
		log.Printf("auth error: %v", err)
		return jwtAuth{}, fmt.Errorf("credentials not equal: %v", err)
	}
	log.Printf("authenticated '%s'", u.Principal())
	token, err := createToken(fmt.Sprint(u.Principal()), a.secret)
	if err != nil {
		log.Printf("token error: %v", err)
		return nil, fmt.Errorf("token creation failed: %v", err)
	}
	return jwtAuth{token}, nil
}

//takes principal and password and returns credential struct
func (a jwtAuthMan) NewCredentials(principal interface{}, password string) auth.Credentials {
	u := userPassCredentials{principal: principal, password: password}
	return u
}

func (a jwtAuthMan) Filter(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//check auth in requests
		cookie, err := r.Cookie("pyt")
		if err != nil {
			http.Error(w, "missing auth", http.StatusUnauthorized)
			return
		}
		sub, err := verifyToken(cookie.Value, a.secret)
		if err != nil {
			http.Error(w, "access not allowed", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), "sub", sub)

		h.ServeHTTP(w, r.WithContext(ctx))

	}
}

//creates a new instance of the jwt auth manager
func NewJWTAuth(secret string) jwtAuthMan {
	return jwtAuthMan{secret: secret}
}

//helpers

//compares passwords
func compare(a, b []byte) error {
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

//takes a token string and the server secret and parses and validates token and returns
//the subject i.e username
func verifyToken(t string, secret string) (uint, error) {
	//todo add more validations and checks
	log.Printf("validating token: %v", t)
	token, err := jwt.ParseWithClaims(t, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		log.Printf("parsing error: %v", err)
		return 0, fmt.Errorf("unable to parse token: %v", err)
	}
	c := token.Claims.(*jwt.StandardClaims)
	uid, err := strconv.ParseUint(c.Subject, 10, 64)
	if err != nil {
		log.Println(err)
		return 0, fmt.Errorf("unable to parse sub: %v", err)
	}
	return uint(uid), nil
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
