package dto

import "time"

type TransactionRequest struct {
	UserID         int64   `json:"user_id" validate:"required,gt=0"`
	TenorMonths    int     `json:"tenor_months" validate:"required,oneof=1 2 3 6"`
	ContractNumber string  `json:"contract_number" validate:"required"`
	OTRAmount      float64 `json:"otr_amount" validate:"required,gt=0"`
	AdminFee       float64 `json:"admin_fee" validate:"omitempty,gte=0"`
	Installment    float64 `json:"installment_amount" validate:"omitempty,gte=0"`
	Interest       float64 `json:"interest_amount" validate:"omitempty,gte=0"`
	AssetName      string  `json:"asset_name" validate:"required"`
}

type TransactionResponse struct {
	ID             int64     `json:"id"`
	UserID         int64     `json:"user_id"`
	ContractNumber string    `json:"contract_number"`
	OTRAmount      float64   `json:"otr_amount"`
	AdminFee       float64   `json:"admin_fee"`
	Installment    float64   `json:"installment_amount"`
	Interest       float64   `json:"interest_amount"`
	AssetName      string    `json:"asset_name"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
