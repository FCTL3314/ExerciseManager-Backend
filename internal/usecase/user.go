package usecase

import (
	"ExerciseManager/internal/domain"
	"errors"
	"gorm.io/gorm"
)

type UserUsecase struct {
	userRepository domain.UserRepository
}

func NewUserUsecase(userRepository domain.UserRepository) *UserUsecase {
	return &UserUsecase{userRepository: userRepository}
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

func (uu *UserUsecase) Create(createUser *domain.CreateUser) (*domain.User, error) {
	if _, err := uu.userRepository.GetByUsername(createUser.Username); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return uu.userRepository.Create(createUser.ToUser())
		}
		return nil, err
	}

	return &domain.User{}, &domain.ErrObjectUniqueConstraint{Fields: []string{"username"}}
}

func (uu *UserUsecase) Update(id uint, updateUser *domain.UpdateUser) (*domain.User, error) {
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

func (uu *UserUsecase) Delete(id uint) error {
	if _, err := uu.userRepository.GetById(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.ErrObjectNotFound
		}
		return err
	}

	return uu.userRepository.Delete(id)
}
