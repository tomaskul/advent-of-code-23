package day07

import (
	"reflect"
	"testing"
)

func Test_cashInJokers(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected map[byte]int
	}{
		{
			name:  "sample T55J5",
			input: "T55J5",
			expected: map[byte]int{
				'T': 1,
				'5': 4,
			},
		},
		{
			name:  "sample KTJJT",
			input: "KTJJT",
			expected: map[byte]int{
				'T': 4,
				'K': 1,
			},
		},
		{
			name:  "sample QQQJA",
			input: "QQQJA",
			expected: map[byte]int{
				'Q': 4,
				'A': 1,
			},
		},
		{
			name:  "actul data J4443",
			input: "J4443",
			expected: map[byte]int{
				'4': 4,
				'3': 1,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			sut := newCard(tc.input)
			actual := sut.cashInJokers()

			if !reflect.DeepEqual(tc.expected, actual) {
				t.Errorf("expected: %v, got: %v", tc.expected, actual)
			}
		})
	}
}

func TestCard_getType(t *testing.T) {
	tests := []struct {
		name         string
		input        string
		expectedType int
	}{
		{
			name:         "actual - cash in jokers - JQQQT",
			input:        "JQQQT",
			expectedType: FourOfKindType,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			sut := newCard(tc.input)
			actual := sut.getType(true)
			if tc.expectedType != actual {
				t.Errorf("expected: %d, got: %d", tc.expectedType, actual)
			}
		})
	}
}
