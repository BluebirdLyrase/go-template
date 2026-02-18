package models

type Country struct {
	Code string `gorm:"primaryKey;size:2"`
	Name string `gorm:"size:100;not null;unique"`
}
