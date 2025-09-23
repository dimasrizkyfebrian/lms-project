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

// ... keyFunc and errorHandler remain the same ...
func keyFunc(c *gin.Context) string { return c.ClientIP() }
func errorHandler(c *gin.Context, info ratelimit.Info) { c.JSON(http.StatusTooManyRequests, gin.H{"error": "Too many requests"}) }

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, reading from environment")
	}
	
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		log.Fatal("DB_DSN environment variable not set")
	}
	
	// --- LOAD JWT KEY FROM ENV ---
	jwtKey := os.Getenv("JWT_SECRET_KEY")
	if jwtKey == "" {
		log.Fatal("JWT_SECRET_KEY environment variable not set")
	}

	// --- Database Connection ---
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	db.AutoMigrate(&model.User{})
	log.Println("Database migration completed.")

	// --- Initialize Layers (Dependency Injection) ---
	userRepository := repository.NewUserRepository(db)
	// Pass the jwtKey to the UserService
	userService := service.NewUserService(userRepository, jwtKey)
	userHandler := handler.NewUserHandler(userService)

	// ... Rest of the server setup remains the same ...
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

	apiV1 := router.Group("/api/v1")
	{
		apiV1.POST("/register", rateLimitMiddleware, userHandler.Register)
		apiV1.POST("/login", userHandler.Login)
		apiV1.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "UP"})
		})
		
		authorized := apiV1.Group("/")
		authorized.Use(middleware.AuthMiddleware())
		{
			authorized.GET("/me", userHandler.GetMe)
		}
	}

	log.Println("Starting server on http://localhost:8080")
	router.Run(":8080")
}