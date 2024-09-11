package repository

import (
	"ExerciseManager/internal/domain"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) Get(query interface{}, args ...interface{}) (*domain.User, error) {
	var user domain.User
	result := ur.db.Where(query, args...).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (ur *UserRepository) GetByUsername(username string) (*domain.User, error) {
	return ur.Get("username = ?", username)
}

func (ur *UserRepository) List(query interface{}, args ...interface{}) ([]*domain.User, error) {
	var users []*domain.User
	result := ur.db.Where(query, args...).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (ur *UserRepository) Create(user *domain.User) (*domain.User, error) {
	result := ur.db.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (ur *UserRepository) Delete(id uint) error {
	if result := ur.db.Where("id = ?", id).Delete(&domain.User{}); result.Error != nil {
		return result.Error
	}
	return nil
}
