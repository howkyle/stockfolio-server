package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/howkyle/stockfolio-server/analysis"
	"gorm.io/gorm"
)

type server struct {
	port string
	db   *gorm.DB
}

//creates server instance
func Create(port string, db *gorm.DB) server {
	return server{port, db}
}

// configures routes and returns router
func router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/analyze", analysis.AnalysisHandler()).Methods("POST")
	return r
}

//starts the server
func (s *server) Start() {
	fmt.Printf("starting server on port %s", s.port)
	log.Fatal(http.ListenAndServe(s.port, router()))
}
