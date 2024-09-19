package repository

import (
	"ExerciseManager/internal/domain"
	"gorm.io/gorm"
)

type ExerciseRepository struct {
	db        *gorm.DB
	toPreload []string
}

func NewExerciseRepository(db *gorm.DB) *ExerciseRepository {
	return &ExerciseRepository{db: db, toPreload: []string{"User", "Workouts"}}
}

func (er *ExerciseRepository) GetById(id int64) (*domain.Exercise, error) {
	return er.Get(&domain.FilterParams{
		Query: "id = ?",
		Args:  []interface{}{id},
	})
}

func (er *ExerciseRepository) Get(filterParams *domain.FilterParams) (*domain.Exercise, error) {
	var exercise domain.Exercise
	query := er.db.Where(filterParams.Query, filterParams.Args...)
	query = applyPreloadsForGORMQuery(query, er.toPreload)
	if err := (query.First(&exercise)).Error; err != nil {
		return nil, err
	}

	return &exercise, nil
}

func (er *ExerciseRepository) Fetch(params *domain.Params) ([]*domain.Exercise, error) {
	var exercises []*domain.Exercise
	query := er.db.Where(params.Filter.Query, params.Filter.Args...)
	query = query.Order(params.Order)
	query = applyPreloadsForGORMQuery(query, er.toPreload)
	if params.Pagination.Limit != 0 {
		query = query.Limit(params.Pagination.Limit).Offset(params.Pagination.Offset)
	}
	if err := (query.Find(&exercises)).Error; err != nil {
		return nil, err
	}

	return exercises, nil
}

func (er *ExerciseRepository) Create(exercise *domain.Exercise) (*domain.Exercise, error) {
	if err := (er.db.Save(&exercise)).Error; err != nil {
		return nil, err
	}

	query := applyPreloadsForGORMQuery(er.db.Model(&domain.Exercise{}), er.toPreload)
	if err := query.First(exercise).Error; err != nil {
		return nil, err
	}

	return exercise, nil
}

func (er *ExerciseRepository) Update(exercise *domain.Exercise) (*domain.Exercise, error) {
	if err := (er.db.Save(&exercise)).Error; err != nil {
		return nil, err
	}

	query := applyPreloadsForGORMQuery(er.db.Model(&domain.Exercise{}), er.toPreload)
	if err := query.First(exercise).Error; err != nil {
		return nil, err
	}

	return exercise, nil
}

func (er *ExerciseRepository) Delete(id int64) error {
	result := er.db.Where("id = ?", id).Delete(&domain.Exercise{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (er *ExerciseRepository) Count(params *domain.FilterParams) (int64, error) {
	var count int64
	if err := (er.db.Model(&domain.Exercise{}).Where(params.Query, params.Args...).Count(&count)).Error; err != nil {
		return 0, err
	}
	return count, nil
}
