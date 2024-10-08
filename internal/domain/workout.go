package domain

import (
	"time"
)

type Workout struct {
	ID               int64
	Name             string
	Description      string
	UserID           int64
	User             *User
	WorkoutExercises []*WorkoutExercise `gorm:"foreignKey:WorkoutID"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type WorkoutExercise struct {
	ID         int64
	WorkoutID  int64
	Workout    *Workout
	ExerciseID int64
	Exercise   *Exercise
	BreakTime  time.Duration
	CreatedAt  time.Time
}

type ResponseWorkout struct {
	ID               int64                      `json:"id"`
	Name             string                     `json:"name"`
	Description      string                     `json:"description"`
	User             *ResponseUser              `json:"user"`
	WorkoutExercises []*ResponseWorkoutExercise `json:"workout_exercises"`
	CreatedAt        time.Time                  `json:"created_at"`
	UpdatedAt        time.Time                  `json:"updated_at"`
}

type CreateWorkoutRequest struct {
	Name        string `json:"name" binding:"required,min=1,max=128"`
	Description string `json:"description" binding:"omitempty,max=516"`
}

type UpdateWorkoutRequest struct {
	Name        *string `json:"name,omitempty" binding:"omitempty,min=1,max=128"`
	Description *string `json:"description,omitempty" binding:"omitempty,max=516"`
}

type ResponseWorkoutExercise struct {
	ID        int64                   `json:"id"`
	Exercise  *ResponseNestedExercise `json:"exercise"`
	BreakTime time.Duration           `json:"break_time"`
	CreatedAt time.Time               `json:"created_at"`
}

type AddExerciseToWorkoutRequest struct {
	ExerciseID int64         `json:"exercise_id" binding:"required"`
	BreakTime  time.Duration `json:"break_time" binding:"required"`
}

type UpdateWorkoutExerciseRequest struct {
	BreakTime *time.Duration `json:"break_time,omitempty" binding:"omitempty"`
}

func NewWorkoutFromCreateRequest(req *CreateWorkoutRequest) *Workout {
	return &Workout{
		Name:        req.Name,
		Description: req.Description,
	}
}

func (w *Workout) ToResponseWorkout() *ResponseWorkout {
	rw := &ResponseWorkout{
		ID:               w.ID,
		Name:             w.Name,
		Description:      w.Description,
		WorkoutExercises: ToResponseWorkoutExercises(w.WorkoutExercises),
		CreatedAt:        w.CreatedAt,
		UpdatedAt:        w.UpdatedAt,
	}

	if w.User != nil {
		rw.User = w.User.ToResponseUser()
	}

	return rw
}

func (w *Workout) ApplyUpdate(req *UpdateWorkoutRequest) {
	if req.Name != nil {
		w.Name = *req.Name
	}

	if req.Description != nil {
		w.Description = *req.Description
	}
}

func (we *WorkoutExercise) ToResponseWorkoutExercise() *ResponseWorkoutExercise {
	rw := &ResponseWorkoutExercise{
		ID:        we.ID,
		BreakTime: we.BreakTime,
		CreatedAt: we.CreatedAt,
	}

	if we.Exercise != nil {
		rw.Exercise = we.Exercise.ToResponseNestedExercise()
	}

	return rw
}

func (we *WorkoutExercise) ApplyUpdate(req *UpdateWorkoutExerciseRequest) {
	if req.BreakTime != nil {
		we.BreakTime = *req.BreakTime
	}
}

func ToResponseWorkouts(workouts []*Workout) []*ResponseWorkout {
	responseWorkouts := make([]*ResponseWorkout, len(workouts))
	for i, workout := range workouts {
		responseWorkouts[i] = workout.ToResponseWorkout()
	}
	return responseWorkouts
}

func ToResponseWorkoutExercises(workoutExercises []*WorkoutExercise) []*ResponseWorkoutExercise {
	responseWorkoutExercise := make([]*ResponseWorkoutExercise, len(workoutExercises))
	for i, workoutExercise := range workoutExercises {
		responseWorkoutExercise[i] = workoutExercise.ToResponseWorkoutExercise()
	}
	return responseWorkoutExercise
}
