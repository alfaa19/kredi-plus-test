package service

import (
	"github.com/alfaa19/kredi-plus-test/internal/model"
	"github.com/alfaa19/kredi-plus-test/internal/repository"
)

type UserService interface {
	Create(user *model.User) error
	GetByID(id int64) (*model.User, error)
	Update(user *model.User) error
	Delete(id int64) error
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) Create(user *model.User) error {
	return s.userRepo.Create(user)
}

func (s *userService) GetByID(id int64) (*model.User, error) {
	return s.userRepo.GetByID(id)
}

func (s *userService) Update(user *model.User) error {
	return s.userRepo.Update(user)
}

func (s *userService) Delete(id int64) error {
	return s.userRepo.Delete(id)
}
