package users

import (
	"errors"
	"strings"
)

// Service berisi logika bisnis untuk entitas User.
type Service struct {
	repo *Repository
}

// NewService membuat instance service baru.
func NewService(repo *Repository) *Service { return &Service{repo: repo} }

// GetAll mengembalikan semua user.
func (s *Service) GetAll() ([]User, error) { return s.repo.FindAll() }

// GetByID mencari user berdasarkan ID.
func (s *Service) GetByID(id uint) (User, error) {
	u, err := s.repo.FindByID(id)
	if err != nil {
		return User{}, err
	}
	return *u, nil
}

// Create membuat user baru dengan validasi sederhana.
func (s *Service) Create(u User) (User, error) {
	u.Username = strings.TrimSpace(u.Username)
	u.Email = strings.TrimSpace(u.Email)
	if u.Username == "" {
		return User{}, errors.New("username tidak boleh kosong")
	}
	if u.Email == "" {
		return User{}, errors.New("email tidak boleh kosong")
	}
	return s.repo.Create(u)
}

// Update memperbarui semua field user.
func (s *Service) Update(id uint, data User) (User, error) {
	u, err := s.repo.FindByID(id)
	if err != nil {
		return User{}, err
	}

	// Update field satu per satu
	if data.Username != "" {
		u.Username = data.Username
	}
	if data.Name != "" {
		u.Name = data.Name
	}
	if data.Email != "" {
		u.Email = data.Email
	}
	if data.Password != "" {
		u.Password = data.Password
	}

	if err := s.repo.Update(u); err != nil {
		return User{}, err
	}
	return *u, nil
}

// Delete menghapus user berdasarkan ID.
func (s *Service) Delete(id uint) error {
	u, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(u)
}
