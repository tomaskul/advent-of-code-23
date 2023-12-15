package day14

import (
	"reflect"
	"testing"
)

func TestBoard_createNew(t *testing.T) {
	// Arrange
	input := []string{
		"O....#....",
		"O.OO#....#",
		".....##...",
		"OO.#O....O",
		".O.....O#.",
		"O.#..O.#.#",
		"..O..#O..O",
		".......O..",
		"#....###..",
		"#OO..#....",
	}
	expected := [][]item{
		{'O', '.', '.', '.', '.', '#', '.', '.', '.', '.'},
		{'O', '.', 'O', 'O', '#', '.', '.', '.', '.', '#'},
		{'.', '.', '.', '.', '.', '#', '#', '.', '.', '.'},
		{'O', 'O', '.', '#', 'O', '.', '.', '.', '.', 'O'},
		{'.', 'O', '.', '.', '.', '.', '.', 'O', '#', '.'},
		{'O', '.', '#', '.', '.', 'O', '.', '#', '.', '#'},
		{'.', '.', 'O', '.', '.', '#', 'O', '.', '.', 'O'},
		{'.', '.', '.', '.', '.', '.', '.', 'O', '.', '.'},
		{'#', '.', '.', '.', '.', '#', '#', '#', '.', '.'},
		{'#', 'O', 'O', '.', '.', '#', '.', '.', '.', '.'},
	}

	// Act.
	actual := newBoard(input)

	// Assert.
	if !reflect.DeepEqual(expected, actual.data) {
		t.Errorf("expected: %v, got: %v", expected, actual)
	}
}

func TestBoard_moveNorth(t *testing.T) {
	tests := []struct {
		name          string
		input         []string
		x             int
		y             int
		expectedMove  bool
		expectedBoard [][]item
	}{
		{
			name: "can't move - empty space",
			input: []string{
				".",
				"O",
			},
			x: 0, y: 0,
			expectedMove: false,
			expectedBoard: [][]item{
				{'.'},
				{'O'},
			},
		},
		{
			name: "can move - moveable",
			input: []string{
				".",
				"O",
			},
			x: 0, y: 1,
			expectedMove: true,
			expectedBoard: [][]item{
				{'O'},
				{'.'},
			},
		},
		{
			name: "can't move - immoveable",
			input: []string{
				"#",
				"O",
			},
			x: 0, y: 1,
			expectedMove: false,
			expectedBoard: [][]item{
				{'#'},
				{'O'},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			sut := newBoard(tc.input)
			actual := sut.moveNorth(tc.y, tc.x)
			if tc.expectedMove != actual {
				t.Fatalf("expected: %v, got: %v", tc.expectedMove, actual)
			}
			if !reflect.DeepEqual(tc.expectedBoard, sut.data) {
				t.Errorf("expected: %v, got: %v", tc.expectedBoard, sut.data)
			}
		})
	}
}

func TestBoad_TiltNorth(t *testing.T) {
	// Arrange
	input := []string{
		"O....#....",
		"O.OO#....#",
		".....##...",
		"OO.#O....O",
		".O.....O#.",
		"O.#..O.#.#",
		"..O..#O..O",
		".......O..",
		"#....###..",
		"#OO..#....",
	}
	expected := [][]item{
		{'O', 'O', 'O', 'O', '.', '#', '.', 'O', '.', '.'},
		{'O', 'O', '.', '.', '#', '.', '.', '.', '.', '#'},
		{'O', 'O', '.', '.', 'O', '#', '#', '.', '.', 'O'},
		{'O', '.', '.', '#', '.', 'O', 'O', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '#', '.'},
		{'.', '.', '#', '.', '.', '.', '.', '#', '.', '#'},
		{'.', '.', 'O', '.', '.', '#', '.', 'O', '.', 'O'},
		{'.', '.', 'O', '.', '.', '.', '.', '.', '.', '.'},
		{'#', '.', '.', '.', '.', '#', '#', '#', '.', '.'},
		{'#', '.', '.', '.', '.', '#', '.', '.', '.', '.'},
	}

	// Act.
	sut := newBoard(input)

	sut.TiltBoardNorth()

	if !reflect.DeepEqual(expected, sut.data) {
		t.Errorf("expected: %v, got: %v", expected, sut.data)
	}
}
