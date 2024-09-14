package accesscontrol

type UserChecker interface {
	CanAccessUser(authenticatedUserID, targetUserID uint) bool
}

type UserAccess struct{}

func NewUserAccess() *UserAccess {
	return &UserAccess{}
}

func (ac *UserAccess) CanAccessUser(authenticatedUserID, targetUserID uint) bool {
	if authenticatedUserID != targetUserID {
		return false
	}
	return true
}
