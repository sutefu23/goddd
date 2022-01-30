package collection

import "goddd/domain/entity"

type Users struct {
	Users []*entity.User
}

func NewUsers() *Users {
	return &Users{
		Users: []*entity.User{},
	}
}

// コレクションにuserを1つ追加
func (u *Users) Add(user *entity.User) {
	u.Users = append(u.Users, user)
}
