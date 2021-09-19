package company

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func CompaniesByPortfolioHandler(s Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//get user id
		params := mux.Vars(r)
		pid := params["pid"]
		if pid == "" {
			http.Error(w, "missing param: portfolio id", http.StatusBadRequest)
			return
		}

		id, err := strconv.ParseUint(pid, 10, 64)
		if err != nil {
			http.Error(w, "invalid param: portfolio id", http.StatusBadRequest)
			return
		}

		c, err := s.CompaniesByPortfolio(uint(id))
		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusInternalServerError)
		}

		response, err := json.Marshal(c)
		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusInternalServerError)
		}

		w.Header().Add("content-type", "application/json")
		fmt.Fprint(w, string(response))
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

func GetReportHandler(s Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		rid := params["rid"]
		if rid == "" {
			http.Error(w, "invalid param: report id", http.StatusBadRequest)
			return
		}
		id, err := strconv.ParseUint(rid, 10, 64)
		if err != nil {
			log.Println(err)
			http.Error(w, "invalid param: report id", http.StatusBadRequest)
			return
		}
		result, err := s.GetReport(uint(id))
		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
		response, err := json.Marshal(result)
		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusInternalServerError)
		}

		w.Header().Add("content-type", "application/json")

		fmt.Fprint(w, string(response))
	}
}

func GetReportsByCompanyHandler(s Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		cid := params["cid"]
		if cid == "" {
			http.Error(w, "invalid param: company id", http.StatusBadRequest)
			return
		}
		id, err := strconv.ParseUint(cid, 10, 64)
		if err != nil {
			log.Println(err)
			http.Error(w, "invalid param: company id", http.StatusBadRequest)
			return
		}
		result, err := s.GetReportsByCompany(uint(id))
		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
		response, err := json.Marshal(result)
		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusInternalServerError)
		}

		w.Header().Add("content-type", "application/json")

		fmt.Fprint(w, string(response))
	}
}
