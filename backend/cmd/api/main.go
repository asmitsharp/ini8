package main

import (
	"context"
	"ini8/internal/config"
	"ini8/internal/handler"
	"ini8/internal/middleware"
	"ini8/internal/repository"
	"ini8/internal/service"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func runMigrations(db *gorm.DB) {
	// Read the SQL migration file
	sqlFilePath := filepath.Join("migrations", "registrations_table.sql")
	sqlBytes, err := os.ReadFile(sqlFilePath)
	if err != nil {
		log.Fatalf("Failed to read migration file: %v", err)
	}

	// Execute the SQL command
	if err := db.Exec(string(sqlBytes)).Error; err != nil {
		log.Fatalf("Failed to execute migration: %v", err)
	}

	log.Println("Migrations have been successfully run.")
}

func main() {
	// if err := godotenv.Load(); err != nil {
	// 	log.Printf("Environment File Not Found")
	// }

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load Config: %v", err)
	}

	db, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to Database: %v", err)
	}

	log.Println("Database connection established.")

	// Run migrations
	runMigrations(db)

	repo := repository.NewRegistrationRepository(db)
	svc := service.NewRegistrationService(repo)
	h := handler.NewRegistrationHandler(svc)

	r := gin.Default()

	r.Use(middleware.Logger())
	r.Use(middleware.CORS())
	r.Use(middleware.Recovery())

	api := r.Group("/api/")
	{
		api.POST("/registrations", h.Create)
		api.GET("/registrations/:id", h.Get)
		api.GET("/registrations", h.List)
		api.PUT("/registrations/:id", h.Update)
		api.DELETE("/registrations/:id", h.Delete)
	}

	srv := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	log.Printf("Server is up and running on port %s", cfg.Port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server has been shut down gracefully.")
}
