package entities_test

import (
	"testing"

	"github.com/cleysonph/clean-arch-go/enterprise/entities"
)

func TestNewCpfShouldReturnErrorWhenLessThan11Characters(t *testing.T) {
	_, err := entities.NewCpf("123")
	if err == nil {
		t.Error("expected error, got nil")
	}
	if err != entities.ErrInvalidCpf {
		t.Error("expected error to be ErrInvalidCpf, got", err)
	}
}

func TestNewCpfShouldReturnErrorWhenGreaterThan11Characters(t *testing.T) {
	_, err := entities.NewCpf("123456789011")
	if err == nil {
		t.Error("expected error, got nil")
	}
	if err != entities.ErrInvalidCpf {
		t.Error("expected error to be ErrInvalidCpf, got", err)
	}
}

func TestNewCpfShouldReturnErrorWhenHasInvalidCharacters(t *testing.T) {
	_, err := entities.NewCpf("1234567890a")
	if err == nil {
		t.Error("expected error, got nil")
	}
	if err != entities.ErrInvalidCpf {
		t.Error("expected error to be ErrInvalidCpf, got", err)
	}
}

func TestNewCpfShouldReturnCpfWhenValid(t *testing.T) {
	cpf, err := entities.NewCpf("12345678901")
	if err != nil {
		t.Error("expected no error, got", err)
	}
	if cpf.Value != "12345678901" {
		t.Error("expected cpf to be 12345678901, got", cpf.Value)
	}
}
