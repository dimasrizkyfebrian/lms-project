// cmd/web/main.go
package main

import (
	"lms-project/backend/internal/handler"
	"lms-project/backend/internal/model"
	"lms-project/backend/internal/repository"
	"lms-project/backend/internal/service"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// --- Database Connection (sama seperti sebelumnya) ---
	dsn := "host=localhost user=postgres password=20022003 dbname=lms_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	db.AutoMigrate(&model.User{})
	log.Println("Database migration completed.")

	// --- Initialize Layers (sama seperti sebelumnya) ---
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	// --- HTTP Server Setup (using Gin) ---
	router := gin.Default()

	// --- 2. TERAPKAN MIDDLEWARE CORS ---
	// Konfigurasi ini memberitahu backend untuk mengizinkan
	// request dari frontend Anda (localhost:3000).
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"} // Izinkan origin frontend
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	router.Use(cors.New(config))

	// API v1 route group
	apiV1 := router.Group("/api/v1")
	{
		apiV1.POST("/register", userHandler.Register)
		apiV1.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "UP"})
		})
	}

	log.Println("Starting server on http://localhost:8080")
	router.Run(":8080")
}