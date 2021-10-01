package user

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/howkyle/stockfolio-server/portfolio"
	"github.com/howkyle/uman"
)

func SignUpHandler(u uman.UserManager, a uman.AuthManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body UserSignup
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			http.Error(w, "unable to read payload", http.StatusBadRequest)
			return
		}
		hashedPass, err := uman.NewUserPassCredentials(body.UserName, body.Password).Hash()
		if err != nil {
			http.Error(w, "unable to hash credentials", http.StatusInternalServerError)
			return
		}

		ut := User{Username: body.UserName, Password: hashedPass, Portfolio: portfolio.Portfolio{Title: fmt.Sprintf("%v's Portfolio", body.UserName)}}

		_, err = u.Create(ut)
		if err != nil {
			log.Println(err)
			http.Error(w, "unable to create user", http.StatusInternalServerError)
		}

	}
}

func LoginHandler(u uman.UserManager, a uman.AuthManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		login := struct{ UserName, Password string }{}
		err := json.NewDecoder(r.Body).Decode(&login)
		if err != nil {
			log.Printf("unable to decode request body:%v", err)
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}

		user, err := u.Retrieve(User{Username: login.UserName})
		if err != nil {
			log.Println(err)
			http.Error(w, "failed login", http.StatusUnauthorized)
			return
		}

		cred := uman.NewUserPassCredentials(fmt.Sprintf("%v", user.GetID()), user.GetPassword())
		auth, err := a.Authenticate(cred, login.Password)
		if err != nil {
			log.Println(err)
			http.Error(w, "failed login", http.StatusUnauthorized)
			return
		}

		cookie := http.Cookie{Name: "pyt", Value: auth.Auth(), Domain: "localhost"}
		http.SetCookie(w, &cookie)
	}
}

func LogOutHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{Name: "pyt", Expires: time.Now()})
	}
}
