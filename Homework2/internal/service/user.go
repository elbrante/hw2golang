package service

import (
	"Homework2/internal/model"
)

type UserRepository interface {
	GetAll() []*model.User
}

func NewUserService(repository UserRepository) *UserService {
	return &UserService{userRepository: repository}
}

type UserService struct {
	userRepository UserRepository
}

func (service *UserService) GetAll() []*model.User {
	return service.userRepository.GetAll()
}
