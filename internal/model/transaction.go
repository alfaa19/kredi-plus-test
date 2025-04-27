package model

import "time"

type Transaction struct {
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
