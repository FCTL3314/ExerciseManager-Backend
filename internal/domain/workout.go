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

type ResponseWorkout struct {
	ID          int64       `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	UserID      int64       `json:"user_id"`
	User        *User       `json:"user"`
	Exercises   []*Exercise `json:"exercises"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
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

func (w *Workout) ToResponseWorkout() *ResponseWorkout {
	return &ResponseWorkout{
		ID:          w.ID,
		Name:        w.Name,
		Description: w.Description,
		UserID:      w.UserID,
		User:        w.User,
		Exercises:   w.Exercises,
		CreatedAt:   w.CreatedAt,
		UpdatedAt:   w.UpdatedAt,
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

func ToResponseWorkouts(users []*Workout) []*ResponseWorkout {
	responseWorkouts := make([]*ResponseWorkout, len(users))
	for i, workout := range users {
		responseWorkouts[i] = workout.ToResponseWorkout()
	}
	return responseWorkouts
}
