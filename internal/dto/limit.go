package dto

import "time"

type LimitResponse struct {
	ID          int64     `json:"id"`
	UserID      int64     `json:"user_id"`
	TenorMonths int       `json:"tenor_months"`
	LimitAmount float64   `json:"limit_amount"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type LimitRequest struct {
	UserID      int64   `json:"user_id" validate:"required,gt=0"`
	TenorMonths int     `json:"tenor_months" validate:"required,oneof=1 2 3 6"` // Cuma boleh 1,2,3,6
	LimitAmount float64 `json:"limit_amount" validate:"required,gt=0"`
}
