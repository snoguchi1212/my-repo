package main

import "testing"

func TestEvenOrOdd(t *testing.T) {
	result := EvenOrOdd(8)
	if result != "Even" {
		t.Errorf("expected: even, actual: %s", result)
	}
}
