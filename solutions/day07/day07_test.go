package day07

import "testing"

// func TestCardData_handScore(t *testing.T) {
// 	tests := []struct {
// 		name     string
// 		input    card
// 		expected int
// 	}{
// 		{
// 			name:     "sample 1",
// 			input:    fromString("AAAAA"),
// 			expected: 13 * 5,
// 		},
// 		{
// 			name:     "sample 2",
// 			input:    fromString("22222"),
// 			expected: 1 * 5,
// 		},
// 	}

// 	for _, tc := range tests {
// 		t.Run(tc.name, func(t *testing.T) {
// 			actual := tc.input.handScore()
// 			if tc.expected != actual {
// 				t.Errorf("expected: %d, got: %d", tc.expected, actual)
// 			}
// 		})
// 	}
// }

// func TestCardData_SortedCorectly(t *testing.T) {
// 	tests := []struct {
// 		name     string
// 		input    []string
// 		expected int
// 	}{
// 		{
// 			name: "sample",
// 			input: []string{"32T3K 765",
// 				"T55J5 684",
// 				"KK677 28",
// 				"KTJJT 220",
// 				"QQQJA 483"},
// 			expected: 13 * 5,
// 		},
// 	}

// 	for _, tc := range tests {
// 		t.Run(tc.name, func(t *testing.T) {
// 			items := parsePt1(tc.input)
// 			actual := sortPt1(items)
// 			t.Errorf("%+v", actual)
// 			// actual := tc.input.handScore()
// 			// if tc.expected != actual {
// 			// 	t.Errorf("expected: %d, got: %d", tc.expected, actual)
// 			// }
// 		})
// 	}
// }

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
			items := parsePt1(tc.input)
			sortedItems := sortPt1(items)
			actual := calculateWinnings(sortedItems)

			if tc.expected != actual {
				t.Errorf("expected: %d, got: %d", tc.expected, actual)
			}
		})
	}
}
