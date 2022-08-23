package entities_test

import (
	"testing"

	"github.com/cleysonph/clean-arch-go/enterprise/entities"
)

func TestNewClientNameShouldReturnErrorWhenIsEmpty(t *testing.T) {
	_, err := entities.NewClientName("")
	if err == nil {
		t.Error("expected error, got nil")
	}
	if err != entities.ErrInvalidClientName {
		t.Error("expected error to be ErrInvalidClientName, got", err)
	}
}

func TestNewClientNameShouldReturnErrorWhenLessThan3Characters(t *testing.T) {
	_, err := entities.NewClientName("aa")
	if err == nil {
		t.Error("expected error, got nil")
	}
	if err != entities.ErrInvalidClientName {
		t.Error("expected error to be ErrInvalidClientName, got", err)
	}
}

func TestNewClientNameShouldReturnErrorWhenHasNonValidCharacters(t *testing.T) {
	_, err := entities.NewClientName("aa@")
	if err == nil {
		t.Error("expected error, got nil")
	}
	if err != entities.ErrInvalidClientName {
		t.Error("expected error to be ErrInvalidClientName, got", err)
	}

	_, err = entities.NewClientName("aa1")
	if err == nil {
		t.Error("expected error, got nil")
	}
	if err != entities.ErrInvalidClientName {
		t.Error("expected error to be ErrInvalidClientName, got", err)
	}
}

func TestNewClientNameShouldReturnClientNameWhenIsValid(t *testing.T) {
	name, err := entities.NewClientName("Test")
	if err != nil {
		t.Error("expected no error, got", err)
	}
	if name.Value != "Test" {
		t.Error("expected name to be Test, got", name.Value)
	}
}

func TestNewClientShouldReturnClientWhenIsValid(t *testing.T) {
	client, err := entities.NewClient("Test", "Test", "test@mail.com", "12345678901")
	if err != nil {
		t.Error("expected no error, got", err)
	}
	if client.FirstName != "Test" {
		t.Error("expected first name to be Test, got", client.FirstName)
	}
	if client.LastName != "Test" {
		t.Error("expected last name to be Test, got", client.LastName)
	}
	if client.Email != "test@mail.com" {
		t.Error("expected email to be test@mail.com, got", client.Email)
	}
	if client.Cpf != "12345678901" {
		t.Error("expected cpf to be 12345678901, got", client.Cpf)
	}
}
