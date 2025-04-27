package repository

import (
	"database/sql"

	"github.com/alfaa19/kredi-plus-test/internal/model"
)

type TransactionRepository interface {
	Create(transaction *model.Transaction) error
	GetByID(id int64) (*model.Transaction, error)
	GetAllByUserID(userID int64) ([]model.Transaction, error)
}

type transactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) TransactionRepository {
	return &transactionRepository{db: db}
}

func (r *transactionRepository) Create(transaction *model.Transaction) error {
	query := `INSERT INTO transactions (user_id, contract_number, otr_amount, admin_fee, installment_amount, interest_amount, asset_name) VALUES (?, ?, ?, ?, ?, ?, ?)`
	_, err := r.db.Exec(query, transaction.UserID, transaction.ContractNumber, transaction.OTRAmount, transaction.AdminFee, transaction.Installment, transaction.Interest, transaction.AssetName)
	return err
}

func (r *transactionRepository) GetByID(id int64) (*model.Transaction, error) {
	query := `SELECT id, user_id, contract_number, otr_amount, admin_fee, installment_amount, interest_amount, asset_name, created_at, updated_at FROM transactions WHERE id = ?`
	row := r.db.QueryRow(query, id)

	var trx model.Transaction
	err := row.Scan(&trx.ID, &trx.UserID, &trx.ContractNumber, &trx.OTRAmount, &trx.AdminFee, &trx.Installment, &trx.Interest, &trx.AssetName, &trx.CreatedAt, &trx.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &trx, nil
}

func (r *transactionRepository) GetAllByUserID(userID int64) ([]model.Transaction, error) {
	query := `SELECT id, user_id, contract_number, otr_amount, admin_fee, installment_amount, interest_amount, asset_name, created_at, updated_at FROM transactions WHERE user_id = ?`
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []model.Transaction
	for rows.Next() {
		var trx model.Transaction
		if err := rows.Scan(&trx.ID, &trx.UserID, &trx.ContractNumber, &trx.OTRAmount, &trx.AdminFee, &trx.Installment, &trx.Interest, &trx.AssetName, &trx.CreatedAt, &trx.UpdatedAt); err != nil {
			return nil, err
		}
		transactions = append(transactions, trx)
	}
	return transactions, nil
}
