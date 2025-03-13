package main

import (
	"fmt"
	"testing"
)

func TestStringInterpolation(t *testing.T) {
	t.Logf("Running TestStringInterpolation")
	name := "John"
	age := 30
	expected := "My name is John and I am 30 years old."
	result := fmt.Sprintf("My name is %s and I am %d years old.", name, age)

	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}
