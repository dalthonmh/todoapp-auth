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
			log.Println("✅ Conectado a la base de datos")
			return db
		}

		log.Printf("Intento %d: error al conectar a la base de datos: %v", i+1, err)
		time.Sleep(3 * time.Second)
	}

	log.Fatal("❌ No se pudo conectar a la base de datos después de múltiples intentos")
	return nil
}
