package day03

import (
	"reflect"
	"testing"
)

func Test_parseInputData(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected any
	}{
		{
			name:     "sample",
			input:    []string{},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := parseInputData(tt.input)
			if !reflect.DeepEqual(tt.expected, actual) {
				t.Errorf("expected: %v, got: %v", tt.expected, actual)
			}
		})
	}
}
