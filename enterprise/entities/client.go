package entities

import (
	"regexp"
)

func NewClientName(value string) (ClientName, error) {
	if len(value) < 3 {
		return ClientName{}, ErrInvalidClientName
	} else if !regexp.MustCompile(`^([a-zA-Z]+)$`).MatchString(value) {
		return ClientName{}, ErrInvalidClientName
	}
	return ClientName{value}, nil
}

type ClientName struct {
	Value string
}

func NewClient(firstName string, lastName string, email string, cpf string) (Client, error) {
	firstNameValue, err := NewClientName(firstName)
	if err != nil {
		return Client{}, err
	}
	lastNameValue, err := NewClientName(lastName)
	if err != nil {
		return Client{}, err
	}
	emailValue, err := NewEmail(email)
	if err != nil {
		return Client{}, err
	}
	cpfValue, err := NewCpf(cpf)
	if err != nil {
		return Client{}, err
	}
	return Client{firstNameValue.Value, lastNameValue.Value, emailValue.Value, cpfValue.Value}, nil
}

type Client struct {
	FirstName string
	LastName  string
	Email     string
	Cpf       string
}
