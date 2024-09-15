package domain

import (
	"time"
)

type Workout struct {
	ID          int64
	Name        string
	Description string
	UserID      int64
	User        *User
	Exercises   []*Exercise `gorm:"many2many:workout_exercises;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type CreateWorkoutRequest struct {
	Name        string `json:"name" binding:"required,min=1,max=128"`
	Description string `json:"description" binding:"required,min=1,max=256"`
}

type UpdateWorkoutRequest struct {
	Name        *string `json:"name" binding:"required,min=1,max=128"`
	Description *string `json:"description" binding:"required,min=1,max=256"`
}

func NewWorkoutFromCreateRequest(req *CreateWorkoutRequest) *Workout {
	return &Workout{
		Name:        req.Name,
		Description: req.Description,
	}
}

func (w *Workout) ApplyUpdate(req *UpdateWorkoutRequest) {
	if req.Name != nil {
		w.Name = *req.Name
	}

	if req.Description != nil {
		w.Description = *req.Description
	}
}
