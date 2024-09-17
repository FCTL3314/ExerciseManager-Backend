package usecase

import (
	"ExerciseManager/bootstrap"
	"ExerciseManager/internal/auth"
	"ExerciseManager/internal/domain"
	"ExerciseManager/internal/errormapper"
	"ExerciseManager/internal/permission"
	"ExerciseManager/internal/tokenutil"
)

type UserUsecase struct {
	userRepository domain.UserRepository
	accessManager  permission.AccessPolicy
	passwordHasher auth.PasswordHasher
	tokenManager   tokenutil.IJWTTokenManager
	errorMapper    errormapper.Chain
	cfg            *bootstrap.Config
}

func NewUserUsecase(
	userRepository domain.UserRepository,
	accessManager permission.AccessPolicy,
	passwordHasher auth.PasswordHasher,
	tokenManager tokenutil.IJWTTokenManager,
	errorMapper errormapper.Chain,
) *UserUsecase {
	return &UserUsecase{
		userRepository: userRepository,
		accessManager:  accessManager,
		passwordHasher: passwordHasher,
		tokenManager:   tokenManager,
		errorMapper:    errorMapper,
	}
}

func (uu *UserUsecase) GetById(id int64) (*domain.User, error) {
	user, err := uu.userRepository.GetById(id)
	if err != nil {
		return nil, uu.errorMapper.MapError(err)
	}

	return user, nil
}

func (uu *UserUsecase) Get(params *domain.FilterParams) (*domain.User, error) {
	user, err := uu.userRepository.Get(params)
	if err != nil {
		return nil, uu.errorMapper.MapError(err)
	}
	return user, nil
}

func (uu *UserUsecase) List(params *domain.Params) (*domain.PaginatedResult[*domain.User], error) {
	users, err := uu.userRepository.Fetch(params)
	if err != nil {
		return nil, uu.errorMapper.MapError(err)
	}

	count, err := uu.userRepository.Count(&domain.FilterParams{})
	if err != nil {
		return nil, uu.errorMapper.MapError(err)
	}

	return &domain.PaginatedResult[*domain.User]{Results: users, Count: count}, nil
}

func (uu *UserUsecase) Create(createUserRequest *domain.CreateUserRequest) (*domain.User, error) {
	hashedPassword, err := uu.passwordHasher.Hash(createUserRequest.Password)
	if err != nil {
		return nil, uu.errorMapper.MapError(err)
	}

	user := domain.NewUserFromCreateRequest(createUserRequest)
	user.Password = hashedPassword

	createdUser, err := uu.userRepository.Create(user)
	if err != nil {
		return nil, uu.errorMapper.MapError(err)
	}
	return createdUser, nil
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
		return nil, uu.errorMapper.MapError(err)
	}
	refreshToken, err := uu.tokenManager.CreateUserRefreshToken(targetUser)
	if err != nil {
		return nil, uu.errorMapper.MapError(err)
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
		return nil, uu.errorMapper.MapError(err)
	}
	refreshToken, err := uu.tokenManager.CreateUserRefreshToken(targetUser)
	if err != nil {
		return nil, uu.errorMapper.MapError(err)
	}

	return &domain.TokensResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (uu *UserUsecase) Update(authUserId int64, id int64, updateUserRequest *domain.UpdateUserRequest) (*domain.User, error) {
	userToUpdate, err := uu.userRepository.GetById(id)
	if err != nil {
		return nil, uu.errorMapper.MapError(err)
	}

	if !uu.accessManager.HasAccess(authUserId, userToUpdate) {
		return nil, domain.ErrAccessDenied
	}

	userToUpdate.ApplyUpdate(updateUserRequest)
	updatedUser, err := uu.userRepository.Update(userToUpdate)
	if err != nil {
		return nil, uu.errorMapper.MapError(err)
	}
	return updatedUser, nil
}

func (uu *UserUsecase) Delete(authUserId int64, id int64) error {
	user, err := uu.userRepository.GetById(id)
	if err != nil {
		return uu.errorMapper.MapError(err)
	}

	if !uu.accessManager.HasAccess(authUserId, user) {
		return domain.ErrAccessDenied
	}

	return uu.userRepository.Delete(id)
}
