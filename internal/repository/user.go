package repository

import (
	"database/sql"

	"github.com/alfaa19/kredi-plus-test/internal/model"
)

type UserRepository interface {
	Create(user *model.User) error
	GetByID(id int64) (*model.User, error)
	Update(user *model.User) error
	Delete(id int64) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *model.User) error {
	query := `INSERT INTO users (nik, full_name, legal_name, place_of_birth, date_of_birth, salary, ktp_photo_url, selfie_photo_url) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := r.db.Exec(query, user.NIK, user.FullName, user.LegalName, user.PlaceOfBirth, user.DateOfBirth, user.Salary, user.KTPPhotoURL, user.SelfiePhotoURL)
	return err
}

func (r *userRepository) GetByID(id int64) (*model.User, error) {
	query := `SELECT id, nik, full_name, legal_name, place_of_birth, date_of_birth, salary, ktp_photo_url, selfie_photo_url, created_at, updated_at FROM users WHERE id = ?`
	row := r.db.QueryRow(query, id)

	var user model.User
	err := row.Scan(&user.ID, &user.NIK, &user.FullName, &user.LegalName, &user.PlaceOfBirth, &user.DateOfBirth, &user.Salary, &user.KTPPhotoURL, &user.SelfiePhotoURL, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Update(user *model.User) error {
	query := `UPDATE users SET full_name = ?, legal_name = ?, place_of_birth = ?, date_of_birth = ?, salary = ?, ktp_photo_url = ?, selfie_photo_url = ? WHERE id = ?`
	_, err := r.db.Exec(query, user.FullName, user.LegalName, user.PlaceOfBirth, user.DateOfBirth, user.Salary, user.KTPPhotoURL, user.SelfiePhotoURL, user.ID)
	return err
}

func (r *userRepository) Delete(id int64) error {
	query := `DELETE FROM users WHERE id = ?`
	_, err := r.db.Exec(query, id)
	return err
}
