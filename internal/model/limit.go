package model

import "time"

type Limit struct {
	ID          int64     `json:"id"`
	UserID      int64     `json:"user_id"`
	TenorMonths int       `json:"tenor_months"`
	LimitAmount float64   `json:"limit_amount"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
