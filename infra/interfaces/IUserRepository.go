package interfaces

import "goddd/domain/entity"

type IUserRepository interface {
	Save(user *entity.User) (*entity.User, error)
	Create(user *entity.User) (*entity.User, error)
	FindMany() []*entity.User
	FindById() (*entity.User, error)
	FindOne() (*entity.User, error)
}
