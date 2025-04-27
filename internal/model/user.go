package model

import "time"

type User struct {
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
