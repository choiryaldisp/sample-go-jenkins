package main

import "testing"

func TestSum(t *testing.T) {
	a := 1
	b := 2
	expected := 3

	actual := Sum(a, b)
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}
