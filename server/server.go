package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/howkyle/stockfolio-server/analysis"
	"github.com/howkyle/stockfolio-server/auth"
	"github.com/howkyle/stockfolio-server/company"
	"github.com/howkyle/stockfolio-server/jwt_auth"
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
	companyService   company.Service
	analysisService  analysis.Service
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
	r.HandleFunc("/report", s.authManager.Filter(company.AddReportHandler(s.companyService))).Methods("POST")
	r.HandleFunc("/report/{rid}", s.authManager.Filter(company.GetReportHandler(s.companyService))).Methods("GET")
	r.HandleFunc("/portfolio", s.authManager.Filter(portfolio.GetPortfolioHandler(s.portfolioService))).Methods("GET")
	r.HandleFunc("/portfolio/{pid}/assets", s.authManager.Filter(company.CompaniesByPortfolioHandler(s.companyService))).Methods("GET")
	r.HandleFunc("/company", s.authManager.Filter(company.NewCompanyHandler(s.companyService))).Methods("POST")
	r.HandleFunc("/company/{cid}/reports", s.authManager.Filter(company.GetReportsByCompanyHandler(s.companyService))).Methods("GET")
	r.HandleFunc("/analysis/quick", s.authManager.Filter(analysis.QuickAnalysisHandler(s.analysisService))).Methods("POST")
	r.HandleFunc("/analysis/report", s.authManager.Filter(analysis.ReportAnalysisHandler(s.analysisService))).Methods("POST")

	s.router = r
}

func (s *server) configServices() {
	ur := user.NewRepository(s.db)
	pr := portfolio.NewRepository(s.db)
	cr := company.NewRepository(s.db)
	s.authManager = jwt_auth.NewJWTAuth(s.secret)
	s.userService = user.CreateService(ur, s.authManager)
	s.companyService = company.CreateService(cr)
	s.portfolioService = portfolio.CreateService(pr)
	s.analysisService = analysis.CreateService()
}

//starts the server
func (s *server) Start() {
	fmt.Printf("starting server on port %s", s.port)
	log.Fatal(http.ListenAndServe(s.port, s.router))
}
