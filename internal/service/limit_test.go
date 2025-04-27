package service

import (
	"errors"
	"testing"

	"github.com/alfaa19/kredi-plus-test/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockLimitRepoForService struct {
	mock.Mock
}

func (m *MockLimitRepoForService) Create(limit *model.Limit) error {
	args := m.Called(limit)
	return args.Error(0)
}

func (m *MockLimitRepoForService) GetByUserIDAndTenor(userID int64, tenor int) (*model.Limit, error) {
	args := m.Called(userID, tenor)
	return args.Get(0).(*model.Limit), args.Error(1)
}

func (m *MockLimitRepoForService) GetAllByUserID(userID int64) ([]model.Limit, error) {
	args := m.Called(userID)
	return args.Get(0).([]model.Limit), args.Error(1)
}

func (m *MockLimitRepoForService) Update(limit *model.Limit) error {
	args := m.Called(limit)
	return args.Error(0)
}

func TestCreateLimit_Success(t *testing.T) {
	limitRepo := new(MockLimitRepoForService)
	service := NewLimitService(limitRepo)

	limit := &model.Limit{ID: 1, UserID: 1}

	limitRepo.On("Create", limit).Return(nil)

	err := service.Create(limit)
	assert.NoError(t, err)
}

func TestCreateLimit_Fail(t *testing.T) {
	limitRepo := new(MockLimitRepoForService)
	service := NewLimitService(limitRepo)

	limit := &model.Limit{ID: 1, UserID: 1}

	limitRepo.On("Create", limit).Return(errors.New("failed create limit"))

	err := service.Create(limit)
	assert.Error(t, err)
}

func TestGetLimitByUserIDAndTenor_Success(t *testing.T) {
	limitRepo := new(MockLimitRepoForService)
	service := NewLimitService(limitRepo)

	limit := &model.Limit{ID: 1, UserID: 1, TenorMonths: 6}

	limitRepo.On("GetByUserIDAndTenor", limit.UserID, limit.TenorMonths).Return(limit, nil)

	result, err := service.GetByUserIDAndTenor(limit.UserID, limit.TenorMonths)
	assert.NoError(t, err)
	assert.Equal(t, limit, result)
}

func TestGetLimitByUserIDAndTenor_Fail(t *testing.T) {
	limitRepo := new(MockLimitRepoForService)
	service := NewLimitService(limitRepo)

	limitRepo.On("GetByUserIDAndTenor", int64(99), 6).Return(&model.Limit{}, errors.New("limit not found"))

	result, err := service.GetByUserIDAndTenor(99, 6)
	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestGetAllLimitsByUserID_Success(t *testing.T) {
	limitRepo := new(MockLimitRepoForService)
	service := NewLimitService(limitRepo)

	userID := int64(1)
	limits := []model.Limit{
		{ID: 1, UserID: userID},
	}

	limitRepo.On("GetAllByUserID", userID).Return(limits, nil)

	result, err := service.GetAllByUserID(userID)
	assert.NoError(t, err)
	assert.Equal(t, limits, result)
}

func TestGetAllLimitsByUserID_Fail(t *testing.T) {
	limitRepo := new(MockLimitRepoForService)
	service := NewLimitService(limitRepo)

	userID := int64(99)

	limitRepo.On("GetAllByUserID", userID).Return([]model.Limit{}, errors.New("failed get all limits"))

	result, err := service.GetAllByUserID(userID)
	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestUpdateLimit_Success(t *testing.T) {
	limitRepo := new(MockLimitRepoForService)
	service := NewLimitService(limitRepo)

	limit := &model.Limit{ID: 1, UserID: 1}

	limitRepo.On("Update", limit).Return(nil)

	err := service.Update(limit)
	assert.NoError(t, err)
}

func TestUpdateLimit_Fail(t *testing.T) {
	limitRepo := new(MockLimitRepoForService)
	service := NewLimitService(limitRepo)

	limit := &model.Limit{ID: 1, UserID: 1}

	limitRepo.On("Update", limit).Return(errors.New("failed update limit"))

	err := service.Update(limit)
	assert.Error(t, err)
}
