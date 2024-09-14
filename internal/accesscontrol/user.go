package accesscontrol

import "ExerciseManager/internal/domain"

type UserChecker interface {
	CanAccessUser(authenticatedUserID, targetUserID uint) error
}

type User struct {
}

func (ac *User) CanAccessUser(authenticatedUserID, targetUserID uint) error {
	if authenticatedUserID != targetUserID {
		return domain.ErrAccessDenied
	}
	return nil
}
