package portfolio

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/howkyle/stockfolio-server/cust_error"
)

func GetPortfolioHandler(s Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userid := r.Context().Value("sub")

		uid, ok := userid.(uint)
		if !ok {
			http.Error(w, "something went wrong", http.StatusInternalServerError)
			return
		}
		p, err := s.Portfolio(uid)
		if err != nil {
			if errors.Is(err, cust_error.NotFound) {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		body, err := json.Marshal(p)
		if err != nil {
			http.Error(w, "somthing went wrong", http.StatusInternalServerError)
			return
		}
		w.Header().Add("content-type", "application/json")
		fmt.Fprintf(w, "%s", string(body))
	}

}
