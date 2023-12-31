package day03

import (
	"reflect"
	"regexp"
	"testing"

	"github.com/tomaskul/advent-of-code-23/util"
)

type args struct {
	subject    string
	rowAbove   string
	subjectRow string
	rowBelow   string
}

func Test_regex(t *testing.T) {
	tests := []struct {
		name            string
		input           string
		expectedMatches []string
	}{
		{
			name:            "actual - weird 206",
			input:           "886.......206..............*6.......595=.....*.85........*..............................286..$...23.....436.................................",
			expectedMatches: []string{"886", "206", "6", "595", "85", "286", "23", "436"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			numMatch, _ := regexp.Compile("\\d+")
			actual := numMatch.FindAllString(tc.input, -1)
			if !reflect.DeepEqual(tc.expectedMatches, actual) {
				t.Errorf("expected: %v, got: %v", tc.expectedMatches, actual)
			}
		})
	}
}

func Test_isAdjacentToSymbol(t *testing.T) {
	tests := []struct {
		name      string
		arguments args
		expected  bool
	}{
		{
			name: "sample - rows[:2] - 467",
			arguments: args{
				subject:    "467",
				rowAbove:   "",
				subjectRow: "467..114..",
				rowBelow:   "...*......",
			},
			expected: true,
		},
		{
			name: "sample - rows[:2] - 114",
			arguments: args{
				subject:    "114",
				rowAbove:   "",
				subjectRow: "467..114..",
				rowBelow:   "...*......",
			},
			expected: false,
		},
		{
			name: "sample - rows[3:5] - 617",
			arguments: args{
				subject:    "617",
				rowAbove:   "......#...",
				subjectRow: "617*......",
				rowBelow:   ".....+.58.",
			},
			expected: true,
		},
		{
			name: "adjacent right - 58",
			arguments: args{
				subject:    "58",
				rowAbove:   "617*......",
				subjectRow: ".....+.58*",
				rowBelow:   "..592.....",
			},
			expected: true,
		},
		{
			name: "sample - rows[4:6] - 58",
			arguments: args{
				subject:    "58",
				rowAbove:   "617*......",
				subjectRow: ".....+.58.",
				rowBelow:   "..592.....",
			},
			expected: false,
		},
		{
			name: "not adj - top right overshoot",
			arguments: args{
				subject:    "617",
				rowAbove:   "....###...",
				subjectRow: "617.......",
				rowBelow:   ".....+.58.",
			},
			expected: false,
		},
		{
			name: "not adj - bottom right overshoot",
			arguments: args{
				subject:    "617",
				rowAbove:   "......#...",
				subjectRow: "617.......",
				rowBelow:   "....++.58.",
			},
			expected: false,
		},
		{
			name: "not adj - top left overshoot",
			arguments: args{
				subject:    "77",
				rowAbove:   "#.........",
				subjectRow: "..77......",
				rowBelow:   ".....+.58.",
			},
			expected: false,
		},
		{
			name: "not adj - bottom left overshoot",
			arguments: args{
				subject:    "77",
				rowAbove:   "......#...",
				subjectRow: "..77......",
				rowBelow:   "#....+.58.",
			},
			expected: false,
		},
	}

	nonFillerRegex, _ := regexp.Compile("[^.0-9]")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, _ := isAdjacentToSymbol(tt.arguments.subject, tt.arguments.rowAbove, tt.arguments.subjectRow, tt.arguments.rowBelow, nonFillerRegex, 1)
			if tt.expected != actual {
				t.Errorf("expected: %v, got: %v", tt.expected, actual)
			}

		})
	}
}

