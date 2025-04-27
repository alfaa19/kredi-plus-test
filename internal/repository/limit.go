package repository

import (
	"database/sql"

	"github.com/alfaa19/kredi-plus-test/internal/model"
)

type LimitRepository interface {
	Create(limit *model.Limit) error
	GetByUserIDAndTenor(userID int64, tenor int) (*model.Limit, error)
	GetAllByUserID(userID int64) ([]model.Limit, error)
	Update(limit *model.Limit) error
}

type limitRepository struct {
	db *sql.DB
}

func NewLimitRepository(db *sql.DB) LimitRepository {
	return &limitRepository{db: db}
}

func (r *limitRepository) Create(limit *model.Limit) error {
	query := `INSERT INTO limits (user_id, tenor_months, limit_amount) VALUES (?, ?, ?)`
	_, err := r.db.Exec(query, limit.UserID, limit.TenorMonths, limit.LimitAmount)
	return err
}

func (r *limitRepository) GetByUserIDAndTenor(userID int64, tenor int) (*model.Limit, error) {
	query := `SELECT id, user_id, tenor_months, limit_amount, created_at, updated_at FROM limits WHERE user_id = ? AND tenor_months = ?`
	row := r.db.QueryRow(query, userID, tenor)

	var limit model.Limit
	err := row.Scan(&limit.ID, &limit.UserID, &limit.TenorMonths, &limit.LimitAmount, &limit.CreatedAt, &limit.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &limit, nil
}

func (r *limitRepository) GetAllByUserID(userID int64) ([]model.Limit, error) {
	query := `SELECT id, user_id, tenor_months, limit_amount, created_at, updated_at FROM limits WHERE user_id = ?`
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var limits []model.Limit
	for rows.Next() {
		var limit model.Limit
		if err := rows.Scan(&limit.ID, &limit.UserID, &limit.TenorMonths, &limit.LimitAmount, &limit.CreatedAt, &limit.UpdatedAt); err != nil {
			return nil, err
		}
		limits = append(limits, limit)
	}
	return limits, nil
}

func (r *limitRepository) Update(limit *model.Limit) error {
	query := `UPDATE limits SET limit_amount = ? WHERE id = ?`
	_, err := r.db.Exec(query, limit.LimitAmount, limit.ID)
	return err
}
