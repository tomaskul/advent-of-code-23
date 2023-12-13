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
			subjectRow := row[afterIndex:]

			isAdj, newAfterIndex := isAdjacentToSymbol(number, above, subjectRow, below)
			afterIndex += newAfterIndex
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

	leftIndex := subjectIndex
	if subjectIndex-1 > -1 {
		leftIndex -= 1
	}
	rightIndex := subjectEndIndex
	if subjectEndIndex+1 <= len(subjectRow) {
		rightIndex += 1
	}

	if nonFillerRegex.MatchString(subjectRow[leftIndex:rightIndex]) {
		return true, subjectEndIndex
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

	// fmt.Printf("subject:%q\ncurRow:%q\n", subject, subjectRow)
	// fmt.Printf("above:  %v\n", aboveSearchSpace)
	// fmt.Printf("current:%v\n", subjectRow[leftIndex:rightIndex])
	// fmt.Printf("below:  %v\n\n", belowSearchSpace)

	return false, subjectEndIndex
}

func (s *Day03) PrintPart2() {
	fmt.Println("WIP")
	fmt.Println(util.Sum(numbersAdjacentToGears(s.rows)))
}

func numbersAdjacentToGears(input []string) []int {
	// gearMatch, _ := regexp.Compile("\\*")
	// numMatch, _ := regexp.Compile("\\d+")
	result := make([]int, 0)

	return result
}
