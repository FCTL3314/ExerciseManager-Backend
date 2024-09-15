package usecase

import (
	"ExerciseManager/bootstrap"
	"ExerciseManager/internal/accesscontrol"
	"ExerciseManager/internal/auth"
	"ExerciseManager/internal/domain"
	"ExerciseManager/internal/tokenutil"
	"errors"
	"gorm.io/gorm"
	"strconv"
)

type UserUsecase struct {
	userRepository    domain.UserRepository
	userAccessChecker accesscontrol.UserChecker
	passwordHasher    auth.PasswordHasher
	tokenManager      *tokenutil.TokenManager
	cfg               *bootstrap.Config
}

func NewUserUsecase(
	userRepository domain.UserRepository,
	userAccessChecker accesscontrol.UserChecker,
	passwordHasher auth.PasswordHasher,
	tokenManager *tokenutil.TokenManager,
	cfg *bootstrap.Config,
) *UserUsecase {
	return &UserUsecase{
		userRepository:    userRepository,
		userAccessChecker: userAccessChecker,
		passwordHasher:    passwordHasher,
		tokenManager:      tokenManager,
		cfg:               cfg,
	}
}

func (uu *UserUsecase) GetById(id uint) (*domain.User, error) {
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

func (uu *UserUsecase) Create(createUser *domain.CreateUserRequest) (*domain.User, error) {
	hashedPassword, err := uu.passwordHasher.Hash(createUser.Password)
	if err != nil {
		return &domain.User{}, err
	}
	createUser.Password = hashedPassword

	if _, err := uu.userRepository.GetByUsername(createUser.Username); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			user := createUser.ToUser()
			return uu.userRepository.Create(user)
		}
		return nil, err
	}

	return &domain.User{}, &domain.ErrObjectUniqueConstraint{Fields: []string{"username"}}
}

func (uu *UserUsecase) Login(loginUser *domain.LoginUserRequest) (*domain.TokensResponse, error) {
	targetUser, err := uu.userRepository.GetByUsername(loginUser.Username)
	if err != nil {
		return nil, domain.ErrInvalidAuthCredentials
	}

	err = uu.passwordHasher.Compare(targetUser.Password, loginUser.Password)
	if err != nil {
		return nil, domain.ErrInvalidAuthCredentials
	}

	accessToken, err := uu.tokenManager.CreateAccessToken(targetUser)
	if err != nil {
		return nil, err
	}
	refreshToken, err := uu.tokenManager.CreateRefreshToken(targetUser)
	if err != nil {
		return nil, err
	}

	return &domain.TokensResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (uu *UserUsecase) RefreshTokens(refreshTokenRequest *domain.RefreshTokenRequest) (*domain.TokensResponse, error) {
	userIDString, err := uu.tokenManager.ExtractUserIDFromRefreshToken(refreshTokenRequest.RefreshToken)
	if err != nil {
		return nil, domain.ErrInvalidAuthCredentials
	}

	userID, err := strconv.ParseUint(userIDString, 10, 64)
	if err != nil {
		return nil, domain.ErrInvalidAuthCredentials
	}

	targetUser, err := uu.userRepository.GetById(uint(userID))
	if err != nil {
		return nil, domain.ErrInvalidAuthCredentials
	}

	accessToken, err := uu.tokenManager.CreateAccessToken(targetUser)
	if err != nil {
		return nil, err
	}
	refreshToken, err := uu.tokenManager.CreateRefreshToken(targetUser)
	if err != nil {
		return nil, err
	}

	return &domain.TokensResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (uu *UserUsecase) Update(authUserId uint, id uint, updateUser *domain.UpdateUserRequest) (*domain.User, error) {
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

	updateUser.ApplyToUser(userToUpdate)

	return uu.userRepository.Update(userToUpdate)
}

func (uu *UserUsecase) Delete(authUserId uint, id uint) error {
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
