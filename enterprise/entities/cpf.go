package entities

import "regexp"

func NewCpf(value string) (Cpf, error) {
	if !regexp.MustCompile(`^([0-9]{11})$`).MatchString(value) {
		return Cpf{}, ErrInvalidCpf
	}
	return Cpf{value}, nil
}

type Cpf struct {
	Value string
}
