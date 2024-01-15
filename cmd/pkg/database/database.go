package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connect to the database
func Connect() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=Zippo123$ dbname=GO_DB port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
