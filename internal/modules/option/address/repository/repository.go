package repository

import (
	"my-api/internal/modules/option/address/models"

	"gorm.io/gorm"
)

type AddressRepository struct {
	db *gorm.DB
}

func NewAddressRepository(db *gorm.DB) *AddressRepository {
	return &AddressRepository{db: db}
}

func (r *AddressRepository) GetCityByID(id uint) (*models.City, error) {
	var city models.City

	if err := r.db.
		First(&city, id).Error; err != nil {

		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &city, nil
}

func (r *AddressRepository) GetAllCities() ([]*models.City, error) {
	var cities []*models.City

	if err := r.db.
		Find(&cities).Error; err != nil {
		return nil, err
	}

	return cities, nil
}

func (r *AddressRepository) GetCitiesByCountryCode(code string) ([]*models.City, error) {
	var cities []*models.City

	if err := r.db.
		Where("country_code = ?", code).
		Find(&cities).Error; err != nil {
		return nil, err
	}

	return cities, nil
}

func (r *AddressRepository) GetCountryByCode(code string) (*models.Country, error) {
	var country models.Country

	if err := r.db.First(&country, "code = ?", code).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &country, nil
}

func (r *AddressRepository) GetAllCountries() ([]*models.Country, error) {
	var countries []*models.Country

	if err := r.db.Find(&countries).Error; err != nil {
		return nil, err
	}

	return countries, nil
}
