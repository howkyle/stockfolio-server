package analysis

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//handles requests for quick company analysis
func QuickAnalysisHandler(s Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var company QuickAnalysis
		err := decodeBody(r, &company)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		result, err := s.Analyze(company.ToFinancialReport())
		if err != nil {
			http.Error(w, "unable to run analysis", http.StatusInternalServerError)
			return
		}

		response, err := createResponse(result)
		if err != nil {
			http.Error(w, "unable to create response", http.StatusInternalServerError)
		}

		writeJSON(w, response)
	}

}

//handles requests to analyze a stored financial report
func ReportAnalysisHandler(s Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var company ReportAnalysis

		err := decodeBody(r, &company)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		result, err := s.Analyze(company.ToFinancialReport())
		if err != nil {
			http.Error(w, "unable to run analysis", http.StatusInternalServerError)
			return
		}

		response, err := createResponse(result)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		writeJSON(w, response)
	}

}

func decodeBody(r *http.Request, v interface{}) error {
	if r.Body == nil {
		return fmt.Errorf("empty request body")
	}

	err := json.NewDecoder(r.Body).Decode(&v)
	if err != nil {
		log.Printf("unable to decode request body: %v", err)
		return fmt.Errorf("unable to decode request body")
	}

	return nil
}

func createResponse(body interface{}) (string, error) {
	response, err := json.Marshal(body)
	if err != nil {
		return "", fmt.Errorf("unable to create response: %v", err)
	}
	return string(response), nil
}

func writeJSON(w http.ResponseWriter, v string) {
	w.Header().Add("content-type", "application/json")
	fmt.Fprintf(w, v)
}
