package day09

import "testing"

func Test_nextInSequence(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "sample 15..18",
			input:    []int{0, 3, 6, 9, 12, 15},
			expected: 18,
		},
		{
			name:     "sample 21..28",
			input:    []int{1, 3, 6, 10, 15, 21},
			expected: 28,
		},
		{
			name:     "sample 45..68",
			input:    []int{10, 13, 16, 21, 30, 45},
			expected: 68,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := nextInSequence(tc.input)
			if actual != tc.expected {
				t.Errorf("expected: %d, got: %d", tc.expected, actual)
			}
		})
	}
}
