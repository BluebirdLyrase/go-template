package repository

import (
	"gorm.io/gorm"
	"my-api/internal/modules/user/models"
)

type ProfileRepository struct {
	db *gorm.DB
}

func NewProfileRepository(db *gorm.DB) *ProfileRepository {
	return &ProfileRepository{db: db}
}

func (r *ProfileRepository) Save(profile *models.Profile) error {
	return r.db.Save(profile).Error
}

func (r *ProfileRepository) GetByUserID(userID uint) (*models.Profile, error) {
	var profile models.Profile
	if err := r.db.Where("user_id = ?", userID).First(&profile).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &profile, nil
}

func (r *ProfileRepository) GetByID(id uint) (*models.Profile, error) {
	var profile models.Profile
	if err := r.db.First(&profile, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &profile, nil
}

func (r *ProfileRepository) Delete(userID uint) error {
	return r.db.Where("user_id = ?", userID).Delete(&models.Profile{}).Error
}

func (r *ProfileRepository) GetAll() ([]*models.Profile, error) {
	var profiles []*models.Profile
	if err := r.db.Find(&profiles).Error; err != nil {
		return nil, err
	}
	return profiles, nil
}
