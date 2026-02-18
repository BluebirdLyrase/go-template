package mocks

import (
	authmodels "my-api/internal/modules/auth/models"
	// usermodels "my-api/internal/modules/user/models"
	// productmodels "my-api/internal/modules/product/models"
)

// Mock User Repository
type MockUserRepository struct {
	Users []*authmodels.User
}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{
		Users: make([]*authmodels.User, 0),
	}
}

func (m *MockUserRepository) Create(user *authmodels.User) error {
	user.ID = uint(len(m.Users) + 1)
	m.Users = append(m.Users, user)
	return nil
}

func (m *MockUserRepository) GetByID(id uint) (*authmodels.User, error) {
	for _, u := range m.Users {
		if u.ID == id {
			return u, nil
		}
	}
	return nil, nil
}

func (m *MockUserRepository) GetByEmail(email string) (*authmodels.User, error) {
	for _, u := range m.Users {
		if u.Email == email {
			return u, nil
		}
	}
	return nil, nil
}

func (m *MockUserRepository) Update(user *authmodels.User) error {
	for i, u := range m.Users {
		if u.ID == user.ID {
			m.Users[i] = user
			return nil
		}
	}
	return nil
}

func (m *MockUserRepository) Delete(id uint) error {
	for i, u := range m.Users {
		if u.ID == id {
			m.Users = append(m.Users[:i], m.Users[i+1:]...)
			return nil
		}
	}
	return nil
}

func (m *MockUserRepository) GetAll() ([]*authmodels.User, error) {
	return m.Users, nil
}
