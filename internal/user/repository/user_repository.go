package repository

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DbUserConfiguration() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=root dbname=todoapp port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}
