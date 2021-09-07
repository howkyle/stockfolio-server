package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/howkyle/stockfolio-server/analysis"
	"github.com/howkyle/stockfolio-server/user"
	"gorm.io/gorm"
)

type server struct {
	port        string
	router      *mux.Router
	db          *gorm.DB
	userService user.Service
}

//creates server instance
func Create(port string, db *gorm.DB) server {

	s := server{port: port, db: db}
	s.configServices()
	s.configRouter()
	return s
}

// configures routes and returns router
func (s *server) configRouter() {
	r := mux.NewRouter()
	r.HandleFunc("/signup", user.SignUpHandler(s.userService)).Methods("POST")
	r.HandleFunc("/analyze", analysis.AnalysisHandler()).Methods("POST")
	s.router = r
}

func (s *server) configServices() {
	ur := user.NewRepository(s.db)
	s.userService = user.CreateService(ur)
}

//starts the server
func (s *server) Start() {
	fmt.Printf("starting server on port %s", s.port)
	log.Fatal(http.ListenAndServe(s.port, s.router))
}
