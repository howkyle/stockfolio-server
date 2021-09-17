package main

import (
	"log"

	"github.com/howkyle/stockfolio-server/company"
	"github.com/howkyle/stockfolio-server/config"
	"github.com/howkyle/stockfolio-server/portfolio"
	"github.com/howkyle/stockfolio-server/server"
	"github.com/howkyle/stockfolio-server/user"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	config := config.Config()
	db := initDB(config.DB)
	server := server.Create(config.Port, db, config.Secret)
	server.Start()
}

func initDB(connection string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(connection))
	if err != nil {
		panic("unable to connect to database")
	}
	log.Printf("connected to db: %v", db.Name())

	err = db.AutoMigrate(user.User{}, portfolio.Portfolio{}, company.Company{}, company.FinancialReport{})
	if err != nil {
		panic("unable to run db migration: " + err.Error())
	}

	return db
}
