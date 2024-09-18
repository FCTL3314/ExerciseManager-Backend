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
	query := wr.db.Where(params.Query, params.Args...)
	query = query.Preload("User").Preload("Exercises")
	err := (query.First(&workout)).Error
	if err != nil {
		return nil, err
	}

	return &workout, nil
}

func (wr *WorkoutRepository) Fetch(params *domain.Params) ([]*domain.Workout, error) {
	var workouts []*domain.Workout
	query := wr.db.Where(params.Filter.Query, params.Filter.Args...)
	query = query.Order(params.Order)
	query = query.Preload("User").Preload("Exercises")
	if params.Pagination.Limit != 0 {
		query = query.Limit(params.Pagination.Limit).Offset(params.Pagination.Offset)
	}
	if err := (query.Find(&workouts)).Error; err != nil {
		return nil, err
	}

	return workouts, nil
}

func (wr *WorkoutRepository) Create(workout *domain.Workout) (*domain.Workout, error) {
	if err := (wr.db.Save(&workout)).Error; err != nil {
		return nil, err
	}

	if err := wr.db.Preload("User").Preload("Exercises").First(workout).Error; err != nil {
		return nil, err
	}

	return workout, nil
}

func (wr *WorkoutRepository) Update(workout *domain.Workout) (*domain.Workout, error) {
	if err := (wr.db.Save(&workout)).Error; err != nil {
		return nil, err
	}

	if err := wr.db.Preload("User").Preload("Exercises").First(workout).Error; err != nil {
		return nil, err
	}

	return workout, nil
}

func (wr *WorkoutRepository) Delete(id int64) error {
	if err := (wr.db.Where("id = ?", id).Delete(&domain.Workout{})).Error; err != nil {
		return err
	}
	return nil
}

func (wr *WorkoutRepository) Count(params *domain.FilterParams) (int64, error) {
	var count int64
	if err := (wr.db.Model(&domain.Workout{}).Where(params.Query, params.Args...).Count(&count)).Error; err != nil {
		return 0, err
	}
	return count, nil
}
