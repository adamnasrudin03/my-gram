package database

import (
	"fmt"
	"log"

	"github.com/adamnasrudin03/my-gram/app/configs"
	"github.com/adamnasrudin03/my-gram/app/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// SetupDbConnection is creating a new connection to our database
func SetupDbConnection() *gorm.DB {
	configs := configs.GetInstance()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		configs.Dbconfig.Host,
		configs.Dbconfig.Username,
		configs.Dbconfig.Password,
		configs.Dbconfig.Dbname,
		configs.Dbconfig.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to create a connection to database")
	}

	if configs.Dbconfig.DebugMode {
		db = db.Debug()
	}

	if configs.Dbconfig.DbIsMigrate {
		//auto migration entity db
		db.AutoMigrate(
			&entity.User{},
			&entity.Photo{},
			&entity.SocialMedia{},
			&entity.Comment{},
		)
	}

	log.Println("Connection Database Success!")
	return db
}

// CloseDbConnection method is closing a connection between your app and your db
func CloseDbConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}

	dbSQL.Close()
}
