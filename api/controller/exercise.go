package controller

import (
	"ExerciseManager/bootstrap"
	"ExerciseManager/internal/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DefaultExerciseController struct {
	usecase      domain.ExerciseUsecase
	errorHandler *ErrorHandler
	Logger       bootstrap.Logger
	cfg          *bootstrap.Config
}

func NewExerciseController(
	usecase domain.ExerciseUsecase,
	errorHandler *ErrorHandler,
	logger bootstrap.Logger,
	cfg *bootstrap.Config,
) *DefaultExerciseController {
	return &DefaultExerciseController{
		usecase:      usecase,
		errorHandler: errorHandler,
		Logger:       logger,
		cfg:          cfg,
	}
}

func (wc *DefaultExerciseController) Get(c *gin.Context) {
	id, err := getParamAsInt64(c, "id")
	if err != nil {
		wc.errorHandler.Handle(c, err)
		return
	}

	exercise, err := wc.usecase.GetById(id)

	if err != nil {
		wc.errorHandler.Handle(c, err)
		return
	}

	responseExercise := exercise.ToResponseExercise()

	c.JSON(http.StatusOK, responseExercise)
}

func (wc *DefaultExerciseController) List(c *gin.Context) {
	paginationParams, err := getPaginationParams(c, wc.cfg.Pagination.MaxExerciseLimit)
	if err != nil {
		wc.errorHandler.Handle(c, err)
		return
	}

	params := domain.Params{
		Pagination: paginationParams,
	}

	paginatedResult, err := wc.usecase.List(&params)
	if err != nil {
		wc.errorHandler.Handle(c, err)
		return
	}

	responseExercises := domain.ToResponseExercises(paginatedResult.Results)

	paginatedResponse := domain.PaginatedResponse{
		Count:   paginatedResult.Count,
		Limit:   paginationParams.Limit,
		Offset:  paginationParams.Offset,
		Results: responseExercises,
	}

	c.JSON(http.StatusOK, paginatedResponse)
}

func (wc *DefaultExerciseController) Create(c *gin.Context) {
	var exercise domain.CreateExerciseRequest
	if err := c.ShouldBindJSON(&exercise); err != nil {
		c.JSON(http.StatusBadRequest, domain.NewValidationErrorResponse(err.Error()))
		return
	}

	authUserId := c.GetInt64(string(UserIDContextKey))

	createdExercise, err := wc.usecase.Create(authUserId, &exercise)
	if err != nil {
		wc.errorHandler.Handle(c, err)
		return
	}

	responseExercise := createdExercise.ToResponseExercise()

	c.JSON(http.StatusCreated, responseExercise)
}

func (wc *DefaultExerciseController) Update(c *gin.Context) {
	id, err := getParamAsInt64(c, "id")
	if err != nil {
		wc.errorHandler.Handle(c, err)
		return
	}

	var exercise domain.UpdateExerciseRequest
	if err := c.ShouldBindJSON(&exercise); err != nil {
		c.JSON(http.StatusBadRequest, domain.NewValidationErrorResponse(err.Error()))
		return
	}

	authUserId := c.GetInt64(string(UserIDContextKey))

	updatedExercise, err := wc.usecase.Update(authUserId, id, &exercise)
	if err != nil {
		wc.errorHandler.Handle(c, err)
		return
	}

	responseExercise := updatedExercise.ToResponseExercise()

	c.JSON(http.StatusOK, responseExercise)
}

func (wc *DefaultExerciseController) Delete(c *gin.Context) {
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
