package db

import (
	"ge-rest-api/src/config"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

func Connect() *gorm.DB {

	db, err := gorm.Open(postgres.Open(config.GetDbUrl()), &gorm.Config{})

	config.CheckError(err, "Error connecting to database")

	return db
}

func Close(db *gorm.DB) {
	dbSQL, _ := db.DB()
	close := dbSQL.Close()

	config.CheckError(close, "Error closing database connection")

	dbSQL.Close()

}
