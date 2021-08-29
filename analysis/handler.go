package analysis

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/howkyle/stockfolio-server/domain/entity"
)

func AnalysisHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "method not allowed", http.StatusForbidden)
			return
		}

		if r.Body == nil {
			http.Error(w, "empty request body", http.StatusBadRequest)
			return
		}

		var company entity.Company
		// err := json.Unmarshal(r.Body, &company)
		err := json.NewDecoder(r.Body).Decode(&company)
		if err != nil {
			http.Error(w, "unable to decode request body", http.StatusBadRequest)
			return
		}

		result, err := Analyze(&company)
		response, _ := json.Marshal(result)
		w.Header().Add("content-type", "application/json")
		fmt.Fprintf(w, string(response))
	}

}
