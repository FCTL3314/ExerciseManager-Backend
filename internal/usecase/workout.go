package usecase

import (
	"ExerciseManager/bootstrap"
	"ExerciseManager/internal/domain"
	"ExerciseManager/internal/errormapper"
	"ExerciseManager/internal/permission"
	"reflect"
)

type WorkoutUsecase struct {
	workoutRepository         domain.WorkoutRepository
	exerciseRepository        domain.ExerciseRepository
	workoutExerciseRepository domain.WorkoutExerciseRepository
	accessManager             permission.AccessPolicy
	errorMapper               errormapper.Chain
	cfg                       *bootstrap.Config
}

func NewWorkoutUsecase(
	workoutRepository domain.WorkoutRepository,
	exerciseRepository domain.ExerciseRepository,
	workoutExerciseRepository domain.WorkoutExerciseRepository,
	accessManager permission.AccessPolicy,
	errorMapper errormapper.Chain,
	cfg *bootstrap.Config,
) *WorkoutUsecase {
	return &WorkoutUsecase{
		workoutRepository:         workoutRepository,
		exerciseRepository:        exerciseRepository,
		workoutExerciseRepository: workoutExerciseRepository,
		accessManager:             accessManager,
		errorMapper:               errorMapper,
		cfg:                       cfg,
	}
}

func (wu *WorkoutUsecase) GetById(id int64) (*domain.Workout, error) {
	workout, err := wu.workoutRepository.GetById(id)
	if err != nil {
		return nil, wu.errorMapper.MapError(err)
	}

	return workout, nil
}

func (wu *WorkoutUsecase) Get(params *domain.FilterParams) (*domain.Workout, error) {
	workout, err := wu.workoutRepository.Get(params)
	if err != nil {
		return nil, wu.errorMapper.MapError(err)
	}
	return workout, nil
}

func (wu *WorkoutUsecase) List(params *domain.Params) (*domain.PaginatedResult[*domain.Workout], error) {
	workouts, err := wu.workoutRepository.Fetch(params)
	if err != nil {
		return nil, wu.errorMapper.MapError(err)
	}

	count, err := wu.workoutRepository.Count(&domain.FilterParams{})
	if err != nil {
		return nil, wu.errorMapper.MapError(err)
	}

	return &domain.PaginatedResult[*domain.Workout]{Results: workouts, Count: count}, nil
}

func (wu *WorkoutUsecase) Create(authUserId int64, createWorkoutRequest *domain.CreateWorkoutRequest) (*domain.Workout, error) {
	workout := domain.NewWorkoutFromCreateRequest(createWorkoutRequest)
	workout.UserID = authUserId
	return wu.workoutRepository.Create(workout)
}

func (wu *WorkoutUsecase) AddExercise(authUserId, workoutId int64, addExerciseRequest *domain.AddExerciseToWorkoutRequest) (*domain.Workout, error) {
	workout, err := wu.workoutRepository.GetById(workoutId)
	if err != nil {
		return nil, wu.errorMapper.MapError(err)
	}

	if !wu.accessManager.HasAccess(authUserId, workout) {
		return nil, domain.ErrAccessDenied
	}

	if len(workout.WorkoutExercises) >= wu.cfg.Workout.MaxExercisesCount {
		return nil, &domain.ErrMaxRelatedObjectsNumberReached{
			ParentObjectName:  reflect.TypeOf(domain.Workout{}).Name(),
			RelatedObjectName: reflect.TypeOf(domain.Exercise{}).Name(),
			Limit:             wu.cfg.Workout.MaxExercisesCount,
		}
	}

	exercise, err := wu.exerciseRepository.GetById(addExerciseRequest.ExerciseID)
	if err != nil {
		return nil, wu.errorMapper.MapError(err)
	}

	workoutExercise := domain.WorkoutExercise{
		WorkoutID:  workout.ID,
		ExerciseID: exercise.ID,
		BreakTime:  addExerciseRequest.BreakTime,
	}
	_, err = wu.workoutExerciseRepository.Create(&workoutExercise)
	if err != nil {
		return nil, wu.errorMapper.MapError(err)
	}

	workout, err = wu.workoutRepository.GetById(workoutId)
	if err != nil {
		return nil, wu.errorMapper.MapError(err)
	}

	return workout, nil
}

func (wu *WorkoutUsecase) UpdateExercise(authUserId, workoutId, workoutExerciseId int64, updateWorkoutExerciseRequest *domain.UpdateWorkoutExerciseRequest) (*domain.Workout, error) {
	workoutExerciseToUpdate, err := wu.workoutExerciseRepository.GetById(workoutExerciseId)
	if err != nil {
		return nil, wu.errorMapper.MapError(err)
	}

	if workoutId != workoutExerciseToUpdate.WorkoutID {
		return nil, domain.ErrObjectNotFound
	}

	if !wu.accessManager.HasAccess(authUserId, workoutExerciseToUpdate.Workout) {
		return nil, domain.ErrAccessDenied
	}

	workoutExerciseToUpdate.ApplyUpdate(updateWorkoutExerciseRequest)
	if _, err := wu.workoutExerciseRepository.Update(workoutExerciseToUpdate); err != nil {
		return nil, wu.errorMapper.MapError(err)
	}

	workout, err := wu.workoutRepository.GetById(workoutId)
	if err != nil {
		return nil, wu.errorMapper.MapError(err)
	}

	return workout, nil
}

func (wu *WorkoutUsecase) RemoveExercise(authUserId, workoutId, workoutExerciseId int64) (*domain.Workout, error) {
	workout, err := wu.workoutRepository.GetById(workoutId)
	if err != nil {
		return nil, wu.errorMapper.MapError(err)
	}

	if !wu.accessManager.HasAccess(authUserId, workout) {
		return nil, domain.ErrAccessDenied
	}

	err = wu.workoutExerciseRepository.Delete(workoutExerciseId)
	if err != nil {
		return nil, wu.errorMapper.MapError(err)
	}

	workout, err = wu.workoutRepository.GetById(workoutId)
	if err != nil {
		return nil, wu.errorMapper.MapError(err)
	}

	return workout, nil
}

func (wu *WorkoutUsecase) Update(authUserId int64, id int64, updateWorkoutRequest *domain.UpdateWorkoutRequest) (*domain.Workout, error) {
	workoutToUpdate, err := wu.workoutRepository.GetById(id)
	if err != nil {
		return nil, wu.errorMapper.MapError(err)
	}

	if !wu.accessManager.HasAccess(authUserId, workoutToUpdate) {
		return nil, domain.ErrAccessDenied
	}

	workoutToUpdate.ApplyUpdate(updateWorkoutRequest)

	updatedWorkout, err := wu.workoutRepository.Update(workoutToUpdate)
	if err != nil {
		return nil, wu.errorMapper.MapError(err)
	}
	return updatedWorkout, nil
}

func (wu *WorkoutUsecase) Delete(authUserId int64, id int64) error {
	workout, err := wu.workoutRepository.GetById(id)
	if err != nil {
		return wu.errorMapper.MapError(err)
	}

	if !wu.accessManager.HasAccess(authUserId, workout) {
		return domain.ErrAccessDenied
	}

	return wu.workoutRepository.Delete(id)
}
