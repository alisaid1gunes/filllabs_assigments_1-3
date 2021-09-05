package main

import (
	"testing"
)

func TestFindMostRepeated(t *testing.T) {
	x := []string{"apple", "pie", "apple", "red", "red", "red"}
	result := findMostRepeated(x)
	expected := "red"
	if result != expected {
		t.Errorf("add() test returned an unexpected result: got %v want %v", result, expected)
	}
}
