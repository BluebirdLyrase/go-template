package database

import (
	"log"

	"my-api/internal/modules/auth/models"
	auth_model "my-api/internal/modules/auth/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

func Connect(dsn string) *DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("✅ Connected to database")

	return &DB{db}
}

func (db *DB) AutoMigrate() error {
	return db.DB.AutoMigrate(
		&auth_model.User{},
	)
}

func (db *DB) Close() error {
	sqlDB, err := db.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

// Mock DB for testing
type MockDB struct {
	Users []*models.User
}

func NewMockDB() *MockDB {
	return &MockDB{
		Users: make([]*models.User, 0),
	}
}

// Helper method to convert MockDB to GORM-like interface
func (m *MockDB) GetDB() *gorm.DB {
	return &gorm.DB{}
}
