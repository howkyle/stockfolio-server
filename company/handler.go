package company

import (
	"encoding/json"
	"log"
	"net/http"
)

func NewCompanyHandler(s Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Body == nil {
			http.Error(w, "empty request body", http.StatusBadRequest)
			return
		}

		var company AddCompany

		err := json.NewDecoder(r.Body).Decode(&company)
		if err != nil {
			http.Error(w, "unable to decode request body", http.StatusBadRequest)
			return
		}
		_, err = s.AddCompany(company.Company())
		if err != nil {
			log.Println(err)
			http.Error(w, "somthing went wrong", http.StatusInternalServerError)
			return
		}
		//add to db with user id as key
	}
}

func AddReportHandler(s Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Body == nil {
			http.Error(w, "empty request body", http.StatusBadRequest)
			return
		}

		var a AddReport
		err := json.NewDecoder(r.Body).Decode(&a)
		if err != nil {
			http.Error(w, "unable to decode request body", http.StatusBadRequest)
			return
		}

		_, err = s.AddReport(a.Report())
		if err != nil {
			log.Println(err)
			http.Error(w, "something went wrong", http.StatusInternalServerError)
			return
		}
	}
}
