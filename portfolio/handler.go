package portfolio

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func GetPortfolioHandler(s Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userid := r.Context().Value("sub")

		uid, ok := userid.(uint)
		log.Printf("user id %v  %v %v, %v", fmt.Sprintf("%T", userid), userid, uid, ok)
		if !ok {
			http.Error(w, "somthing went wrong", http.StatusInternalServerError)
			return
		}
		p, err := s.Portfolio(uid)
		if err != nil {
			http.Error(w, "somthing went wrong", http.StatusInternalServerError)
			return
		}
		body, err := json.Marshal(p)
		if err != nil {
			http.Error(w, "somthing went wrong", http.StatusInternalServerError)
			return
		}
		w.Header().Add("content-type", "application/json")
		fmt.Fprintf(w, "%v", body)
	}
}
func NewCompanyHandler(s Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Body == nil {
			http.Error(w, "empty request body", http.StatusBadRequest)
			return
		}

		var company Company

		err := json.NewDecoder(r.Body).Decode(&company)
		if err != nil {
			http.Error(w, "unable to decode request body", http.StatusBadRequest)
			return
		}
		_, err = s.AddCompany(company)
		if err != nil {
			log.Println(err)
			http.Error(w, "somthing went wrong", http.StatusInternalServerError)
			return
		}
		//add to db with user id as key
	}
}
