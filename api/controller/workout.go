package controller

import (
	"ExerciseManager/bootstrap"
	"ExerciseManager/internal/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DefaultWorkoutController struct {
	usecase      domain.WorkoutUsecase
	errorHandler *ErrorHandler
	Logger       bootstrap.Logger
	cfg          *bootstrap.Config
}

func NewWorkoutController(
	usecase domain.WorkoutUsecase,
	errorHandler *ErrorHandler,
	logger bootstrap.Logger,
	cfg *bootstrap.Config,
) *DefaultWorkoutController {
	return &DefaultWorkoutController{
		usecase:      usecase,
		errorHandler: errorHandler,
		Logger:       logger,
		cfg:          cfg,
	}
}

func (wc *DefaultWorkoutController) Get(c *gin.Context) {
	id, err := getParamAsInt64(c, "id")
	if err != nil {
		wc.errorHandler.Handle(c, err)
		return
	}

	workout, err := wc.usecase.GetById(id)

	if err != nil {
		wc.errorHandler.Handle(c, err)
		return
	}

	responseWorkout := workout.ToResponseWorkout()

	c.JSON(http.StatusOK, responseWorkout)
}

func (wc *DefaultWorkoutController) List(c *gin.Context) {
	params, err := getParams(c, wc.cfg.Pagination.MaxWorkoutLimit)
	if err != nil {
		wc.errorHandler.Handle(c, err)
		return
	}

	paginatedResult, err := wc.usecase.List(&params)
	if err != nil {
		wc.errorHandler.Handle(c, err)
		return
	}

	responseWorkouts := domain.ToResponseWorkouts(paginatedResult.Results)

	paginatedResponse := domain.PaginatedResponse{
		Count:   paginatedResult.Count,
		Limit:   params.Pagination.Limit,
		Offset:  params.Pagination.Offset,
		Results: responseWorkouts,
	}

	c.JSON(http.StatusOK, paginatedResponse)
}

func (wc *DefaultWorkoutController) Create(c *gin.Context) {
	var workout domain.CreateWorkoutRequest
	if err := c.ShouldBindJSON(&workout); err != nil {
		c.JSON(http.StatusBadRequest, domain.NewValidationErrorResponse(err.Error()))
		return
	}

	authUserId := c.GetInt64(string(UserIDContextKey))

	createdWorkout, err := wc.usecase.Create(authUserId, &workout)
	if err != nil {
		wc.errorHandler.Handle(c, err)
		return
	}

	responseWorkout := createdWorkout.ToResponseWorkout()

	c.JSON(http.StatusCreated, responseWorkout)
}

func (wc *DefaultWorkoutController) AddExercise(c *gin.Context) {
	workoutId, err := getParamAsInt64(c, "id")
	if err != nil {
		wc.errorHandler.Handle(c, err)
		return
	}

	var addExerciseRequest domain.AddExerciseToWorkoutRequest
	if err := c.ShouldBindJSON(&addExerciseRequest); err != nil {
		c.JSON(http.StatusBadRequest, domain.NewValidationErrorResponse(err.Error()))
		return
	}

	authUserId := c.GetInt64(string(UserIDContextKey))

	workout, err := wc.usecase.AddExercise(authUserId, workoutId, &addExerciseRequest)
	if err != nil {
		wc.errorHandler.Handle(c, err)
		return
	}

	responseWorkout := workout.ToResponseWorkout()

	c.JSON(http.StatusCreated, responseWorkout)
}

func (wc *DefaultWorkoutController) UpdateExercise(c *gin.Context) {
	workoutId, err := getParamAsInt64(c, "id")
	if err != nil {
		wc.errorHandler.Handle(c, err)
		return
	}

	workoutExerciseId, err := getParamAsInt64(c, "workoutExerciseId")
	if err != nil {
		wc.errorHandler.Handle(c, err)
		return
	}

	var updateWorkoutExerciseRequest domain.UpdateWorkoutExerciseRequest
	if err := c.ShouldBindJSON(&updateWorkoutExerciseRequest); err != nil {
		c.JSON(http.StatusBadRequest, domain.NewValidationErrorResponse(err.Error()))
		return
	}

	authUserId := c.GetInt64(string(UserIDContextKey))

	workoutExercise, err := wc.usecase.UpdateExercise(authUserId, workoutId, workoutExerciseId, &updateWorkoutExerciseRequest)
	if err != nil {
		wc.errorHandler.Handle(c, err)
		return
	}

	responseWorkout := workoutExercise.ToResponseWorkout()

	c.JSON(http.StatusCreated, responseWorkout)
}

func (wc *DefaultWorkoutController) RemoveExercise(c *gin.Context) {
	workoutId, err := getParamAsInt64(c, "id")
	if err != nil {
		wc.errorHandler.Handle(c, err)
		return
	}
	exerciseId, err := getParamAsInt64(c, "exerciseId")
	if err != nil {
		wc.errorHandler.Handle(c, err)
		return
	}

	authUserId := c.GetInt64(string(UserIDContextKey))

	workout, err := wc.usecase.RemoveExercise(authUserId, workoutId, exerciseId)
	if err != nil {
		wc.errorHandler.Handle(c, err)
		return
	}

	responseWorkout := workout.ToResponseWorkout()

	c.JSON(http.StatusOK, responseWorkout)
}

func (wc *DefaultWorkoutController) Update(c *gin.Context) {
	id, err := getParamAsInt64(c, "id")
	if err != nil {
		wc.errorHandler.Handle(c, err)
		return
	}

	var workout domain.UpdateWorkoutRequest
	if err := c.ShouldBindJSON(&workout); err != nil {
		c.JSON(http.StatusBadRequest, domain.NewValidationErrorResponse(err.Error()))
		return
	}

	authUserId := c.GetInt64(string(UserIDContextKey))

	updatedWorkout, err := wc.usecase.Update(authUserId, id, &workout)
	if err != nil {
		wc.errorHandler.Handle(c, err)
		return
	}

	responseWorkout := updatedWorkout.ToResponseWorkout()

	c.JSON(http.StatusOK, responseWorkout)
}

func (wc *DefaultWorkoutController) Delete(c *gin.Context) {
	id, err := getParamAsInt64(c, "id")
	if err != nil {
		wc.errorHandler.Handle(c, err)
		return
	}

	authUserId := c.GetInt64(string(UserIDContextKey))

	if err := wc.usecase.Delete(authUserId, id); err != nil {
		wc.errorHandler.Handle(c, err)
		return
	}
	c.Status(http.StatusNoContent)
}
