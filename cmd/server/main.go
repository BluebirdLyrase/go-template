package main

import (
	"log"

	"my-api/internal/modules/auth"
	"my-api/internal/shared/config"
	"my-api/internal/shared/database"
	"my-api/internal/shared/router"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Connect to database
	db := database.Connect(cfg.GetDSN())
	defer db.Close()

	// Auto-migrate models
	if err := db.AutoMigrate(); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	// Initialize router
	r := router.New()

	// Initialize and register modules
	initializeModules(r, db.DB)

	// Start server
	log.Printf("🚀 Server starting on port %s", cfg.ServerPort)
	if err := r.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func initializeModules(router *gin.Engine, db *gorm.DB) {
	// Auth module
	authModule := auth.NewModule(db)
	authModule.RegisterRoutes(router)
	log.Println("✅ Auth module registered")
}
