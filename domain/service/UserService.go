package service

import (
	"goddd/domain/entity"
	"goddd/infra/interfaces"
)

type UserService struct {
	userRepository interfaces.IUserRepository
}

func (u *UserService) createUser(user *entity.User) (*entity.User, error) {
	newUser, err := u.userRepository.Save(user)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}

func NewUserService(userRepository interfaces.IUserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}
