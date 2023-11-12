package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Database *gorm.DB
}

func NewDatabase() (*Database, error) {
	dsn := "host=localhost user=postgres password=shalin123 dbname=users-gorm port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return &Database{
		Database: db,
	}, nil
}

// Creates the table for a model
func (db *Database) InitialMigration(entity interface{}) {
	db.Database.AutoMigrate(&entity)
}
