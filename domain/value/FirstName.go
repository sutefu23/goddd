package value

import (
	"errors"
	"unicode/utf8"
)

type FirstName struct {
	value string
}

func NewFirstName(value string) (*FirstName, error) {
	if utf8.RuneCountInString(value) > 10 {
		return nil, errors.New("firstName must be under 10")
	}
	return &FirstName{value: value}, nil
}
