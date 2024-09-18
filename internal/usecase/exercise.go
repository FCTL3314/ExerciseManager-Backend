package usecase

import (
	"ExerciseManager/internal/domain"
	"ExerciseManager/internal/errormapper"
	"ExerciseManager/internal/permission"
)

type ExerciseUsecase struct {
	exerciseRepository domain.ExerciseRepository
	accessManager      permission.AccessPolicy
	errorMapper        errormapper.Chain
}

func NewExerciseUsecase(
	exerciseRepository domain.ExerciseRepository,
	accessManager permission.AccessPolicy,
	errorMapper errormapper.Chain,
) *ExerciseUsecase {
	return &ExerciseUsecase{
		exerciseRepository: exerciseRepository,
		accessManager:      accessManager,
		errorMapper:        errorMapper,
	}
}

func (eu *ExerciseUsecase) GetById(id int64) (*domain.Exercise, error) {
	exercise, err := eu.exerciseRepository.GetById(id)
	if err != nil {
		return nil, eu.errorMapper.MapError(err)
	}

	return exercise, nil
}

func (eu *ExerciseUsecase) Get(params *domain.FilterParams) (*domain.Exercise, error) {
	exercise, err := eu.exerciseRepository.Get(params)
	if err != nil {
		return nil, eu.errorMapper.MapError(err)
	}
	return exercise, nil
}

func (eu *ExerciseUsecase) List(params *domain.Params) (*domain.PaginatedResult[*domain.Exercise], error) {
	exercises, err := eu.exerciseRepository.Fetch(params)
	if err != nil {
		return nil, eu.errorMapper.MapError(err)
	}

	count, err := eu.exerciseRepository.Count(&domain.FilterParams{})
	if err != nil {
		return nil, eu.errorMapper.MapError(err)
	}

	return &domain.PaginatedResult[*domain.Exercise]{Results: exercises, Count: count}, nil
}

func (eu *ExerciseUsecase) Create(authUserId int64, createExerciseRequest *domain.CreateExerciseRequest) (*domain.Exercise, error) {
	exercise := domain.NewExerciseFromCreateRequest(createExerciseRequest)
	exercise.UserID = authUserId
	return eu.exerciseRepository.Create(exercise)
}

func (eu *ExerciseUsecase) Update(authUserId int64, id int64, updateExerciseRequest *domain.UpdateExerciseRequest) (*domain.Exercise, error) {
	exerciseToUpdate, err := eu.exerciseRepository.GetById(id)
	if err != nil {
		return nil, eu.errorMapper.MapError(err)
	}

	if !eu.accessManager.HasAccess(authUserId, exerciseToUpdate) {
		return nil, domain.ErrAccessDenied
	}

	exerciseToUpdate.ApplyUpdate(updateExerciseRequest)

	updatedExercise, err := eu.exerciseRepository.Update(exerciseToUpdate)
	if err != nil {
		return nil, eu.errorMapper.MapError(err)
	}
	return updatedExercise, nil
}

func (eu *ExerciseUsecase) Delete(authUserId int64, id int64) error {
	exercise, err := eu.exerciseRepository.GetById(id)
	if err != nil {
		return eu.errorMapper.MapError(err)
	}

	if !eu.accessManager.HasAccess(authUserId, exercise) {
		return domain.ErrAccessDenied
	}

	return eu.exerciseRepository.Delete(id)
}
