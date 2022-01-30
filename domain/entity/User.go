package entity

import "goddd/domain/value"

type User struct {
	FirstName value.FirstName
	LastName  value.LastName
}

func (u *User) ChangeFirstName(firstName string) (*User, error) {
	newUser, err := NewUser(firstName, u.LastName.Value)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func NewUser(firstname string, lastName string) (*User, error) {
	fName, err := value.NewFirstName(firstname)
	if err != nil {
		return nil, err
	}
	lName, err := value.NewLastName(lastName)
	if err != nil {
		return nil, err
	}

	return &User{
		FirstName: *fName,
		LastName:  *lName,
	}, nil
}
