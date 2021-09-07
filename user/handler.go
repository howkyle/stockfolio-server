package user

import (
	"encoding/json"
	"log"
	"net/http"
)

func SignUpHandler(s Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("signin up user")
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

	}
}
