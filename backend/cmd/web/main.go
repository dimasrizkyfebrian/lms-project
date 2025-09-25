// cmd/web/main.go
package main

import (
	"lms-project/backend/internal/handler"
	"lms-project/backend/internal/middleware"
	"lms-project/backend/internal/model"
	"lms-project/backend/internal/repository"
	"lms-project/backend/internal/service"
	"log"
	"net/http"
	"os"
	"time"

	ratelimit "github.com/JGLTechnologies/gin-rate-limit"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func keyFunc(c *gin.Context) string { return c.ClientIP() }
func errorHandler(c *gin.Context, info ratelimit.Info) { c.JSON(http.StatusTooManyRequests, gin.H{"error": "Too many requests"}) }

func main() {
	err := godotenv.Load()
	if err != nil { log.Println("Warning: .env file not found") }
	dsn := os.Getenv("DB_DSN")
	if dsn == "" { log.Fatal("DB_DSN not set") }
	jwtKey := os.Getenv("JWT_SECRET_KEY")
	if jwtKey == "" { log.Fatal("JWT_SECRET_KEY not set") }
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil { log.Fatal("Failed to connect to database") }
	db.AutoMigrate(&model.User{}, &model.Course{})
	seedCourses(db)

	// --- Initialize Layers ---
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository, jwtKey)
	userHandler := handler.NewUserHandler(userService)

	// ADDED: Initialize course layers
	courseRepository := repository.NewCourseRepository(db)
	courseService := service.NewCourseService(courseRepository)
	courseHandler := handler.NewCourseHandler(courseService)


	// ... (Server Setup & Middleware) ...
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	router.Use(cors.New(config))
	store := ratelimit.InMemoryStore(&ratelimit.InMemoryOptions{
		Rate:  time.Second * 10,
		Limit: 1,
	})
	rateLimitMiddleware := ratelimit.RateLimiter(store, &ratelimit.Options{
		ErrorHandler: errorHandler,
		KeyFunc:      keyFunc,
	})

	// --- API v1 Route Group ---
	apiV1 := router.Group("/api/v1")
	{
		// Public routes
		apiV1.POST("/register", rateLimitMiddleware, userHandler.Register)
		apiV1.POST("/login", userHandler.Login)
		
		// Protected routes
		authorized := apiV1.Group("/")
		authorized.Use(middleware.AuthMiddleware())
		{
			authorized.GET("/me", userHandler.GetMe)
			authorized.GET("/courses", courseHandler.GetAll)
			authorized.POST("/courses", courseHandler.Create)
			authorized.GET("/courses/:id", courseHandler.GetByID)
		}
	}

	log.Println("Starting server on http://localhost:8080")
	router.Run(":8080")
}

func seedCourses(db *gorm.DB) {
	var count int64
	db.Model(&model.Course{}).Count(&count)
	if count == 0 {
		log.Println("Seeding courses table...")
		courses := []model.Course{
			{Title: "Introduction to Go", Description: "Learn the fundamentals of the Go programming language.", UserID: 1},
			{Title: "Building Web Apps with Nuxt", Description: "Create modern and fast web applications with Nuxt 4.", UserID: 1},
			{Title: "Mastering PostgreSQL", Description: "A deep dive into PostgreSQL for developers.", UserID: 1},
		}
		if err := db.Create(&courses).Error; err != nil {
			log.Fatal("Failed to seed courses:", err)
		}
		log.Println("Courses seeded successfully.")
	}
}