package models

import (
	"gorm.io/gorm"
	"time"
)

type Profile struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `gorm:"index" json:"user_id"`
	Bio       string    `json:"bio"`
	AvatarURL string    `json:"avatar_url"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
	City      string    `json:"city"`
	Country   string    `json:"country"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type UpdateProfileRequest struct {
	Bio       string `json:"bio"`
	AvatarURL string `json:"avatar_url"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	City      string `json:"city"`
	Country   string `json:"country"`
}

type ProfileResponse struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	Bio       string    `json:"bio"`
	AvatarURL string    `json:"avatar_url"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
	City      string    `json:"city"`
	Country   string    `json:"country"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
