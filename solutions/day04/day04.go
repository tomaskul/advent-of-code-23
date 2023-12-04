package day04

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tomaskul/advent-of-code-23/util"
)

type Day04 struct {
	SessionCookie string
	rows          []string
}

type scratchcard struct {
	id            int
	winningValues []int
	actualValues  []int
}

func (c *scratchcard) calculatePoints() int {
	result := 0

	for _, actual := range c.actualValues {
		for _, winning := range c.winningValues {
			if actual == winning {
				if result == 0 {
					result = 1
				} else {
					result = result * 2
				}
			}
		}
	}

	return result
}

func (s *Day04) getData() {
	if s.rows == nil {
		s.rows = util.GetRows("https://adventofcode.com/2023/day/4/input", s.SessionCookie)
	}
}

func (s *Day04) PrintPart1() {
	s.getData()

	scratchcards := parseInputData(s.rows)
	actual := make([]int, 0)
	for _, scratchcard := range scratchcards {
		actual = append(actual, scratchcard.calculatePoints())
	}
	fmt.Println(util.Sum(actual))
}

func parseInputData(rows []string) []scratchcard {
	cardsPlayed := make([]scratchcard, len(rows))
	for i, row := range rows {
		card := scratchcard{}

		tokens := strings.Split(row, ":")
		card.id, _ = strconv.Atoi(strings.TrimPrefix(tokens[0], "Card "))

		sides := strings.Split(tokens[1], "|")
		card.winningValues = getIntValues(strings.Split(sides[0], " "))
		card.actualValues = getIntValues(strings.Split(sides[1], " "))

		cardsPlayed[i] = card
	}

	return cardsPlayed
}

func getIntValues(input []string) []int {
	result := make([]int, 0)
	for _, v := range input {
		integerValue, err := strconv.Atoi(v)
		if err != nil {
			continue
		}
		result = append(result, integerValue)
	}
	return result
}

func (s *Day04) PrintPart2() {
	s.getData()
}
