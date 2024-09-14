package accesscontrol

type UserChecker interface {
	CanAccessUser(authenticatedUserID, targetUserID uint) bool
}

type User struct {
}

func (ac *User) CanAccessUser(authenticatedUserID, targetUserID uint) bool {
	if authenticatedUserID != targetUserID {
		return false
	}
	return true
}
