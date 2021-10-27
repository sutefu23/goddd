package value

import (
	"errors"
	"unicode/utf8"
)

type LastName struct {
	value string
}

func NewLastName(value string) (*LastName, error) {
	if utf8.RuneCountInString(value) > 10 {
		return nil, errors.New("lastName must be under 10")

	}
	return &LastName{value: value}, nil
}
