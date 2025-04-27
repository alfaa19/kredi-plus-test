package service

import (
	"database/sql"
	"errors"

	"github.com/alfaa19/kredi-plus-test/internal/model"
	"github.com/alfaa19/kredi-plus-test/internal/repository"
	locker "github.com/alfaa19/kredi-plus-test/pkg"
)

type TransactionService interface {
	CreateTransaction(tx *model.Transaction, tenorMonths int) error
	GetByID(id int64) (*model.Transaction, error)
	GetAllByUserID(userID int64) ([]model.Transaction, error)
}

type transactionService struct {
	db        *sql.DB
	limitRepo repository.LimitRepository
	trxRepo   repository.TransactionRepository
	locker    *locker.Locker
}

func NewTransactionService(db *sql.DB, limitRepo repository.LimitRepository, trxRepo repository.TransactionRepository, locker *locker.Locker) TransactionService {
	return &transactionService{
		db:        db,
		limitRepo: limitRepo,
		trxRepo:   trxRepo,
		locker:    locker,
	}
}

// Create Transaction + Update Limits
func (s *transactionService) CreateTransaction(trx *model.Transaction, tenorMonths int) error {
	s.locker.Lock(trx.UserID)
	defer s.locker.Unlock(trx.UserID)

	txDB, err := s.db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			txDB.Rollback()
		} else {
			txDB.Commit()
		}
	}()

	limit, err := s.limitRepo.GetByUserIDAndTenor(trx.UserID, tenorMonths)
	if err != nil {
		return err
	}

	if limit.LimitAmount < trx.OTRAmount {
		return errors.New("limit tenor tidak mencukupi")
	}

	allLimits, err := s.limitRepo.GetAllByUserID(trx.UserID)
	if err != nil {
		return err
	}

	for _, l := range allLimits {
		newLimit := l.LimitAmount - trx.OTRAmount
		if newLimit < 0 {
			newLimit = 0
		}
		l.LimitAmount = newLimit

		if err := s.limitRepo.Update(&l); err != nil {
			return err
		}
	}

	if err := s.trxRepo.Create(trx); err != nil {
		return err
	}

	return nil
}

// New Method: Get By Transaction ID
func (s *transactionService) GetByID(id int64) (*model.Transaction, error) {
	return s.trxRepo.GetByID(id)
}

// New Method: Get All Transactions By User ID
func (s *transactionService) GetAllByUserID(userID int64) ([]model.Transaction, error) {
	return s.trxRepo.GetAllByUserID(userID)
}
