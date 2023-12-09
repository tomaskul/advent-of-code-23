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
	rows, _ := util.GetCachedRows("https://adventofcode.com/2023/day/3/input", "5", ".txt", sessionCookie)
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
	numMatch, _ := regexp.Compile("\\d{1,7}")
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

		isAdj, afterIndex := false, 0
		for _, number := range numbersInRow {
			above, below := "", ""
			if rowAbove != "" {
				above = fmt.Sprintf("%s", rowAbove[afterIndex:])
			}
			if rowBelow != "" {
				below = fmt.Sprintf("%s", rowBelow[afterIndex:])
			}
			curr := fmt.Sprintf("%s", row[afterIndex:])

			if i < 3 {
				fmt.Printf("number: %q\n", number)
				fmt.Printf("above:  %v\n", above)
				fmt.Printf("current:%v\n", curr)
				fmt.Printf("below:  %v\n\n\n", below)
			}
			isAdj, afterIndex = isAdjacentToSymbol(number, above, curr, below)
			if isAdj {
				value, _ := strconv.Atoi(number)
				result = append(result, value)
			}
		}

		//fmt.Printf("rowIndex: %d, map: %v\n", i, hasDuplicates(numbersInRow))
	}

	return result
}

func isAdjacentToSymbol(subject, rowAbove, subjectRow, rowBelow string) (bool, int) {
	nonFillerRegex, _ := regexp.Compile("[^0-9.]")
	subjectIndex := strings.Index(subjectRow, subject)
	subjectLength := len(subject)

	leftIndex := subjectIndex
	if subjectIndex-1 > -1 {
		leftIndex = subjectIndex - 1
		if nonFillerRegex.Match([]byte{subjectRow[leftIndex]}) {
			return true, subjectIndex + len(subject)
		}
	}

	rightIndex := subjectIndex + subjectLength - 1
	if subjectIndex+subjectLength < len(subjectRow) {
		if nonFillerRegex.Match([]byte{subjectRow[subjectIndex+subjectLength]}) {
			return true, subjectIndex + len(subject)
		}
		rightIndex = subjectIndex + subjectLength + 1
	}

	if rowAbove != "" {
		searchSpace := rowAbove[leftIndex:rightIndex]
		if nonFillerRegex.Match([]byte(searchSpace)) {
			return true, subjectIndex + len(subject)
		}
	}

	if rowBelow != "" {
		searchSpace := rowBelow[leftIndex:rightIndex]
		if nonFillerRegex.Match([]byte(searchSpace)) {
			return true, subjectIndex + len(subject)
		}
	}

	return false, subjectIndex + len(subject)
}

func hasDuplicates(values []string) map[string]int {
	result := make(map[string]int)
	for _, v := range values {
		_, ok := result[v]
		if ok {
			result[v] = result[v] + 1
		} else {
			result[v] = 1
		}
	}
	return result
}
