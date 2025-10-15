package users

import "gorm.io/gorm"

// Repository menangani operasi database untuk entitas User.
type Repository struct {
	db *gorm.DB
}

// NewRepository membuat instance repository baru.
func NewRepository(db *gorm.DB) *Repository { return &Repository{db: db} }

// FindAll mengembalikan semua data user.
func (r *Repository) FindAll() ([]User, error) {
	var out []User
	return out, r.db.Find(&out).Error
}

// FindByID mencari user berdasarkan ID.
func (r *Repository) FindByID(id uint) (*User, error) {
	var u User
	if err := r.db.First(&u, id).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

// FindByEmail mencari user berdasarkan email.
func (r *Repository) FindByEmail(email string) (*User, error) {
	var u User
	if err := r.db.Where("email = ?", email).First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

// Create menambahkan user baru ke database.
func (r *Repository) Create(u User) (User, error) {
	return u, r.db.Create(&u).Error
}

// Update memperbarui data user yang sudah ada.
func (r *Repository) Update(u *User) error {
	return r.db.Save(u).Error
}

// Delete menghapus user dari database.
func (r *Repository) Delete(u *User) error {
	return r.db.Delete(u).Error
}
