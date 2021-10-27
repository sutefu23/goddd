package usecase

import "goddd/domain/entity"

type IUserRepository interface {
	save() (*entity.User, error)
	create() (*entity.User, error)
	findMany()
	findOne() (*entity.User, error)
}

type UserService struct {
	userRepository IUserRepository
}

func NewUserService(userRepository IUserRepository) {
	userservice := UserService{userRepository: userRepository}
}
