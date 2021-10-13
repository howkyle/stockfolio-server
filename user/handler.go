package user

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func SignUpHandler(s Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body UserRegistration
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			http.Error(w, "unable to read payload", http.StatusBadRequest)
			return
		}

		_, err = s.Register(body.User())
		if err != nil {
			log.Println(err)
			http.Error(w, "unable to create user", http.StatusInternalServerError)
		}
	}
}

func LoginHandler(s Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var login UserRegistration
		err := json.NewDecoder(r.Body).Decode(&login)
		if err != nil {
			log.Printf("unable to decode request body:%v", err)
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
		auth, err := s.Signin(login.User())
		if err != nil {

		}
		// user, err := u.Retrieve(User{Username: login.UserName})
		// if err != nil {
		// 	log.Println(err)
		// 	http.Error(w, "failed login", http.StatusUnauthorized)
		// 	return
		// }

		// cred := uman.NewUserPassCredentials(fmt.Sprintf("%v", user.GetID()), user.GetPassword())
		// auth, err := a.Authenticate(cred, login.Password)
		// if err != nil {
		// 	log.Println(err)
		// 	http.Error(w, "failed login", http.StatusUnauthorized)
		// 	return
		// }

		// cookie := http.Cookie{Name: "pyt", Value: auth.Auth(), Domain: "localhost"}
		cookie, ok := auth.(http.Cookie)
		if !ok {

		}
		http.SetCookie(w, &cookie)
	}
}

func LogOutHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{Name: "pyt", Expires: time.Now()})
	}
}
