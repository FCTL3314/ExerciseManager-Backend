package usecase

import (
	"ExerciseManager/internal/domain"
	"ExerciseManager/internal/errormapper"
	"ExerciseManager/internal/permission"
)

type WorkoutUsecase struct {
	workoutRepository domain.WorkoutRepository
	accessManager     permission.AccessPolicy
	errorMapper       errormapper.Chain
}

func NewWorkoutUsecase(
	workoutRepository domain.WorkoutRepository,
	accessManager permission.AccessPolicy,
	errorMapper errormapper.Chain,
) *WorkoutUsecase {
	return &WorkoutUsecase{
		workoutRepository: workoutRepository,
		accessManager:     accessManager,
		errorMapper:       errorMapper,
	}
}

func (ur *WorkoutUsecase) GetById(id int64) (*domain.Workout, error) {
	workout, err := ur.workoutRepository.GetById(id)
	if err != nil {
		return nil, ur.errorMapper.MapError(err)
	}

	return workout, nil
}

func (ur *WorkoutUsecase) Get(params *domain.FilterParams) (*domain.Workout, error) {
	workout, err := ur.workoutRepository.Get(params)
	if err != nil {
		return nil, ur.errorMapper.MapError(err)
	}
	return workout, nil
}

func (ur *WorkoutUsecase) List(params *domain.Params) (*domain.PaginatedResult[*domain.Workout], error) {
	workouts, err := ur.workoutRepository.Fetch(params)
	if err != nil {
		return nil, ur.errorMapper.MapError(err)
	}

	count, err := ur.workoutRepository.Count(&domain.FilterParams{})
	if err != nil {
		return nil, ur.errorMapper.MapError(err)
	}

	return &domain.PaginatedResult[*domain.Workout]{Results: workouts, Count: count}, nil
}

func (ur *WorkoutUsecase) Create(createWorkoutRequest *domain.CreateWorkoutRequest) (*domain.Workout, error) {
	workout := domain.NewWorkoutFromCreateRequest(createWorkoutRequest)
	return ur.workoutRepository.Create(workout)
}

func (ur *WorkoutUsecase) Update(authUserId int64, id int64, updateWorkoutRequest *domain.UpdateWorkoutRequest) (*domain.Workout, error) {
	workoutToUpdate, err := ur.workoutRepository.GetById(id)
	if err != nil {
		return nil, ur.errorMapper.MapError(err)
	}

	if !ur.accessManager.HasAccess(authUserId, workoutToUpdate) {
		return nil, domain.ErrAccessDenied
	}

	workoutToUpdate.ApplyUpdate(updateWorkoutRequest)

	updatedWorkout, err := ur.workoutRepository.Update(workoutToUpdate)
	if err != nil {
		return nil, ur.errorMapper.MapError(err)
	}
	return updatedWorkout, nil
}

func (ur *WorkoutUsecase) Delete(authUserId int64, id int64) error {
	workout, err := ur.workoutRepository.GetById(id)
	if err != nil {
		return ur.errorMapper.MapError(err)
	}

	if !ur.accessManager.HasAccess(authUserId, workout) {
		return domain.ErrAccessDenied
	}

	return ur.workoutRepository.Delete(id)
}
