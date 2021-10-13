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
			log.Println(err)
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
		auth, err := s.Signin(login.User())
		if err != nil {
			http.Error(w, "failed authentication", http.StatusUnauthorized)
			return
		}
		cookie, ok := auth.(http.Cookie)
		if !ok {
			http.Error(w, "failed to set auth", http.StatusInternalServerError)
			return
		}
		http.SetCookie(w, &cookie)
	}
}

func LogOutHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{Name: "pyt", Expires: time.Now()})
	}
}
