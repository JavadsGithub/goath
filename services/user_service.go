package services

import (
	"github.com/JavadsGithub/goath/models"
	"github.com/JavadsGithub/goath/repositories"
)

// FIXME: should get an interface!
// HINT: should not be aware of repository
type UserService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (us *UserService) GetUserByID(id string) (*models.User, error) {
	return us.userRepo.FindByID(id)
}

func (us *UserService) GetUserByEmail(email string) (*models.User, error) {
	return us.userRepo.FindByEmail(email)
}

func (us *UserService) GetAllUsers() ([]*models.User, error) {
	return us.userRepo.FindAll()
}

func (us *UserService) CreateUser(user *models.User) error {
	// !
	return us.userRepo.Save(user)
}
