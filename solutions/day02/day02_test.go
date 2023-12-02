package day02

import (
	"reflect"
	"testing"
)

func Test_parseInputData(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []game
	}{
		{
			name: "sample",
			input: []string{
				"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
				"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
				"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
				"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
				"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			},
			expected: []game{
				{
					id:   1,
					sets: set{rgb{b: 3, r: 4}, rgb{r: 1, g: 2, b: 6}, rgb{g: 2}},
				},
				{
					id:   2,
					sets: set{rgb{b: 1, g: 2}, rgb{g: 3, b: 4, r: 1}, rgb{g: 1, b: 1}},
				},
				{
					id:   3,
					sets: set{rgb{g: 8, b: 6, r: 20}, rgb{b: 5, r: 4, g: 13}, rgb{g: 5, r: 1}},
				},
				{
					id:   4,
					sets: set{rgb{g: 1, r: 3, b: 6}, rgb{g: 3, r: 6}, rgb{g: 3, b: 15, r: 14}},
				},
				{
					id:   5,
					sets: set{rgb{r: 6, b: 1, g: 3}, rgb{b: 2, r: 1, g: 2}},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := parseInputData(tt.input)
			if !reflect.DeepEqual(tt.expected, actual) {
				t.Errorf("expected: %+v, got: %+v", tt.expected, actual)
			}
		})
	}
}

func Test_countPossibleGames(t *testing.T) {
	tests := []struct {
		name             string
		inputGames       []game
		inputExpectation rgb
		expected         []int
	}{
		{
			name: "sample",
			inputGames: []game{
				{
					id:   1,
					sets: set{rgb{b: 3, r: 4}, rgb{r: 1, g: 2, b: 6}, rgb{g: 2}},
				},
				{
					id:   2,
					sets: set{rgb{b: 1, g: 2}, rgb{g: 3, b: 4, r: 1}, rgb{g: 1, b: 1}},
				},
				{
					id:   3,
					sets: set{rgb{g: 8, b: 6, r: 20}, rgb{b: 5, r: 4, g: 13}, rgb{g: 5, r: 1}},
				},
				{
					id:   4,
					sets: set{rgb{g: 1, r: 3, b: 6}, rgb{g: 3, r: 6}, rgb{g: 3, b: 15, r: 14}},
				},
				{
					id:   5,
					sets: set{rgb{r: 6, b: 1, g: 3}, rgb{b: 2, r: 1, g: 2}},
				},
			},
			inputExpectation: rgb{r: 12, g: 13, b: 14},
			expected:         []int{1, 2, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := countPossibleGames(tt.inputExpectation, tt.inputGames)
			if !reflect.DeepEqual(tt.expected, actual) {
				t.Errorf("expected: %v, got: %v", tt.expected, actual)
			}
		})
	}
}
