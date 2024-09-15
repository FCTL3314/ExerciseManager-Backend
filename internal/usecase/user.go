package usecase

import (
	"ExerciseManager/bootstrap"
	"ExerciseManager/internal/accesscontrol"
	"ExerciseManager/internal/auth"
	"ExerciseManager/internal/domain"
	"ExerciseManager/internal/tokenutil"
	"errors"
	"gorm.io/gorm"
)

type UserUsecase struct {
	userRepository    domain.UserRepository
	userAccessChecker accesscontrol.UserChecker
	passwordHasher    auth.PasswordHasher
	tokenManager      tokenutil.IJWTTokenManager
	cfg               *bootstrap.Config
}

func NewUserUsecase(
	userRepository domain.UserRepository,
	userAccessChecker accesscontrol.UserChecker,
	passwordHasher auth.PasswordHasher,
	tokenManager tokenutil.IJWTTokenManager,
) *UserUsecase {
	return &UserUsecase{
		userRepository:    userRepository,
		userAccessChecker: userAccessChecker,
		passwordHasher:    passwordHasher,
		tokenManager:      tokenManager,
	}
}

func (uu *UserUsecase) GetById(id int64) (*domain.User, error) {
	user, err := uu.userRepository.GetById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrObjectNotFound
		}
		return nil, err
	}

	return user, nil
}

func (uu *UserUsecase) Get(params *domain.FilterParams) (*domain.User, error) {
	user, err := uu.userRepository.Get(params)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrObjectNotFound
		}
		return nil, err
	}
	return user, nil
}

func (uu *UserUsecase) List(params *domain.Params) (*domain.PaginatedResult[*domain.User], error) {
	users, err := uu.userRepository.Fetch(params)
	if err != nil {
		return &domain.PaginatedResult[*domain.User]{}, err
	}

	count, err := uu.userRepository.Count(&domain.FilterParams{})
	if err != nil {
		return &domain.PaginatedResult[*domain.User]{}, err
	}

	return &domain.PaginatedResult[*domain.User]{Results: users, Count: count}, nil
}

func (uu *UserUsecase) Create(createUserRequest *domain.CreateUserRequest) (*domain.User, error) {
	hashedPassword, err := uu.passwordHasher.Hash(createUserRequest.Password)
	if err != nil {
		return nil, err
	}

	user := domain.NewUserFromCreateRequest(createUserRequest)
	user.Password = hashedPassword

	if _, err := uu.userRepository.GetByUsername(user.Username); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return uu.userRepository.Create(user)
		}
		return nil, err
	}

	return nil, &domain.ErrObjectUniqueConstraint{Fields: []string{"username"}}
}

func (uu *UserUsecase) Login(loginUserRequest *domain.LoginUserRequest) (*domain.TokensResponse, error) {
	targetUser, err := uu.userRepository.GetByUsername(loginUserRequest.Username)
	if err != nil {
		return nil, domain.ErrInvalidAuthCredentials
	}

	err = uu.passwordHasher.Compare(targetUser.Password, loginUserRequest.Password)
	if err != nil {
		return nil, domain.ErrInvalidAuthCredentials
	}

	accessToken, err := uu.tokenManager.CreateUserAccessToken(targetUser)
	if err != nil {
		return nil, err
	}
	refreshToken, err := uu.tokenManager.CreateUserRefreshToken(targetUser)
	if err != nil {
		return nil, err
	}

	return &domain.TokensResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (uu *UserUsecase) RefreshTokens(refreshTokenRequest *domain.RefreshTokenRequest) (*domain.TokensResponse, error) {
	userId, err := uu.tokenManager.ExtractUserIDFromRefreshToken(refreshTokenRequest.RefreshToken)
	if err != nil {
		return nil, domain.ErrInvalidAuthCredentials
	}

	targetUser, err := uu.userRepository.GetById(userId)
	if err != nil {
		return nil, domain.ErrInvalidAuthCredentials
	}

	accessToken, err := uu.tokenManager.CreateUserAccessToken(targetUser)
	if err != nil {
		return nil, err
	}
	refreshToken, err := uu.tokenManager.CreateUserRefreshToken(targetUser)
	if err != nil {
		return nil, err
	}

	return &domain.TokensResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (uu *UserUsecase) Update(authUserId int64, id int64, updateUserRequest *domain.UpdateUserRequest) (*domain.User, error) {
	if !uu.userAccessChecker.CanAccessUser(authUserId, id) {
		return nil, domain.ErrAccessDenied
	}

	userToUpdate, err := uu.userRepository.GetById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrObjectNotFound
		}
		return nil, err
	}

	userToUpdate.ApplyUpdate(updateUserRequest)

	return uu.userRepository.Update(userToUpdate)
}

func (uu *UserUsecase) Delete(authUserId int64, id int64) error {
	if !uu.userAccessChecker.CanAccessUser(authUserId, id) {
		return domain.ErrAccessDenied
	}

	if _, err := uu.userRepository.GetById(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.ErrObjectNotFound
		}
		return err
	}

	return uu.userRepository.Delete(id)
}
