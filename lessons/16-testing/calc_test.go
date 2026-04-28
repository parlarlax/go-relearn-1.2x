package main

import (
	"errors"
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive numbers", 2, 3, 5},
		{"negative numbers", -1, -2, -3},
		{"mixed", 5, -3, 2},
		{"zero", 0, 0, 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := Add(tc.a, tc.b)
			if got != tc.expected {
				t.Errorf("Add(%d, %d) = %d, want %d", tc.a, tc.b, got, tc.expected)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	t.Run("normal division", func(t *testing.T) {
		got, err := Divide(10.0, 2.0)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got != 5.0 {
			t.Errorf("Divide(10, 2) = %f, want 5.0", got)
		}
	})

	t.Run("division by zero", func(t *testing.T) {
		_, err := Divide(10.0, 0)
		if err == nil {
			t.Error("expected error for division by zero")
		}
	})

	t.Run("division error type", func(t *testing.T) {
		_, err := Divide(1, 0)
		if !errors.Is(err, ErrDivisionByZero) {
			t.Errorf("unexpected error: %v", err)
		}
	})
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(i, i+1)
	}
}

func BenchmarkDivide(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Divide(float64(i), float64(i+1))
	}
}
