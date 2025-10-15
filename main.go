package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"APITest/database"
	"APITest/internal/users"
)


func main() {
	// Inisialisasi router Gin
	r := gin.Default()

	// Aktifkan CORS untuk frontend di port 5173
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: true,
	}))

	// Koneksi ke database
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Gagal koneksi database: %v", err)
	}

	// Jalankan migrasi otomatis
	if err := db.AutoMigrate(&users.User{}); err != nil {
		log.Fatalf("Gagal migrasi model: %v", err)
	}

	// Inisialisasi layer Repository → Service → Handler
	userRepo := users.NewRepository(db)
	userService := users.NewService(userRepo)
	userHandler := users.NewHandler(userService)

	// Registrasi route user
	users.RegisterRoutes(r, userHandler)

	// Jalankan server pada port 8080
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}
