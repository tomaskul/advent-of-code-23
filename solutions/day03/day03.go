package day03

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/tomaskul/advent-of-code-23/util"
)

type Day03 struct {
	SessionCookie string
	rows          []string
}

func NewDay03Solution(sessionCookie string) *Day03 {
	rows, _ := util.GetCachedRows("https://adventofcode.com/2023/day/3/input", "3", ".txt", sessionCookie)
	return &Day03{
		rows: rows,
	}
}

func (s *Day03) PrintPart1() {
	fmt.Println(util.Sum(numbersWithAdjacentSymbols(s.rows)))
}

func (s *Day03) PrintPart2() {
	fmt.Println("WIP")
}

func numbersWithAdjacentSymbols(input []string) []int {
	numMatch, _ := regexp.Compile("\\d+")
	result := make([]int, 0)

	for i, row := range input {
		numbersInRow := numMatch.FindAllString(row, -1)
		if len(numbersInRow) == 0 {
			continue
		}

		rowAbove, rowBelow := "", ""
		if i-1 > -1 {
			rowAbove = input[i-1]
		}
		if i+1 < len(input) {
			rowBelow = input[i+1]
		}

		var afterIndex int
		for _, number := range numbersInRow {
			above, below := "", ""
			if rowAbove != "" {
				above = rowAbove[afterIndex:]
			}
			if rowBelow != "" {
				below = rowBelow[afterIndex:]
			}
			curr := row[afterIndex:]

			var isAdj bool
			isAdj, afterIndex = isAdjacentToSymbol(number, above, curr, below)
			if isAdj {
				value, _ := strconv.Atoi(number)
				result = append(result, value)
			}
		}
	}
	return result
}

func isAdjacentToSymbol(subject, rowAbove, subjectRow, rowBelow string) (bool, int) {
	nonFillerRegex, _ := regexp.Compile("[^.0-9]")

	subjectIndex := strings.Index(subjectRow, subject)
	subjectEndIndex := subjectIndex + len(subject)

	// Check left side, keep track of the index for diagonal checks.
	leftIndex := subjectIndex
	if subjectIndex-1 > -1 {
		leftIndex -= 1
		if nonFillerRegex.Match([]byte{subjectRow[leftIndex]}) {
			return true, subjectEndIndex
		}
	}

	// Check right side, keep track of the index for diagonal checks.
	rightIndex := subjectEndIndex
	if subjectEndIndex+1 <= len(subjectRow) {
		if nonFillerRegex.Match([]byte{subjectRow[subjectEndIndex]}) {
			return true, subjectEndIndex
		}
		rightIndex += 1
	}

	// Above, below + diagonal checks.
	aboveSearchSpace, belowSearchSpace := "", ""

	if rowAbove != "" {
		aboveSearchSpace = rowAbove[leftIndex:rightIndex]
		if nonFillerRegex.MatchString(aboveSearchSpace) {
			return true, subjectEndIndex
		}
	}

	if rowBelow != "" {
		belowSearchSpace = rowBelow[leftIndex:rightIndex]
		if nonFillerRegex.MatchString(belowSearchSpace) {
			return true, subjectEndIndex
		}
	}

	// fmt.Printf("subject:%q\n", subject)
	// fmt.Printf("above:  %v\n", aboveSearchSpace)
	// fmt.Printf("current:%v\n", subjectRow[leftIndex:rightIndex])
	// fmt.Printf("below:  %v\n\n", belowSearchSpace)
	// fmt.Printf("current:%v\n\n\n", subjectRow)

	return false, subjectEndIndex
}
