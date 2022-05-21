package main

import (
	"testing"
)

func TestFib(t *testing.T) {
	num := Fib(6)

	if num != 8 {
		t.Fatalf("expected 8, got %d", num)
	}
}
