package usecase

import "ExerciseManager/internal/domain"

type UserUsecase struct {
	userRepository domain.UserRepository
}

func NewUserUsecase(userRepository domain.UserRepository) *UserUsecase {
	return &UserUsecase{userRepository: userRepository}
}

func (uu *UserUsecase) Get(params *domain.FilterParams) (*domain.User, error) {
	return uu.userRepository.Get(params)
}

func (uu *UserUsecase) List(params *domain.Params) ([]*domain.User, error) {
	return uu.userRepository.List(params)
}

func (uu *UserUsecase) Create(user *domain.User) (*domain.User, error) {
	return uu.userRepository.Create(user)
}

func (uu *UserUsecase) Delete(id uint) error {
	return uu.userRepository.Delete(id)
}
