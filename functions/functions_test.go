package main

import "testing"

func TestPlus(t *testing.T) {
	result := Plus2(1, 2)
	expected := 3
	if result != expected {
		t.Errorf("Plus(2,3) = %d ; want %d", result, expected)
	}
}

func TestPlus3(t *testing.T) {
	result := Plus3(2, 3, 5)
	expected := 10
	if result != expected {
		t.Errorf("Plus(2,3,5) = %d ; Correct Answer is %d", result, expected)
	}
}
