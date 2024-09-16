package usecase

import (
	"ExerciseManager/internal/domain"
	"errors"
	"gorm.io/gorm"
)

type WorkoutUsecase struct {
	workoutRepository domain.WorkoutRepository
}

func NewWorkoutUsecase(
	workoutRepository domain.WorkoutRepository,
) *WorkoutUsecase {
	return &WorkoutUsecase{
		workoutRepository: workoutRepository,
	}
}

func (ur *WorkoutUsecase) GetById(id int64) (*domain.Workout, error) {
	workout, err := ur.workoutRepository.GetById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrObjectNotFound
		}
		return nil, err
	}

	return workout, nil
}

func (ur *WorkoutUsecase) Get(params *domain.FilterParams) (*domain.Workout, error) {
	workout, err := ur.workoutRepository.Get(params)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrObjectNotFound
		}
		return nil, err
	}
	return workout, nil
}

func (ur *WorkoutUsecase) List(params *domain.Params) (*domain.PaginatedResult[*domain.Workout], error) {
	workouts, err := ur.workoutRepository.Fetch(params)
	if err != nil {
		return &domain.PaginatedResult[*domain.Workout]{}, err
	}

	count, err := ur.workoutRepository.Count(&domain.FilterParams{})
	if err != nil {
		return &domain.PaginatedResult[*domain.Workout]{}, err
	}

	return &domain.PaginatedResult[*domain.Workout]{Results: workouts, Count: count}, nil
}

func (ur *WorkoutUsecase) Create(createWorkoutRequest *domain.CreateWorkoutRequest) (*domain.Workout, error) {
	workout := domain.NewWorkoutFromCreateRequest(createWorkoutRequest)
	return ur.workoutRepository.Create(workout)
}

func (ur *WorkoutUsecase) Update(authUserId int64, id int64, updateWorkoutRequest *domain.UpdateWorkoutRequest) (*domain.Workout, error) {
	// if !ur.accessManager.HasAccessToUser(authUserId, id) {
	// 	return nil, domain.ErrAccessDenied
	// }

	userToUpdate, err := ur.workoutRepository.GetById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrObjectNotFound
		}
		return nil, err
	}

	userToUpdate.ApplyUpdate(updateWorkoutRequest)

	return ur.workoutRepository.Update(userToUpdate)
}

func (ur *WorkoutUsecase) Delete(authUserId int64, id int64) error {
	// if !ur.accessManager.HasAccessToUser(authUserId, id) {
	// 	return domain.ErrAccessDenied
	// }

	if _, err := ur.workoutRepository.GetById(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.ErrObjectNotFound
		}
		return err
	}

	return ur.workoutRepository.Delete(id)
}
