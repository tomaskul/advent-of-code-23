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

func (c *scratchcard) countMatches() int {
	result := 0

	for _, actual := range c.actualValues {
		for _, winning := range c.winningValues {
			if actual == winning {
				result++
				break
			}
		}
	}

	return result
}

func (s *Day04) getData() {
	if s.rows == nil {
		s.rows, _ = util.GetCachedRows("https://adventofcode.com/2023/day/4/input", "4", ".txt", s.SessionCookie)
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
		card.id, _ = strconv.Atoi(strings.TrimSpace(strings.TrimPrefix(tokens[0], "Card")))

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

	fmt.Println("PS. A pretty slow implementation, give it time.")
	cardMap := traverseCards(parseInputData(s.rows))
	total := 0
	for _, v := range cardMap {
		total += v
	}
	fmt.Println(total)
}

func traverseCards(allCards []scratchcard) map[int]int {
	cardLookup := createCardLookup(allCards)

	result := make(map[int]int)
	for _, v := range cardLookup {
		root := node{parent: nil, data: v}
		tree := recurseTree(&root, cardLookup)
		originalCardMapCount := recurseTreeCountCards(tree)
		for id, count := range originalCardMapCount {
			_, ok := result[id]
			if ok {
				result[id] = result[id] + count
			} else {
				result[id] = count
			}
		}
	}

	return result
}

func createCardLookup(cards []scratchcard) map[int]scratchcard {
	result := make(map[int]scratchcard, 0)
	for _, card := range cards {
		result[card.id] = card
	}
	return result
}

func recurseTree(n *node, cardLookup map[int]scratchcard) node {
	children := getChildren(n, cardLookup)
	if children == nil || len(children) == 0 {
		return *n
	}

	resultChildren := []node{}
	for _, child := range children {
		resultChildren = append(resultChildren, recurseTree(&child, cardLookup))
	}
	n.children = resultChildren
	return *n
}

func getChildren(n *node, cardLookup map[int]scratchcard) []node {
	matchingNumbers := n.data.countMatches()
	if matchingNumbers == 0 {
		return []node{}
	}
	result := []node{}
	for i := 0; i < matchingNumbers; i++ {
		scratchcard, ok := cardLookup[n.data.id+i+1]
		if !ok {
			fmt.Printf("unable to find card with id: %d+%d+1\n", n.data.id, i)
			continue
		}
		result = append(result, node{parent: &n.data, data: scratchcard})
	}
	return result
}

func recurseTreeCountCards(n node) map[int]int {
	result := map[int]int{n.data.id: 1}
	if n.children == nil || len(n.children) == 0 {
		return result
	}

	for _, child := range n.children {
		childMap := recurseTreeCountCards(child)
		for id, count := range childMap {
			_, ok := result[id]
			if ok {
				result[id] = result[id] + count
			} else {
				result[id] = count
			}
		}
	}

	return result
}
