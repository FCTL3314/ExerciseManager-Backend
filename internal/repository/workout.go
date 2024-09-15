package repository

import (
	"ExerciseManager/internal/domain"
	"gorm.io/gorm"
)

type WorkoutRepository struct {
	db *gorm.DB
}

func NewWorkoutRepository(db *gorm.DB) *WorkoutRepository {
	return &WorkoutRepository{db: db}
}

func (wr *WorkoutRepository) GetById(id int64) (*domain.Workout, error) {
	return wr.Get(&domain.FilterParams{
		Query: "id = ?",
		Args:  []interface{}{id},
	})
}

func (wr *WorkoutRepository) Get(params *domain.FilterParams) (*domain.Workout, error) {
	var workout domain.Workout
	query := wr.db.Where(params.Query, params.Args...).First(&workout)
	if query.Error != nil {
		return nil, query.Error
	}
	return &workout, nil
}

func (wr *WorkoutRepository) Fetch(params *domain.Params) ([]*domain.Workout, error) {
	var workouts []*domain.Workout
	query := wr.db.Where(params.Filter.Query, params.Filter.Args...)
	query = query.Order(params.Order)
	if params.Pagination.Limit != 0 {
		query = query.Limit(params.Pagination.Limit).Offset(params.Pagination.Offset)
	}
	query = query.Find(&workouts)
	if query.Error != nil {
		return nil, query.Error
	}
	return workouts, nil
}

func (wr *WorkoutRepository) Create(workout *domain.Workout) (*domain.Workout, error) {
	query := wr.db.Save(&workout)
	if query.Error != nil {
		return nil, query.Error
	}
	return workout, nil
}

func (wr *WorkoutRepository) Update(workout *domain.Workout) (*domain.Workout, error) {
	query := wr.db.Save(&workout)
	if query.Error != nil {
		return nil, query.Error
	}
	return workout, nil
}

func (wr *WorkoutRepository) Delete(id int64) error {
	if query := wr.db.Where("id = ?", id).Delete(&domain.Workout{}); query.Error != nil {
		return query.Error
	}
	return nil
}

func (wr *WorkoutRepository) Count(params *domain.FilterParams) (int64, error) {
	var count int64
	query := wr.db.Model(&domain.Workout{}).Where(params.Query, params.Args...).Count(&count)
	if query.Error != nil {
		return 0, query.Error
	}
	return count, nil
}
