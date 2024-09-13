package usecase

import (
	"ExerciseManager/internal/domain"
)

type UserUsecase struct {
	userRepository domain.UserRepository
}

func NewUserUsecase(userRepository domain.UserRepository) *UserUsecase {
	return &UserUsecase{userRepository: userRepository}
}

func (uu *UserUsecase) Get(params *domain.FilterParams) (*domain.User, error) {
	return uu.userRepository.Get(params)
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

func (uu *UserUsecase) Create(user *domain.User) (*domain.User, error) {
	return uu.userRepository.Create(user)
}

func (uu *UserUsecase) Update(user *domain.User) (*domain.User, error) {
	return uu.userRepository.Update(user)
}

func (uu *UserUsecase) Delete(id uint) error {
	return uu.userRepository.Delete(id)
}
