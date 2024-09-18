package domain

import (
	"time"
)

type Exercise struct {
	ID          int64
	Name        string
	Description string
	Duration    time.Duration
	UserID      int64
	User        *User
	Workouts    []*Workout `gorm:"many2many:workout_exercises;"`
}

type WorkoutExercises struct {
	WorkoutID  int64
	Workout    *Workout
	ExerciseID int64
	Exercise   *Exercise
	BreakTime  time.Duration
}

type ResponseExercise struct {
	ID          int64              `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Duration    time.Duration      `json:"duration"`
	UserID      int64              `json:"user_id"`
	User        *ResponseUser      `json:"user"`
	Workouts    []*ResponseWorkout `json:"workouts"`
}

type CreateExerciseRequest struct {
	Name        string        `json:"name" binding:"required,min=1,max=128"`
	Description string        `json:"description" binding:"required,min=1,max=256"`
	Duration    time.Duration `json:"duration" binding:"required"`
}

type UpdateExerciseRequest struct {
	Name        *string        `json:"name,omitempty" binding:"omitempty,min=1,max=128"`
	Description *string        `json:"description,omitempty" binding:"omitempty,min=1,max=256"`
	Duration    *time.Duration `json:"duration,omitempty" binding:"omitempty"`
}

func NewExerciseFromCreateRequest(req *CreateExerciseRequest) *Exercise {
	return &Exercise{
		Name:        req.Name,
		Description: req.Description,
		Duration:    req.Duration,
	}
}

func (e *Exercise) ToResponseExercise() *ResponseExercise {
	return &ResponseExercise{
		ID:          e.ID,
		Name:        e.Name,
		Description: e.Description,
		Duration:    e.Duration,
		UserID:      e.UserID,
		User:        e.User.ToResponseUser(),
		Workouts:    ToResponseWorkouts(e.Workouts),
	}
}

func (e *Exercise) ApplyUpdate(req *UpdateExerciseRequest) {
	if req.Name != nil {
		e.Name = *req.Name
	}

	if req.Description != nil {
		e.Description = *req.Description
	}

	if req.Duration != nil {
		e.Duration = *req.Duration
	}
}

func ToResponseExercises(users []*Exercise) []*ResponseExercise {
	responseExercises := make([]*ResponseExercise, len(users))
	for i, workout := range users {
		responseExercises[i] = workout.ToResponseExercise()
	}
	return responseExercises
}
