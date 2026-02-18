package service

import (
	"my-api/internal/modules/option/address/models"
	"my-api/internal/modules/option/address/repository"
	"my-api/internal/shared/errors"
)

type AddressService struct {
	repo repository.AddressRepositoryInterface
}

func NewAddressService(repo repository.AddressRepositoryInterface) *AddressService {
	return &AddressService{
		repo: repo,
	}
}

func (s *AddressService) GetAllCountries() ([]*models.Country, error) {
	countries, err := s.repo.GetAllCountries()
	if err != nil {
		return nil, errors.ErrInternalServer
	}
	return countries, nil
}

func (s *AddressService) GetCountryByCode(code string) (*models.Country, error) {
	if code == "" {
		return nil, errors.ErrBadRequest
	}

	country, err := s.repo.GetCountryByCode(code)
	if err != nil {
		return nil, errors.ErrInternalServer
	}
	if country == nil {
		return nil, errors.ErrNotFound
	}

	return country, nil
}

func (s *AddressService) GetAllCities() ([]*models.City, error) {
	cities, err := s.repo.GetAllCities()
	if err != nil {
		return nil, errors.ErrInternalServer
	}
	return cities, nil
}

func (s *AddressService) GetCityByID(id uint) (*models.City, error) {
	if id == 0 {
		return nil, errors.ErrBadRequest
	}

	city, err := s.repo.GetCityByID(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}
	if city == nil {
		return nil, errors.ErrNotFound
	}

	return city, nil
}

func (s *AddressService) GetCitiesByCountryCode(code string) ([]*models.City, error) {
	if code == "" {
		return nil, errors.ErrBadRequest
	}

	cities, err := s.repo.GetCitiesByCountryCode(code)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	return cities, nil
}
