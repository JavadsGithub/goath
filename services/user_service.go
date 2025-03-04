package services

import (
	"github.com/JavadsGithub/goath/models"
	"github.com/JavadsGithub/goath/repositories"
)

type UserService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (us *UserService) GetUserByID(id string) (*models.User, error) {
	return us.userRepo.FindByID(id)
}

func (us *UserService) CreateUser(user *models.User) error {
	return us.userRepo.Save(user)
}
