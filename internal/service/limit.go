package service

import (
	"github.com/alfaa19/kredi-plus-test/internal/model"
	"github.com/alfaa19/kredi-plus-test/internal/repository"
)

type LimitService interface {
	Create(limit *model.Limit) error
	GetByUserIDAndTenor(userID int64, tenor int) (*model.Limit, error)
	GetAllByUserID(userID int64) ([]model.Limit, error)
	Update(limit *model.Limit) error
}

type limitService struct {
	limitRepo repository.LimitRepository
}

func NewLimitService(limitRepo repository.LimitRepository) LimitService {
	return &limitService{limitRepo: limitRepo}
}

func (s *limitService) Create(limit *model.Limit) error {
	return s.limitRepo.Create(limit)
}

func (s *limitService) GetByUserIDAndTenor(userID int64, tenor int) (*model.Limit, error) {
	return s.limitRepo.GetByUserIDAndTenor(userID, tenor)
}

func (s *limitService) GetAllByUserID(userID int64) ([]model.Limit, error) {
	return s.limitRepo.GetAllByUserID(userID)
}

func (s *limitService) Update(limit *model.Limit) error {
	return s.limitRepo.Update(limit)
}
