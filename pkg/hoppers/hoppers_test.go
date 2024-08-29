package hoppers

import (
	"errors"
	"testing"
)

func TestHopsCount(t *testing.T) {
	t.Run("Can find path", func(t *testing.T) {
		width, height := 5, 5
		start := Point{4, 0}
		finish := Point{4, 4}
		obstacles := [][2]Point{{Point{1, 2}, Point{4, 3}}}

		count, err := HopsCount(width, height, start, finish, obstacles)
		if count != 7 {
			t.Errorf("Hops should be 7, got: %d", count)
		}
		if err != nil {
			t.Errorf("Error should be nil, got: %s", err.Error())
		}
	})
		
	t.Run("Can find with returns", func(t *testing.T) {
		width, height := 7, 1
		start := Point{3, 0}
		finish := Point{6, 0}
		obstacles := [][2]Point{{Point{4, 0}, Point{5, 0}}}

		count, err := HopsCount(width, height, start, finish, obstacles)
		if count != 7 {
			t.Errorf("Hops should be 7, got: %d", count)
		}
		if err != nil {
			t.Errorf("Error should be nil, got: %s", err.Error())
		}
	})

	t.Run("Path not found", func(t *testing.T) {
		width, height := 4, 1
		start := Point{0, 0}
		finish := Point{3, 0}
		obstacles := [][2]Point{{Point{1, 0}, Point{2, 0}}}

		_, err := HopsCount(width, height, start, finish, obstacles)
		if err == nil { t.Error("Error expected, got nil") }
		if err.Error() != "No solution" {
			t.Errorf("'No solution' error expected, got: %s", err.Error())
		}
	})
}

func TestPresentResult(t *testing.T) {
	tests := []struct {
		name     string
		result     int
		err error
		expected string
	}{
		{"Hops found", 2, nil, "Optimal solution takes 2 hop(s)"},
		{"No solution", -1, errors.New("No solution here"), "No solution here"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := PresentResult(tt.result, tt.err)
			if result != tt.expected {
				t.Errorf("%s. Expected: %s. Got %s", tt.name, result, tt.expected)
			}
		})
	}
}
