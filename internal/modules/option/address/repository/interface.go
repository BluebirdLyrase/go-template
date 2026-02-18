package repository

import "my-api/internal/modules/option/address/models"

type AddressRepositoryInterface interface {
	GetCityByID(id uint) (*models.City, error)
	GetAllCities() ([]*models.City, error)
	GetCitiesByCountryCode(code string) ([]*models.City, error)

	GetCountryByCode(code string) (*models.Country, error)
	GetAllCountries() ([]*models.Country, error)
}
