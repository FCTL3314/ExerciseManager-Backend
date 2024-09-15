package accesscontrol

type UserChecker interface {
	CanAccessUser(authenticatedUserID, targetUserID int64) bool
}

type UserAccess struct{}

func NewUserAccess() *UserAccess {
	return &UserAccess{}
}

func (ac *UserAccess) CanAccessUser(authenticatedUserID, targetUserID int64) bool {
	if authenticatedUserID != targetUserID {
		return false
	}
	return true
}
