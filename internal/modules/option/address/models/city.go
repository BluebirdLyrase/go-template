package models

type City struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"size:150;not null;index"`

	CountryCode string  `gorm:"size:2;not null;index"`
	Country     Country `gorm:"foreignKey:CountryCode;references:Code"`
}
