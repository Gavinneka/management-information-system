package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectDB membuat koneksi ke database PostgreSQL dan mengembalikan instance *gorm.DB
func ConnectDB() (*gorm.DB, error) {
	// Ganti sesuai konfigurasi kamu
	dsn := "host=localhost user=postgres password=admin dbname=postgres port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("gagal koneksi ke database: %w", err)
	}

	log.Println("✅ Database terkoneksi!")
	return db, nil
}
