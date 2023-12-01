package day01

import (
	"reflect"
	"testing"
)

func Test_getAllCalibrationValues(t *testing.T) {
	input := []string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	}
	expected := []int{12, 38, 15, 77}

	sut := &Day01{}
	actual := sut.getAllCalibrationValues(input)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected: %v, got: %v", expected, actual)
	}
}

func Test_getRealCalibrationValues(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []int
	}{
		{
			name:     "sample",
			input:    []string{"two1nine", "eightwothree", "abcone2threexyz", "xtwone3four", "4nineeightseven2", "zoneight234", "7pqrstsixteen"},
			expected: []int{29, 83, 13, 24, 42, 14, 76},
		},
		{
			name:     "real",
			input:    []string{"4b", "rvgvhpdtwo17xzbxnfjrmfjqxf", "sevenlzsrq6oneightm"},
			expected: []int{44, 27, 78},
		},
		{
			name: "real 2",
			input: []string{"qkxt88twohqzntfcsfournine53", "bmcxgsdjtl2", "74htmkdfg583srrlxbhrjv74", "sixseven6tjtsthqsr", "sqpkddjfvfn2vpncllssfqlzjkcfivetwofour1oneightltn",
				"6smrndvvbhkzpffzfggvzfznzvmkmglvptfour", "45fttcdmvpl", "862",
			},
			expected: []int{83, 22, 74, 66, 28, 64, 45, 82},
		},
		{
			name: "real 3",
			input: []string{"1threesvrfvccqpnqzvhkq1", "44five", "one61", "twoqxqlkkrfj7six51sixjfgjbfx", "4bdldfqtb6",
				"xrftwohszhtkhq9",
				"threennsixfhtgpvdnnx6kxxcpx3twoone",
				"grhbqhtl5threesevenkscph",
				"sevenbrmttfxrlm9kdvmvjgbpz7",
				"26one4one29two",
				"fone76nine",
				"four62",
				"6eightftkdjhnqdsdpone8five4two",
				"mpqltxpzqfone1ninekthxjqjf",
			},

			expected: []int{11, 45, 11, 26, 46, 29, 31, 57, 77, 22, 19, 42, 62, 19},
		},
	}

	sut := &Day01{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := sut.getRealCalibrationValues(tt.input)

			if !reflect.DeepEqual(tt.expected, actual) {
				t.Errorf("expected: %v, got: %v", tt.expected, actual)
			}
		})
	}
}
