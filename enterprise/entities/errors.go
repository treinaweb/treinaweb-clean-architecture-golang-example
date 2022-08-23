package entities

import "errors"

var (
	ErrInvalidClientName = errors.New("invalid client name")
	ErrInvalidCpf        = errors.New("invalid cpf")
	ErrInvalidEmail      = errors.New("invalid email")
)
