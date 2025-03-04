package repositories

import (
	"github.com/JavadsGithub/goath/models"
	"github.com/JavadsGithub/goath/utils"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) FindByID(id string) (*models.User, error) {
	var user models.User
	if err := ur.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := ur.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepository) Save(user *models.User) error {
	hasedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hasedPassword
	return ur.db.Create(user).Error
}
