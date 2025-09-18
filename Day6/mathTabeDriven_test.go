// Go convention: write one test function with multiple test cases in a slice.

package mathutils

import "testing"

func TestAdd_TableDriven(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"2+2", 2, 2, 6},
		{"-1+1", -1, 1, 0},
		{"54+34", 54, 34, 88},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Expected %d got %d", tt.expected, result)
			}
		})
	}
}
