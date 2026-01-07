// utils/db.go
package utils

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectWithRetry() *gorm.DB {
	dsn := os.Getenv("DB_DSN")
	var db *gorm.DB
	var err error

	for i := 0; i < 10; i++ {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			log.Println("Connected to the database successfully")
			return db
		}

		log.Printf("Attempt %d: error connecting to the database: %v", i+1, err)
		time.Sleep(3 * time.Second)
	}

	log.Fatal("Could not connect to the database after multiple attempts")
	return nil
}
