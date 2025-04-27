package service

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/alfaa19/kredi-plus-test/internal/model"
)

// Mock LimitRepository
type MockLimitRepo struct {
	mock.Mock
}

func (m *MockLimitRepo) Create(limit *model.Limit) error {
	args := m.Called(limit)
	return args.Error(0)
}

func (m *MockLimitRepo) GetByUserIDAndTenor(userID int64, tenor int) (*model.Limit, error) {
	args := m.Called(userID, tenor)
	return args.Get(0).(*model.Limit), args.Error(1)
}

func (m *MockLimitRepo) GetAllByUserID(userID int64) ([]model.Limit, error) {
	args := m.Called(userID)
	return args.Get(0).([]model.Limit), args.Error(1)
}

func (m *MockLimitRepo) Update(limit *model.Limit) error {
	args := m.Called(limit)
	return args.Error(0)
}

// Mock TransactionRepository
type MockTransactionRepo struct {
	mock.Mock
}

func (m *MockTransactionRepo) Create(trx *model.Transaction) error {
	args := m.Called(trx)
	return args.Error(0)
}

func (m *MockTransactionRepo) GetByID(id int64) (*model.Transaction, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Transaction), args.Error(1)
}

func (m *MockTransactionRepo) GetAllByUserID(userID int64) ([]model.Transaction, error) {
	args := m.Called(userID)
	return args.Get(0).([]model.Transaction), args.Error(1)
}

func TestCreateTransaction_Success(t *testing.T) {
	limitRepo := new(MockLimitRepo)
	trxRepo := new(MockTransactionRepo)

	service := &transactionService{
		db:        nil, // kita ga test db real
		limitRepo: limitRepo,
		trxRepo:   trxRepo,
	}

	userID := int64(1)
	tenor := 3
	otrAmount := 100000.0

	trx := &model.Transaction{
		UserID:    userID,
		OTRAmount: otrAmount,
	}

	limitRepo.On("GetByUserIDAndTenor", userID, tenor).Return(&model.Limit{
		UserID:      userID,
		TenorMonths: tenor,
		LimitAmount: 200000.0,
	}, nil)

	limitRepo.On("GetAllByUserID", userID).Return([]model.Limit{
		{ID: 1, UserID: userID, LimitAmount: 200000},
		{ID: 2, UserID: userID, LimitAmount: 300000},
	}, nil)

	limitRepo.On("Update", mock.AnythingOfType("*model.Limit")).Return(nil).Twice()

	trxRepo.On("Create", trx).Return(nil)

	err := service.CreateTransaction(trx, tenor)

	assert.NoError(t, err)
}

func TestCreateTransaction_Fail_Limit_NotEnough(t *testing.T) {
	limitRepo := new(MockLimitRepo)
	trxRepo := new(MockTransactionRepo)

	service := &transactionService{
		db:        nil,
		limitRepo: limitRepo,
		trxRepo:   trxRepo,
	}

	userID := int64(1)
	tenor := 6
	otrAmount := 500000.0

	trx := &model.Transaction{
		UserID:    userID,
		OTRAmount: otrAmount,
	}

	limitRepo.On("GetByUserIDAndTenor", userID, tenor).Return(&model.Limit{
		UserID:      userID,
		TenorMonths: tenor,
		LimitAmount: 300000.0,
	}, nil)

	err := service.CreateTransaction(trx, tenor)

	assert.Error(t, err)
	assert.Equal(t, "limit tenor tidak mencukupi", err.Error())
}

func TestGetTransactionByID_Success(t *testing.T) {
	limitRepo := new(MockLimitRepo)
	trxRepo := new(MockTransactionRepo)

	service := &transactionService{
		db:        nil,
		limitRepo: limitRepo,
		trxRepo:   trxRepo,
	}

	trx := &model.Transaction{
		ID:             1,
		UserID:         1,
		ContractNumber: "CTR001",
	}

	trxRepo.On("GetByID", trx.ID).Return(trx, nil)

	result, err := service.GetByID(trx.ID)

	assert.NoError(t, err)
	assert.Equal(t, trx, result)
}

func TestGetTransactionByID_Fail(t *testing.T) {
	limitRepo := new(MockLimitRepo)
	trxRepo := new(MockTransactionRepo)

	service := &transactionService{
		db:        nil,
		limitRepo: limitRepo,
		trxRepo:   trxRepo,
	}

	trxID := int64(99)

	trxRepo.On("GetByID", trxID).Return(&model.Transaction{}, errors.New("transaction not found"))

	result, err := service.GetByID(trxID)

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestGetAllTransactionsByUserID_Success(t *testing.T) {
	limitRepo := new(MockLimitRepo)
	trxRepo := new(MockTransactionRepo)

	service := &transactionService{
		db:        nil,
		limitRepo: limitRepo,
		trxRepo:   trxRepo,
	}

	userID := int64(1)

	trxList := []model.Transaction{
		{ID: 1, UserID: userID, ContractNumber: "CTR001"},
		{ID: 2, UserID: userID, ContractNumber: "CTR002"},
	}

	trxRepo.On("GetAllByUserID", userID).Return(trxList, nil)

	result, err := service.GetAllByUserID(userID)

	assert.NoError(t, err)
	assert.Equal(t, trxList, result)
}

func TestGetAllTransactionsByUserID_Fail(t *testing.T) {
	limitRepo := new(MockLimitRepo)
	trxRepo := new(MockTransactionRepo)

	service := &transactionService{
		db:        nil,
		limitRepo: limitRepo,
		trxRepo:   trxRepo,
	}

	userID := int64(99)

	trxRepo.On("GetAllByUserID", userID).Return([]model.Transaction{}, errors.New("user not found"))

	result, err := service.GetAllByUserID(userID)

	assert.Error(t, err)
	assert.Nil(t, result)
}
