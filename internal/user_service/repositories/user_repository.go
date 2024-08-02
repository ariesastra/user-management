package repositories

import (
	usermodel "aegis_task/internal/user_service/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user usermodel.User) error
	FindByID(id uint) (*usermodel.User, error)
	Update(user usermodel.User) error
	Delete(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(user usermodel.User) error {
	return r.db.Create(&user).Error
}

func (r *userRepository) FindByID(id uint) (*usermodel.User, error) {
	var user usermodel.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Update(user usermodel.User) error {
	return r.db.Save(&user).Error
}

func (r *userRepository) Delete(id uint) error {
	return r.db.Delete(&usermodel.User{}, id).Error
}
