package models_test

import (
	"cronmail/models"
	"testing"
)

func TestPersonAge(t *testing.T) {
	p := models.Person{
		Birthdate: "20/02/1998",
	}
	age, err := p.Age()

	if err != nil {
		t.Error("error is nil")
	}

	if age < 27 {
		t.Error("wrong age")
	}
}

func TestInvalidAge(t *testing.T) {
	p := models.Person{
		Birthdate: "Ciao",
	}

	age, err := p.Age()

	if err == nil {
		t.Error("age shall be invalid")
	}

	if age != 0 {
		t.Error("age shall be zero-value")
	}
}
