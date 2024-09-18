package permission

import (
	"ExerciseManager/internal/domain"
	"reflect"
)

type AccessController interface {
	CanAccessResource(authenticatedUserID int64, resource interface{}) bool
}

type AccessPolicy interface {
	HasAccess(authenticatedUserID int64, resource interface{}) bool
}

type UserAccessPolicy struct{}

func (ua *UserAccessPolicy) HasAccess(authenticatedUserID int64, resource interface{}) bool {
	targetUser, ok := resource.(*domain.User)
	if !ok {
		return false
	}
	return authenticatedUserID == targetUser.ID
}

type WorkoutAccessPolicy struct{}

func (wa *WorkoutAccessPolicy) HasAccess(authenticatedUserID int64, resource interface{}) bool {
	targetWorkout, ok := resource.(*domain.Workout)
	if !ok {
		return false
	}
	return authenticatedUserID == targetWorkout.UserID
}

type ExerciseAccessPolicy struct{}

func (wa *ExerciseAccessPolicy) HasAccess(authenticatedUserID int64, resource interface{}) bool {
	targetExercise, ok := resource.(*domain.Exercise)
	if !ok {
		return false
	}
	return authenticatedUserID == targetExercise.UserID
}

type AccessManager struct {
	policies map[reflect.Type]AccessPolicy
}

func NewAccessManager() *AccessManager {
	return &AccessManager{
		policies: make(map[reflect.Type]AccessPolicy),
	}
}

func (am *AccessManager) RegisterPolicy(resourceType reflect.Type, policy AccessPolicy) {
	am.policies[resourceType] = policy
}

func (am *AccessManager) HasAccess(authenticatedUserID int64, resource interface{}) bool {
	resourceType := reflect.TypeOf(resource)
	policy, exists := am.policies[resourceType]
	if !exists {
		return false
	}
	return policy.HasAccess(authenticatedUserID, resource)
}

func BuildUserAccessManager() *AccessManager {
	accessManager := NewAccessManager()
	accessManager.RegisterPolicy(reflect.TypeOf(&domain.User{}), &UserAccessPolicy{})
	return accessManager
}

func BuildWorkoutAccessManager() *AccessManager {
	accessManager := NewAccessManager()
	accessManager.RegisterPolicy(reflect.TypeOf(&domain.Workout{}), &WorkoutAccessPolicy{})
	return accessManager
}

func BuildExerciseAccessManager() *AccessManager {
	accessManager := NewAccessManager()
	accessManager.RegisterPolicy(reflect.TypeOf(&domain.Exercise{}), &ExerciseAccessPolicy{})
	return accessManager
}

func BuildAllEntitiesAccessManager() *AccessManager {
	accessManager := NewAccessManager()
	accessManager.RegisterPolicy(reflect.TypeOf(&domain.User{}), &UserAccessPolicy{})
	accessManager.RegisterPolicy(reflect.TypeOf(&domain.Workout{}), &WorkoutAccessPolicy{})
	accessManager.RegisterPolicy(reflect.TypeOf(&domain.Exercise{}), &ExerciseAccessPolicy{})
	return accessManager
}
