package day14

import (
	"fmt"
	"regexp"

	"github.com/tomaskul/advent-of-code-23/util"
)

type Day14 struct {
	rows []string
}

func NewDay14Solution(sessionCookie string) *Day14 {
	rows, _ := util.GetCachedRows("https://adventofcode.com/2023/day/14/input", "14", ".txt", sessionCookie)
	return &Day14{
		rows: rows,
	}
}

func (s *Day14) PrintPart1() {
	board := newBoard(s.rows)
	board.TiltBoardNorth()
	fmt.Println(calculateLoad(board))
}

func calculateLoad(b *board) int {
	var result int

	moveableRegex := regexp.MustCompile("[O]")

	for y, row := range b.data {
		matches := moveableRegex.FindAllString(string(row), -1)
		if len(matches) == 0 {
			continue
		}

		result += (len(b.data) - y) * len(matches)
	}

	return result
}

func (s *Day14) PrintPart2() {
	// After certain number of iterations a consistent pattern emerges and repeats into infinity.
	// just need to find the right point in the small pattern at 1000000000th iteration.
	board := newBoard(s.rows)
	board.SpinCycle(1000)
	fmt.Println(calculateLoad(board))
}
