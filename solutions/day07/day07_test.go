package day07

import (
	"reflect"
	"testing"
)

func Test_sortPt1(t *testing.T) {
	tests := []struct {
		name         string
		input        []string
		expectedBids []int
	}{
		{
			name: "sample",
			input: []string{"32T3K 765",
				"T55J5 684",
				"KK677 28",
				"KTJJT 220",
				"QQQJA 483"},
			expectedBids: []int{765, 220, 28, 684, 483},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			items := parse(tc.input)
			sorted := sortPt1(items)

			actual := make([]int, 5)
			for i := 0; i < len(sorted); i++ {
				actual[i] = sorted[i].bid
			}

			if !reflect.DeepEqual(tc.expectedBids, actual) {
				t.Errorf("expected: %v, got: %v", tc.expectedBids, actual)
			}
		})
	}
}

func Test_sortPt2(t *testing.T) {
	tests := []struct {
		name         string
		input        []string
		expectedBids []int
	}{
		{
			name: "sample",
			input: []string{"32T3K 765",
				"T55J5 684",
				"KK677 28",
				"KTJJT 220",
				"QQQJA 483"},
			expectedBids: []int{765, 28, 684, 483, 220},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			items := parse(tc.input)
			sorted := sortPt2(items)

			actual := make([]int, 5)
			for i := 0; i < len(sorted); i++ {
				actual[i] = sorted[i].bid
			}

			if !reflect.DeepEqual(tc.expectedBids, actual) {
				t.Errorf("expected: %v, got: %v", tc.expectedBids, actual)
			}
		})
	}
}

func Test_calculateWinnings(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected int
	}{
		{
			name: "sample",
			input: []string{"32T3K 765",
				"T55J5 684",
				"KK677 28",
				"KTJJT 220",
				"QQQJA 483"},
			expected: 6440,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			items := parse(tc.input)
			sortedItems := sortPt1(items)
			actual := calculateWinnings(sortedItems)

			if tc.expected != actual {
				t.Errorf("expected: %d, got: %d", tc.expected, actual)
			}
		})
	}
}
