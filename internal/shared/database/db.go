package database

import (
	"log"

	"my-api/internal/modules/auth/models"
	auth_model "my-api/internal/modules/auth/models"
	address_model "my-api/internal/modules/option/address/models"

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
		&address_model.Country{},
		&address_model.City{},
	)
}

func (db *DB) Close() error {
	sqlDB, err := db.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

type MockDB struct {
	Users []*models.User
}

func NewMockDB() *MockDB {
	return &MockDB{
		Users: make([]*models.User, 0),
	}
}

func (m *MockDB) GetDB() *gorm.DB {
	return &gorm.DB{}
}
