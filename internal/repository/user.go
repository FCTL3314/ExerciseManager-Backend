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

func (ur *UserRepository) GetById(id int64) (*domain.User, error) {
	return ur.Get(&domain.FilterParams{
		Query: "id = ?",
		Args:  []interface{}{id},
	})
}

func (ur *UserRepository) Get(filterParams *domain.FilterParams) (*domain.User, error) {
	var user domain.User
	if err := (ur.db.Where(filterParams.Query, filterParams.Args...).First(&user)).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepository) GetByUsername(username string) (*domain.User, error) {
	return ur.Get(&domain.FilterParams{
		Query: "username = ?",
		Args:  []interface{}{username},
	})
}

func (ur *UserRepository) Fetch(params *domain.Params) ([]*domain.User, error) {
	var users []*domain.User
	query := ur.db.Where(params.Filter.Query, params.Filter.Args...)
	query = query.Order(params.Order)
	if params.Pagination.Limit != 0 {
		query = query.Limit(params.Pagination.Limit).Offset(params.Pagination.Offset)
	}
	if err := (query.Find(&users)).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (ur *UserRepository) Create(user *domain.User) (*domain.User, error) {
	if err := (ur.db.Save(&user)).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *UserRepository) Update(user *domain.User) (*domain.User, error) {
	if err := (ur.db.Save(&user)).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *UserRepository) Delete(id int64) error {
	if err := (ur.db.Where("id = ?", id).Delete(&domain.User{})).Error; err != nil {
		return err
	}
	return nil
}

func (ur *UserRepository) Count(params *domain.FilterParams) (int64, error) {
	var count int64
	if err := (ur.db.Model(&domain.User{}).Where(params.Query, params.Args...).Count(&count)).Error; err != nil {
		return 0, err
	}
	return count, nil
}
