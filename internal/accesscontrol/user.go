package accesscontrol

type UserChecker interface {
	HasAccessToUser(authenticatedUserID, targetUserID int64) bool
}

type UserAccess struct{}

func NewUserAccess() *UserAccess {
	return &UserAccess{}
}

func (ac *UserAccess) HasAccessToUser(authenticatedUserID, targetUserID int64) bool {
	if authenticatedUserID != targetUserID {
		return false
	}
	return true
}
