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

type node struct {
	parent   *scratchcard
	data     scratchcard
	children []node
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

	fmt.Println(util.Sum(evaluateScratchcards(parseInputData(s.rows))))
}

func evaluateScratchcards(input []scratchcard) []int {
	result := make([]int, 0)
	for _, scratchcard := range input {
		result = append(result, scratchcard.calculatePoints())
	}
	return result
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

func traverseCards(allCards []scratchcard) {
	cardLookup := createCardLookup(allCards)

	root := node{parent: nil, data: cardLookup[1]}
	root.children = root.getChildren(cardLookup)
	for _, childNode := range root.children {

	}
}

func createCardLookup(cards []scratchcard) map[int]scratchcard {
	result := make(map[int]scratchcard, 0)
	for _, card := range cards {
		result[card.id] = card
	}
	return result
}

func (n *node) getChildren(cardLookup map[int]scratchcard) []node {
	matchingNumbers := n.data.calculatePoints()
	if matchingNumbers == 0 {
		return []node{}
	}
	result := []node{}
	for i := 1; i < matchingNumbers+1; i++ {
		result = append(result, node{parent: &n.data, data: cardLookup[n.data.id+i]})
	}
	return result
}

// func findNewChildNodes(dest *[]node, data []scratchcard, cardLookup map[int]scratchcard) {

// }

// func pathThroughScratchcards(allCards []scratchcard) map[int]int {
// 	result := make(map[int]int)
// 	for _, scCard := range allCards {
// 		result[scCard.id] = 0
// 	}

// 	cardsCurrentlyPossessed := []scratchcard{allCards[0]}

// 	inPlay = true
// 	for inPlay {
// 		wins := evaluateScratchcardsPt2(cardsCurrentlyPossessed)
// 		for id, nextSet := range wins {
// 			result[id] = result[id] + nextSet

// 			cardsCurrentlyPossessed = append(cardsCurrentlyPossessed, allCards)
// 		}
// 	}

// 	return result
// }

// func evaluateScratchcardsPt2(input []scratchcard) map[int]int {
// 	result := make(map[int]int, 0)
// 	for _, scratchcard := range input {
// 		nextSet := scratchcard.calculatePoints()
// 		if nextSet == 0 {
// 			continue
// 		}

// 		currentCount, ok := result[scratchcard.id]
// 		if ok {
// 			result[scratchcard.id] = currentCount + nextSet
// 		} else {
// 			result[scratchcard.id] = nextSet
// 		}
// 	}
// 	return result
// }
