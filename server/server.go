package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/howkyle/stockfolio-server/analysis"
	"github.com/howkyle/stockfolio-server/auth"
	"github.com/howkyle/stockfolio-server/auth/jwt_auth"
	"github.com/howkyle/stockfolio-server/portfolio"
	"github.com/howkyle/stockfolio-server/user"
	"gorm.io/gorm"
)

type server struct {
	port             string
	router           *mux.Router
	db               *gorm.DB
	secret           string
	userService      user.Service
	portfolioService portfolio.Service
	authManager      auth.AuthManager
}

//creates server instance
func Create(port string, db *gorm.DB, secret string) server {

	s := server{port: port, db: db, secret: secret}
	s.configServices()
	s.configRouter()
	return s
}

// configures routes and returns router
func (s *server) configRouter() {
	r := mux.NewRouter()
	r.HandleFunc("/login", user.LoginHandler(s.userService)).Methods("POST")
	r.HandleFunc("/signup", user.SignUpHandler(s.userService)).Methods("POST")
	r.HandleFunc("/portfolio", s.authManager.Filter(portfolio.GetPortfolioHandler(s.portfolioService))).Methods("GET")
	r.HandleFunc("/portfolio/add", s.authManager.Filter(portfolio.NewCompanyHandler(s.portfolioService))).Methods("POST")
	r.HandleFunc("/analyze", s.authManager.Filter(analysis.AnalysisHandler())).Methods("POST")
	s.router = r
}

func (s *server) configServices() {
	ur := user.NewRepository(s.db)
	pr := portfolio.NewRepository(s.db)
	s.authManager = jwt_auth.NewJWTAuth(s.secret)
	s.userService = user.CreateService(ur, s.authManager)
	s.portfolioService = portfolio.CreateService(pr)
}

//starts the server
func (s *server) Start() {
	fmt.Printf("starting server on port %s", s.port)
	log.Fatal(http.ListenAndServe(s.port, s.router))
}
