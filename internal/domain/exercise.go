package domain

import (
	"time"
)

type Exercise struct {
	ID          uint
	Name        string
	Description string
	Duration    time.Duration
	Workouts    []*Workout `gorm:"many2many:workout_exercises;"`
}

type WorkoutExercises struct {
	WorkoutID  uint
	Workout    *Workout
	ExerciseID uint
	Exercise   *Exercise
	BreakTime  time.Duration
}
