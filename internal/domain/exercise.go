package domain

import (
	"time"
)

type Exercise struct {
	ID          int64
	Name        string
	Description string
	Duration    time.Duration
	Image       *string
	UserID      int64
	User        *User
	Workouts    []*Workout `gorm:"many2many:workout_exercises;"`
}

type ResponseExercise struct {
	ID          int64         `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Duration    time.Duration `json:"duration"`
	Image       *string       `json:"image"`
	User        *ResponseUser `json:"user"`
}

type ResponseNestedExercise struct {
	ID          int64         `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Duration    time.Duration `json:"duration"`
	Image       *string       `json:"image"`
}

type CreateExerciseRequest struct {
	Name        string        `json:"name" binding:"required,min=1,max=128"`
	Description string        `json:"description" binding:"omitempty,min=1,max=256"`
	Duration    time.Duration `json:"duration" binding:"required"`
	Image       *string       `json:"image" binding:"omitempty,url"`
}

type UpdateExerciseRequest struct {
	Name        *string        `json:"name,omitempty" binding:"omitempty,min=1,max=128"`
	Description *string        `json:"description,omitempty" binding:"omitempty,min=1,max=256"`
	Duration    *time.Duration `json:"duration,omitempty" binding:"omitempty"`
	Image       *string        `json:"image" binding:"omitempty,url"`
}

func NewExerciseFromCreateRequest(req *CreateExerciseRequest) *Exercise {
	return &Exercise{
		Name:        req.Name,
		Description: req.Description,
		Duration:    req.Duration,
		Image:       req.Image,
	}
}

func (e *Exercise) ToResponseExercise() *ResponseExercise {
	re := &ResponseExercise{
		ID:          e.ID,
		Name:        e.Name,
		Description: e.Description,
		Duration:    e.Duration,
		Image:       e.Image,
	}

	if e.User != nil {
		re.User = e.User.ToResponseUser()
	}
	return re
}

func (e *Exercise) ToResponseNestedExercise() *ResponseNestedExercise {
	return &ResponseNestedExercise{
		ID:          e.ID,
		Name:        e.Name,
		Description: e.Description,
		Duration:    e.Duration,
		Image:       e.Image,
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

	if req.Image != nil {
		e.Image = req.Image
	}
}

func ToResponseExercises(exercises []*Exercise) []*ResponseExercise {
	responseExercises := make([]*ResponseExercise, len(exercises))
	for i, exercise := range exercises {
		responseExercises[i] = exercise.ToResponseExercise()
	}
	return responseExercises
}
