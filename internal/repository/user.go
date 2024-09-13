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

func (ur *UserRepository) Get(params *domain.FilterParams) (*domain.User, error) {
	var user domain.User
	query := ur.db.Where(params.Query, params.Args...).First(&user)
	if query.Error != nil {
		return nil, query.Error
	}
	return &user, nil
}

func (ur *UserRepository) GetByUsername(username string) (*domain.User, error) {
	return ur.Get(&domain.FilterParams{
		Query: "username = ?",
		Args:  []interface{}{username},
	})
}

func (ur *UserRepository) List(params *domain.Params) ([]*domain.User, error) {
	var users []*domain.User
	query := ur.db.Where(params.Filter.Query, params.Filter.Args...)
	query = query.Order(params.Order)
	if params.Pagination.Limit != 0 {
		query = query.Limit(params.Pagination.Limit).Offset(params.Pagination.Offset)
	}
	query = query.Find(&users)
	if query.Error != nil {
		return nil, query.Error
	}
	return users, nil
}

func (ur *UserRepository) Create(user *domain.User) (*domain.User, error) {
	query := ur.db.Save(&user)
	if query.Error != nil {
		return nil, query.Error
	}
	return user, nil
}

func (ur *UserRepository) Update(user *domain.User) (*domain.User, error) {
	query := ur.db.Create(&user)
	if query.Error != nil {
		return nil, query.Error
	}
	return user, nil
}

func (ur *UserRepository) Delete(id uint) error {
	if query := ur.db.Where("id = ?", id).Delete(&domain.User{}); query.Error != nil {
		return query.Error
	}
	return nil
}