func Test_numbersWithAdjacentSymbols(t *testing.T) {
	tests := []struct {
		name        string
		input       []string
		expectedAdj []int
		expectedSum int
	}{
		// {
		// 	name: "sample",
		// 	input: []string{
		// 		"467..114..",
		// 		"...*......",
		// 		"..35..633.",
		// 		"......#...",
		// 		"617*......",
		// 		".....+.58.",
		// 		"..592.....",
		// 		"......755.",
		// 		"...$.*....",
		// 		".664.598..",
		// 	},
		// 	expectedAdj: []int{467, 35, 633, 617, 592, 755, 664, 598},
		// 	expectedSum: 4361,
		// },
		// {
		// 	name: "duplicate number",
		// 	input: []string{
		// 		"..#.........",
		// 		".33.....33..",
		// 		"............",
		// 	},
		// 	expectedAdj: []int{33},
		// 	expectedSum: 33,
		// },
		{
			name: "single digit",
			input: []string{
				"......",
				"..8...",
				"-*....",
			},
			expectedAdj: []int{8},
			expectedSum: 8,
		},
		// {
		// 	name: "multiple matches - same row",
		// 	input: []string{
		// 		"#...#...#...",
		// 		"5....5.7....",
		// 	},
		// 	expectedAdj: []int{5, 5, 7},
		// 	expectedSum: 17,
		// },
		// {
		// 	name: "multiple matches - same row",
		// 	input: []string{
		// 		"#...#...#...",
		// 		"5....5.7...=",
		// 		"...........2",
		// 	},
		// 	expectedAdj: []int{5, 5, 7, 2},
		// 	expectedSum: 19,
		// },
		// {
		// 	name: "actual - item 1 match - 2nd item subset",
		// 	input: []string{
		// 		"..............",
		// 		"+81.........8.",
		// 		"..............",
		// 	},
		// 	expectedAdj: []int{81},
		// 	expectedSum: 81,
		// },
		// {
		// 	name: "actual - item 2 match - 2nd item subset",
		// 	input: []string{
		// 		"..............",
		// 		".206.......*6.",
		// 		"..............",
		// 	},
		// 	expectedAdj: []int{6},
		// 	expectedSum: 6,
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualAdj := numbersWithAdjacentSymbols(tt.input)
			actualSum := util.Sum(actualAdj)

			if !reflect.DeepEqual(tt.expectedAdj, actualAdj) {
				t.Errorf("expected: %v, got: %v", tt.expectedAdj, actualAdj)
			}
			if tt.expectedSum != actualSum {
				t.Errorf("expected: %v, got: %v", tt.expectedSum, actualSum)
			}
		})
	}
}

func Test_numbersAdjacentToGears(t *testing.T) {
	tests := []struct {
		name      string
		arguments args
		expected  bool
	}{
		{
			name: "adjacent - valid - top & bottom",
			arguments: args{
				//subject:    "617",
				rowAbove:   "...2...",
				subjectRow: "...*...",
				rowBelow:   "...3...",
			},
			expected: true,
		},
		{
			name: "adjacent - valid - sides",
			arguments: args{
				rowAbove:   ".......",
				subjectRow: "..1*5..",
				rowBelow:   ".......",
			},
			expected: true,
		},
		{
			name: "adjacent - valid - diagonal nw-se",
			arguments: args{
				rowAbove:   "100....",
				subjectRow: "...*...",
				rowBelow:   "....7..",
			},
			expected: true,
		},
		{
			name: "adjacent - valid - diagonal ne-sw",
			arguments: args{
				rowAbove:   "....6..",
				subjectRow: "...*...",
				rowBelow:   ".20....",
			},
			expected: true,
		},
		{
			name: "adjacent - invalid - not enough numbers",
			arguments: args{
				//subject:    "617",
				rowAbove:   "100....",
				subjectRow: "...*...",
				rowBelow:   ".......",
			},
			expected: false,
		},
		{
			name: "adjacent - invalid - too many numbers",
			arguments: args{
				//subject:    "617",
				rowAbove:   "100....",
				subjectRow: "...*800",
				rowBelow:   "...6...",
			},
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

		})
	}
}
