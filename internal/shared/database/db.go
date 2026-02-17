package database

import (
	"log"

	"my-api/internal/modules/auth/models"
	productmodels "my-api/internal/modules/product/models"
	usermodels "my-api/internal/modules/user/models"

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
		&models.User{},
		&usermodels.Profile{},
		&productmodels.Product{},
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
	Users    []*models.User
	Products []*productmodels.Product
}

func NewMockDB() *MockDB {
	return &MockDB{
		Users:    make([]*models.User, 0),
		Products: make([]*productmodels.Product, 0),
	}
}

// Helper method to convert MockDB to GORM-like interface
func (m *MockDB) GetDB() *gorm.DB {
	return &gorm.DB{}
}
