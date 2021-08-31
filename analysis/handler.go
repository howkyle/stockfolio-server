package analysis

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/howkyle/stockfolio-server/portfolio"
)

func AnalysisHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Body == nil {
			http.Error(w, "empty request body", http.StatusBadRequest)
			return
		}

		var company portfolio.Company

		err := json.NewDecoder(r.Body).Decode(&company)
		if err != nil {
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
