package user

import (
	"encoding/json"
	"log"
	"net/http"
)

func SignUpHandler(s Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("signing up user")
		var body UserSignup
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			http.Error(w, "unable to read payload", http.StatusBadRequest)
		}

		s.Signup(body)
	}
}

func LoginHandler(s Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("logging in user")
		login := struct{ UserName, Password string }{}
		err := json.NewDecoder(r.Body).Decode(&login)
		if err != nil {
			log.Printf("unable to decode request body:%v", err)
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
		token, err := s.Login(login.UserName, login.Password)
		if err != nil {
			log.Printf("unable to login:%v", err)
			http.Error(w, "failed login", http.StatusUnauthorized)
			return

		}
		cookie := http.Cookie{Name: "pyt", Value: token, Domain: "localhost"}
		http.SetCookie(w, &cookie)
	}
}
