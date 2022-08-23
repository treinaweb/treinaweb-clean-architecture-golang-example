package entities

import "regexp"

func NewEmail(value string) (Email, error) {
	if !regexp.MustCompile(`^([a-zA-Z0-9_.+-]+@[a-zA-Z0-9_+-]+\.[a-zA-Z]+)$`).MatchString(value) {
		return Email{}, ErrInvalidEmail
	}
	return Email{value}, nil
}

type Email struct {
	Value string
}
