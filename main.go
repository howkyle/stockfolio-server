package main

import (
	"log"

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
	server := server.Create(config.Port, db)
	server.Start()
}

func initDB(connection string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(connection))
	if err != nil {
		panic("unable to connect to database")
	}
	log.Printf("connected to db: %v", db.Name())

	db.AutoMigrate(user.User{}, portfolio.Portfolio{}, portfolio.Company{})

	return db
}
