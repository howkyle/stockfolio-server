package analysis

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func AnalysisHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Body == nil {
			http.Error(w, "empty request body", http.StatusBadRequest)
			return
		}

		var company Company

		err := json.NewDecoder(r.Body).Decode(&company)
		if err != nil {
			log.Printf("unable to decode request body: %v", err)
			http.Error(w, "unable to decode request body", http.StatusBadRequest)
			return
		}

		result, err := Analyze(&company)
		if err != nil {
			http.Error(w, "unable to run analysis", http.StatusInternalServerError)
			return
		}

		response, _ := json.Marshal(result)
		w.Header().Add("content-type", "application/json")
		fmt.Fprintf(w, string(response))
	}

}
