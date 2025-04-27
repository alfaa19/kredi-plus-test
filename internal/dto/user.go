package dto

import "time"

type UserRequest struct {
	NIK            string    `json:"nik" validate:"required,len=16,numeric"`
	FullName       string    `json:"full_name" validate:"required"`
	LegalName      string    `json:"legal_name" validate:"required"`
	PlaceOfBirth   string    `json:"place_of_birth" validate:"required"`
	DateOfBirth    time.Time `json:"date_of_birth" validate:"required"`
	Salary         float64   `json:"salary" validate:"required,gt=0"`
	KTPPhotoURL    string    `json:"ktp_photo_url" validate:"omitempty,url"`
	SelfiePhotoURL string    `json:"selfie_photo_url" validate:"omitempty,url"`
}

type UserResponse struct {
	ID             int64     `json:"id"`
	NIK            string    `json:"nik"`
	FullName       string    `json:"full_name"`
	LegalName      string    `json:"legal_name"`
	PlaceOfBirth   string    `json:"place_of_birth"`
	DateOfBirth    time.Time `json:"date_of_birth"`
	Salary         float64   `json:"salary"`
	KTPPhotoURL    string    `json:"ktp_photo_url"`
	SelfiePhotoURL string    `json:"selfie_photo_url"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
