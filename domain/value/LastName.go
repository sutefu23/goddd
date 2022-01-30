package value

import (
	"errors"
	"unicode/utf8"
)

type LastName struct {
	Value string
}

func NewLastName(value string) (*LastName, error) {
	if utf8.RuneCountInString(value) > 10 {
		return nil, errors.New("lastName must be under 10")

	}
	return &LastName{Value: value}, nil
}
