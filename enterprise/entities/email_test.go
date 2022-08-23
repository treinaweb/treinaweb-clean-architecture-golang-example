package entities_test

import (
	"testing"

	"github.com/cleysonph/clean-arch-go/enterprise/entities"
)

func TestNewEmailShouldReturnErrorWhenNotHasAt(t *testing.T) {
	_, err := entities.NewEmail("testmail.com")
	if err == nil {
		t.Error("expected error, got nil")
	}
	if err != entities.ErrInvalidEmail {
		t.Error("expected error to be ErrInvalidEmail, got", err)
	}
}

func TestNewEmailShouldReturnErrorWhenNotHasDot(t *testing.T) {
	_, err := entities.NewEmail("test@mailcom")
	if err == nil {
		t.Error("expected error, got nil")
	}
	if err != entities.ErrInvalidEmail {
		t.Error("expected error to be ErrInvalidEmail, got", err)
	}
}

func TestNewEmailShouldReturnErrorWhenNotHasValidDomain(t *testing.T) {
	_, err := entities.NewEmail("test@mailcom.")
	if err == nil {
		t.Error("expected error, got nil")
	}
	if err != entities.ErrInvalidEmail {
		t.Error("expected error to be ErrInvalidEmail, got", err)
	}
}

func TestNewEmailShouldReturnEmailWhenIsValid(t *testing.T) {
	email, err := entities.NewEmail("test@mail.com")
	if err != nil {
		t.Error("expected no error, got", err)
	}
	if email.Value != "test@mail.com" {
		t.Error("expected email to be 12345678901, got", email.Value)
	}
}
