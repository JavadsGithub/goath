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

func (ur *UserRepository) FindAll() ([]*models.User, error) {
	users := make([]*models.User, 0)
	if err := ur.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
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

// FIXME: should not get a user model to save! it should be an INPUT!
func (ur *UserRepository) Save(user *models.User) error {
	// FIXME: not a repository logic!
	hasedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hasedPassword
	return ur.db.Create(user).Error
}
