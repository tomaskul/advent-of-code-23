package day04

import (
	"reflect"
	"testing"
)

func Test_parseInputData(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []scratchcard
	}{
		{
			name: "sample",
			input: []string{
				"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
				"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
				"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
				"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
				"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
				"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
			},
			expected: []scratchcard{
				{
					id:            1,
					winningValues: []int{41, 48, 83, 86, 17},
					actualValues:  []int{83, 86, 6, 31, 17, 9, 48, 53},
				},
				{
					id:            2,
					winningValues: []int{13, 32, 20, 16, 61},
					actualValues:  []int{61, 30, 68, 82, 17, 32, 24, 19},
				},
				{
					id:            3,
					winningValues: []int{1, 21, 53, 59, 44},
					actualValues:  []int{69, 82, 63, 72, 16, 21, 14, 1},
				},
				{
					id:            4,
					winningValues: []int{41, 92, 73, 84, 69},
					actualValues:  []int{59, 84, 76, 51, 58, 5, 54, 83},
				},
				{
					id:            5,
					winningValues: []int{87, 83, 26, 28, 32},
					actualValues:  []int{88, 30, 70, 12, 93, 22, 82, 36},
				},
				{
					id:            6,
					winningValues: []int{31, 18, 13, 56, 72},
					actualValues:  []int{74, 77, 10, 23, 35, 67, 36, 11},
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

func TestScratchcard_evaluateScratchcards(t *testing.T) {
	tests := []struct {
		name     string
		input    []scratchcard
		expected []int
	}{
		{
			name: "sample",
			input: []scratchcard{
				{
					id:            1,
					winningValues: []int{41, 48, 83, 86, 17},
					actualValues:  []int{83, 86, 6, 31, 17, 9, 48, 53},
				},
				{
					id:            2,
					winningValues: []int{13, 32, 20, 16, 61},
					actualValues:  []int{61, 30, 68, 82, 17, 32, 24, 19},
				},
				{
					id:            3,
					winningValues: []int{1, 21, 53, 59, 44},
					actualValues:  []int{69, 82, 63, 72, 16, 21, 14, 1},
				},
				{
					id:            4,
					winningValues: []int{41, 92, 73, 84, 69},
					actualValues:  []int{59, 84, 76, 51, 58, 5, 54, 83},
				},
				{
					id:            5,
					winningValues: []int{87, 83, 26, 28, 32},
					actualValues:  []int{88, 30, 70, 12, 93, 22, 82, 36},
				},
				{
					id:            6,
					winningValues: []int{31, 18, 13, 56, 72},
					actualValues:  []int{74, 77, 10, 23, 35, 67, 36, 11},
				},
			},
			expected: []int{
				8, 2, 2, 1, 0, 0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := evaluateScratchcards(tt.input)
			if !reflect.DeepEqual(tt.expected, actual) {
				t.Errorf("expected: %+v, got: %+v", tt.expected, actual)
			}
		})
	}
}

func Test_pathThroughScratchcards(t *testing.T) {
	tests := []struct {
		name              string
		allScratchcards   []scratchcard
		expectedFinalHand map[int]int
	}{
		{
			name: "sample",
			allScratchcards: []scratchcard{
				{
					id:            1,
					winningValues: []int{41, 48, 83, 86, 17},
					actualValues:  []int{83, 86, 6, 31, 17, 9, 48, 53},
				},
				{
					id:            2,
					winningValues: []int{13, 32, 20, 16, 61},
					actualValues:  []int{61, 30, 68, 82, 17, 32, 24, 19},
				},
				{
					id:            3,
					winningValues: []int{1, 21, 53, 59, 44},
					actualValues:  []int{69, 82, 63, 72, 16, 21, 14, 1},
				},
				{
					id:            4,
					winningValues: []int{41, 92, 73, 84, 69},
					actualValues:  []int{59, 84, 76, 51, 58, 5, 54, 83},
				},
				{
					id:            5,
					winningValues: []int{87, 83, 26, 28, 32},
					actualValues:  []int{88, 30, 70, 12, 93, 22, 82, 36},
				},
				{
					id:            6,
					winningValues: []int{31, 18, 13, 56, 72},
					actualValues:  []int{74, 77, 10, 23, 35, 67, 36, 11},
				},
			},
			expectedFinalHand: map[int]int{
				1: 1,
				2: 2,
				3: 4,
				4: 8,
				5: 14,
				6: 1,
			},
		},
	}

}
