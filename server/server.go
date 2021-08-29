package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/howkyle/stockfolio-server/analysis"
)

type server struct {
	port string
}

func Create(port string) server {
	return server{port: port}
}

func (s *server) Start() {
	http.HandleFunc("/analyze", analysis.AnalysisHandler())
	fmt.Printf("starting server on port %s", s.port)
	log.Fatal(http.ListenAndServe(s.port, nil))
}
