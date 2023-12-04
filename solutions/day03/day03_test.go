package day03

import (
	"testing"

	"github.com/tomaskul/advent-of-code-23/util"
)

func Test_isAdjacentToSymbol(t *testing.T) {
	type args struct {
		subject    string
		rowAbove   string
		subjectRow string
		rowBelow   string
	}
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, _ := isAdjacentToSymbol(tt.arguments.subject, tt.arguments.rowAbove, tt.arguments.subjectRow, tt.arguments.rowBelow)
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
		expectedSum int
	}{
		{
			name: "sample",
			input: []string{
				"467..114..",
				"...*......",
				"..35..633.",
				"......#...",
				"617*......",
				".....+.58.",
				"..592.....",
				"......755.",
				"...$.*....",
				".664.598..",
			},
			expectedSum: 4361,
		},
		{
			name: "duplicate number",
			input: []string{
				"..#.........",
				".33.....33..",
				"............",
			},
			expectedSum: 33,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := util.Sum(numbersWithAdjacentSymbols(tt.input))
			if tt.expectedSum != actual {
				t.Errorf("expected: %v, got: %v", tt.expectedSum, actual)
			}
		})
	}
}

// func Test_parseInputData(t *testing.T) {
// 	tests := []struct {
// 		name     string
// 		input    []string
// 		expected any
// 	}{
// 		{
// 			name:     "sample",
// 			input:    []string{},
// 			expected: 0,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			actual := parseInputData(tt.input)
// 			if !reflect.DeepEqual(tt.expected, actual) {
// 				t.Errorf("expected: %v, got: %v", tt.expected, actual)
// 			}
// 		})
// 	}
// }
