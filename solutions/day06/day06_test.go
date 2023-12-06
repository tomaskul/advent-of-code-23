package day06

import (
	"testing"
)

func TestRace_noOptionsToBeat(t *testing.T) {
	tests := []struct {
		name     string
		input    race
		expected int
	}{
		{
			name: "sample 1",
			input: race{
				timeMs:     7,
				distanceMm: 9,
			},
			expected: 4,
		},
		{
			name: "sample 2",
			input: race{
				timeMs:     15,
				distanceMm: 40,
			},
			expected: 8,
		},
		{
			name: "sample 3",
			input: race{
				timeMs:     30,
				distanceMm: 200,
			},
			expected: 9,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.input.noOptionsToBeat()
			if tc.expected != actual {
				t.Errorf("expected: %d, got: %d", tc.expected, actual)
			}
		})
	}
}
