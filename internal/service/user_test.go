package service

import (
	"errors"
	"testing"

	"github.com/alfaa19/kredi-plus-test/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepo struct {
	mock.Mock
}

func (m *MockUserRepo) Create(user *model.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepo) GetByID(id int64) (*model.User, error) {
	args := m.Called(id)
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockUserRepo) Update(user *model.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepo) Delete(id int64) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestCreateUser_Success(t *testing.T) {
	userRepo := new(MockUserRepo)
	service := NewUserService(userRepo)

	user := &model.User{ID: 1, FullName: "Budi"}

	userRepo.On("Create", user).Return(nil)

	err := service.Create(user)
	assert.NoError(t, err)
}

func TestCreateUser_Fail(t *testing.T) {
	userRepo := new(MockUserRepo)
	service := NewUserService(userRepo)

	user := &model.User{ID: 1, FullName: "Budi"}

	userRepo.On("Create", user).Return(errors.New("failed create"))

	err := service.Create(user)
	assert.Error(t, err)
}

func TestGetUserByID_Success(t *testing.T) {
	userRepo := new(MockUserRepo)
	service := NewUserService(userRepo)

	user := &model.User{ID: 1, FullName: "Budi"}

	userRepo.On("GetByID", user.ID).Return(user, nil)

	result, err := service.GetByID(user.ID)
	assert.NoError(t, err)
	assert.Equal(t, user, result)
}

func TestGetUserByID_Fail(t *testing.T) {
	userRepo := new(MockUserRepo)
	service := NewUserService(userRepo)

	userRepo.On("GetByID", int64(99)).Return(&model.User{}, errors.New("user not found"))

	result, err := service.GetByID(99)
	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestUpdateUser_Success(t *testing.T) {
	userRepo := new(MockUserRepo)
	service := NewUserService(userRepo)

	user := &model.User{ID: 1, FullName: "Budi"}

	userRepo.On("Update", user).Return(nil)

	err := service.Update(user)
	assert.NoError(t, err)
}

func TestUpdateUser_Fail(t *testing.T) {
	userRepo := new(MockUserRepo)
	service := NewUserService(userRepo)

	user := &model.User{ID: 1, FullName: "Budi"}

	userRepo.On("Update", user).Return(errors.New("failed update"))

	err := service.Update(user)
	assert.Error(t, err)
}

func TestDeleteUser_Success(t *testing.T) {
	userRepo := new(MockUserRepo)
	service := NewUserService(userRepo)

	userID := int64(1)

	userRepo.On("Delete", userID).Return(nil)

	err := service.Delete(userID)
	assert.NoError(t, err)
}

func TestDeleteUser_Fail(t *testing.T) {
	userRepo := new(MockUserRepo)
	service := NewUserService(userRepo)

	userID := int64(1)

	userRepo.On("Delete", userID).Return(errors.New("failed delete"))

	err := service.Delete(userID)
	assert.Error(t, err)
}
