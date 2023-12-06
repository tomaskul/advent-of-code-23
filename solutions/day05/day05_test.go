package day05

import (
	"reflect"
	"testing"
)

func Test_parseMap(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected map[int]idRangeDef
	}{
		{
			name: "sample seed-to-soil",
			input: []string{
				"50 98 2",
				"52 50 48",
			},
			expected: map[int]idRangeDef{
				50: {srcRangeStart: 98, rangeLen: 2},
				52: {srcRangeStart: 50, rangeLen: 48},
			},
		},
		{
			name: "sample fertilizer-to-water",
			input: []string{
				"49 53 8",
				"0 11 42",
				"42 0 7",
				"57 7 4",
			},
			expected: map[int]idRangeDef{
				49: {srcRangeStart: 53, rangeLen: 8},
				0:  {srcRangeStart: 11, rangeLen: 42},
				42: {srcRangeStart: 0, rangeLen: 7},
				57: {srcRangeStart: 7, rangeLen: 4},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, actual := parseMap(tt.input)
			if !reflect.DeepEqual(tt.expected, actual) {
				t.Errorf("expected: %+v, got: %+v", tt.expected, actual)
			}
		})
	}
}

func Test_parseData(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected data
	}{
		{
			name: "sample truncated",
			input: []string{
				"seeds: 79 14 55 13",
				"",
				"seed-to-soil map:",
				"50 98 2",
				"52 50 48",
				"",
				"soil-to-fertilizer map:",
				"0 15 37",
				"37 52 2",
				"39 0 15",
				"",
			},
			expected: data{
				seeds: []int{79, 14, 55, 13},
				paths: []lookup{
					{50: {srcRangeStart: 98, rangeLen: 2}, 52: {srcRangeStart: 50, rangeLen: 48}},
					{0: {srcRangeStart: 15, rangeLen: 37}, 37: {srcRangeStart: 52, rangeLen: 2}, 39: {srcRangeStart: 0, rangeLen: 15}},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := parseData(tt.input)
			if !reflect.DeepEqual(tt.expected.seeds, actual.seeds) {
				t.Errorf("expected: %+v, got: %+v", tt.expected, actual)
			}

			if len(tt.expected.paths) != len(actual.paths) {
				t.Errorf("expected paths: %d paths, got: %d", len(tt.expected.paths), len(actual.paths))
			}

			for i, v := range tt.expected.paths {
				path := actual.paths[i]
				if !reflect.DeepEqual(v, path) {
					t.Errorf("expected: %+v, got: %+v", v, path)
				}
			}

			// if !reflect.DeepEqual(tt.expected.paths, actual.paths) {
			// 	t.Errorf("expected: %+v, got: %+v", tt.expected, actual)
			// }
		})
	}
}

func Test_getDestinationId(t *testing.T) {
	tests := []struct {
		name        string
		inputSrcId  int
		inputLookup lookup
		expected    int
	}{
		{
			name:       "sample seed-to-soil",
			inputSrcId: 98,
			inputLookup: lookup{
				50: {srcRangeStart: 98, rangeLen: 2},
				52: {srcRangeStart: 50, rangeLen: 48},
			},
			expected: 50,
		},
		{
			name:       "sample seed-to-soil 79",
			inputSrcId: 79,
			inputLookup: lookup{
				50: {srcRangeStart: 98, rangeLen: 2},
				52: {srcRangeStart: 50, rangeLen: 48},
			},
			expected: 81,
		},
		{
			name:       "sample seed-to-soil 14",
			inputSrcId: 14,
			inputLookup: lookup{
				50: {srcRangeStart: 98, rangeLen: 2},
				52: {srcRangeStart: 50, rangeLen: 48},
			},
			expected: 14,
		},
		{
			name:       "sample seed-to-soil 55",
			inputSrcId: 55,
			inputLookup: lookup{
				50: {srcRangeStart: 98, rangeLen: 2},
				52: {srcRangeStart: 50, rangeLen: 48},
			},
			expected: 57,
		},
		{
			name:       "sample seed-to-soil 13",
			inputSrcId: 13,
			inputLookup: lookup{
				50: {srcRangeStart: 98, rangeLen: 2},
				52: {srcRangeStart: 50, rangeLen: 48},
			},
			expected: 13,
		},
		{
			name:       "sample fertilizer-to-water",
			inputSrcId: 12,
			inputLookup: lookup{
				49: {srcRangeStart: 53, rangeLen: 8},
				0:  {srcRangeStart: 11, rangeLen: 42},
				42: {srcRangeStart: 0, rangeLen: 7},
				57: {srcRangeStart: 7, rangeLen: 4},
			},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.inputLookup.getDestinationId(tt.inputSrcId)
			if tt.expected != actual {
				t.Errorf("expected: %d, got: %d", tt.expected, actual)
			}
		})
	}
}

func Test_traverse(t *testing.T) {
	tests := []struct {
		name     string
		input    data
		expected []int
	}{
		{
			name: "sample seed-to-soil",
			input: data{
				seeds: []int{79, 14, 55, 13},
				paths: []lookup{
					{50: {srcRangeStart: 98, rangeLen: 2}, 52: {srcRangeStart: 50, rangeLen: 48}},
					{0: {srcRangeStart: 15, rangeLen: 37}, 37: {srcRangeStart: 52, rangeLen: 2}, 39: {srcRangeStart: 0, rangeLen: 15}},
					{
						49: {srcRangeStart: 53, rangeLen: 8},
						0:  {srcRangeStart: 11, rangeLen: 42},
						42: {srcRangeStart: 0, rangeLen: 7},
						57: {srcRangeStart: 7, rangeLen: 4},
					},
					{88: {srcRangeStart: 18, rangeLen: 7}, 18: {srcRangeStart: 25, rangeLen: 70}},
					{
						45: {srcRangeStart: 77, rangeLen: 23},
						81: {srcRangeStart: 45, rangeLen: 19},
						68: {srcRangeStart: 64, rangeLen: 13},
					},
					{0: {srcRangeStart: 69, rangeLen: 1}, 1: {srcRangeStart: 0, rangeLen: 69}},
					{60: {srcRangeStart: 56, rangeLen: 37}, 56: {srcRangeStart: 93, rangeLen: 4}},
				},
			},
			expected: []int{82, 43, 86, 35},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := traverse(tt.input)
			if !reflect.DeepEqual(tt.expected, actual) {
				t.Errorf("expected: %+v, got: %+v", tt.expected, actual)
			}
		})
	}
}
