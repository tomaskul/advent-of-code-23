package day09

import (
	"fmt"

	"github.com/tomaskul/advent-of-code-23/util"
)

type Day09 struct {
	rows []string
}

func NewDay09Solution(sessionCookie string) *Day09 {
	rows, _ := util.GetCachedRows("https://adventofcode.com/2023/day/14/input", "09", ".txt", sessionCookie)
	return &Day09{
		rows: rows,
	}
}

func (s *Day09) PrintPart1() {
	input := []int{0, 3, 6, 9, 12, 15}
	fmt.Println(nextInSequence(input))
}

func (s *Day09) PrintPart2() {
}

func nextInSequence(sequence []int) int {
	if util.All(sequence, func(val int) bool { return val == 0 }) {
		return 0
	}

	s2 := make([]int, len(sequence)-1)
	for i := 1; i < len(sequence); i++ {
		s2[i-1] = sequence[i] - sequence[i-1]
	}

	return s2[len(s2)-1]
}
