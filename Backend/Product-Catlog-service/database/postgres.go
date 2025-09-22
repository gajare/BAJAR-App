package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func InitDB(databaseURL string) (*gorm.DB, error) {
	db, err := gorm.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	// Enable UUID extension
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")

	return db, nil
}
