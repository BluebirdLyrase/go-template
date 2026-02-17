package service

import (
	"my-api/internal/modules/user/models"
	"my-api/internal/modules/user/repository"
	"my-api/internal/shared/errors"
)

type UserService struct {
	repo *repository.ProfileRepository
}

func NewUserService(repo *repository.ProfileRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetProfile(userID uint) (*models.Profile, error) {
	profile, err := s.repo.GetByUserID(userID)
	if err != nil {
		return nil, errors.ErrUserNotFound
	}
	return profile, nil
}

func (s *UserService) UpdateProfile(userID uint, req *models.UpdateProfileRequest) (*models.Profile, error) {
	profile, err := s.repo.GetByUserID(userID)
	if err != nil || profile == nil {
		// Create new profile if doesn't exist
		profile = &models.Profile{UserID: userID}
	}

	profile.Bio = req.Bio
	profile.AvatarURL = req.AvatarURL
	profile.Phone = req.Phone
	profile.Address = req.Address
	profile.City = req.City
	profile.Country = req.Country

	if err := s.repo.Save(profile); err != nil {
		return nil, errors.ErrInternalServer
	}

	return profile, nil
}

func (s *UserService) CreateProfile(userID uint) (*models.Profile, error) {
	profile := &models.Profile{UserID: userID}
	if err := s.repo.Save(profile); err != nil {
		return nil, errors.ErrInternalServer
	}
	return profile, nil
}
