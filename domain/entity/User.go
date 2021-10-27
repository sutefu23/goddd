package entity

import "goddd/domain/value"

type User struct {
	firstName value.FirstName
	lastName  value.LastName
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
		firstName: *fName,
		lastName:  *lName,
	}, nil
}
