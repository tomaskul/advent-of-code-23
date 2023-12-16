package day14

import (
	"testing"
)

func Test_calculateLoad(t *testing.T) {
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

	// Act.
	sut := newBoard(input)
	sut.TiltBoardNorth()
	actual := calculateLoad(sut)

	// Assert.
	if actual != 136 {
		t.Errorf("expected: 136, got: %d", actual)
	}
}
