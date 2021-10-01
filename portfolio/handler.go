package portfolio

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

func GetPortfolioHandler(s Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userid := fmt.Sprintf("%v", r.Context().Value("sub"))

		uid, err := strconv.ParseUint(userid, 10, 64)
		if err != nil {
			http.Error(w, "something went wrong", http.StatusInternalServerError)
			return
		}
		p, err := s.Portfolio(uint(uid))
		if err != nil {
			if errors.Is(err, NotFound) {
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
