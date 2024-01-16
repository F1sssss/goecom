package database

import (
	"fmt"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

// Connect to the database
func Connect() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=Zippo123$ dbname=GO_DB port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	var err error

	once.Do(func() {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			fmt.Println("Error connecting to the database:", err)
		}
	})

	return db, err
}

// GetDB returns the database connection
func GetDB() *gorm.DB {
	return db
}
