package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/howkyle/stockfolio-server/analysis"
)

type server struct {
	port string
}

func Create(port string) server {
	return server{port: port}
}
func router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/analyze", analysis.AnalysisHandler()).Methods("POST")
	return r
}

func (s *server) Start() {
	fmt.Printf("starting server on port %s", s.port)
	log.Fatal(http.ListenAndServe(s.port, router()))
}
