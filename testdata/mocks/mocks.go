package mocks

import (
	authmodels "my-api/internal/modules/auth/models"
	productmodels "my-api/internal/modules/product/models"
	usermodels "my-api/internal/modules/user/models"
)

// TODO
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

// Mock Profile Repository
type MockProfileRepository struct {
	Profiles []*usermodels.Profile
}

func NewMockProfileRepository() *MockProfileRepository {
	return &MockProfileRepository{
		Profiles: make([]*usermodels.Profile, 0),
	}
}

func (m *MockProfileRepository) Save(profile *usermodels.Profile) error {
	if profile.ID == 0 {
		profile.ID = uint(len(m.Profiles) + 1)
		m.Profiles = append(m.Profiles, profile)
	} else {
		for i, p := range m.Profiles {
			if p.ID == profile.ID {
				m.Profiles[i] = profile
				return nil
			}
		}
		m.Profiles = append(m.Profiles, profile)
	}
	return nil
}

func (m *MockProfileRepository) GetByUserID(userID uint) (*usermodels.Profile, error) {
	for _, p := range m.Profiles {
		if p.UserID == userID {
			return p, nil
		}
	}
	return nil, nil
}

func (m *MockProfileRepository) GetByID(id uint) (*usermodels.Profile, error) {
	for _, p := range m.Profiles {
		if p.ID == id {
			return p, nil
		}
	}
	return nil, nil
}

func (m *MockProfileRepository) Delete(userID uint) error {
	for i, p := range m.Profiles {
		if p.UserID == userID {
			m.Profiles = append(m.Profiles[:i], m.Profiles[i+1:]...)
			return nil
		}
	}
	return nil
}

func (m *MockProfileRepository) GetAll() ([]*usermodels.Profile, error) {
	return m.Profiles, nil
}

// Mock Product Repository
type MockProductRepository struct {
	Products []*productmodels.Product
}

func NewMockProductRepository() *MockProductRepository {
	return &MockProductRepository{
		Products: make([]*productmodels.Product, 0),
	}
}

func (m *MockProductRepository) Create(product *productmodels.Product) error {
	product.ID = uint(len(m.Products) + 1)
	m.Products = append(m.Products, product)
	return nil
}

func (m *MockProductRepository) GetByID(id uint) (*productmodels.Product, error) {
	for _, p := range m.Products {
		if p.ID == id {
			return p, nil
		}
	}
	return nil, nil
}

func (m *MockProductRepository) Update(product *productmodels.Product) error {
	for i, p := range m.Products {
		if p.ID == product.ID {
			m.Products[i] = product
			return nil
		}
	}
	return nil
}

func (m *MockProductRepository) Delete(id uint) error {
	for i, p := range m.Products {
		if p.ID == id {
			m.Products = append(m.Products[:i], m.Products[i+1:]...)
			return nil
		}
	}
	return nil
}

func (m *MockProductRepository) GetAll() ([]*productmodels.Product, error) {
	return m.Products, nil
}

func (m *MockProductRepository) GetByCategory(category string) ([]*productmodels.Product, error) {
	var result []*productmodels.Product
	for _, p := range m.Products {
		if p.Category == category {
			result = append(result, p)
		}
	}
	return result, nil
}
