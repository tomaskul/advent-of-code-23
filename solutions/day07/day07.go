package day07

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/tomaskul/advent-of-code-23/util"
)

type Day07 struct {
	rows []string
}

var cardFaces = []rune{'A', 'K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2'}

type item struct {
	hand card
	bid  int
}

type card struct {
	value  string
	lookup map[byte]int
}

func newCard(input string) card {
	lookup := map[byte]int{}
	for i := 0; i < len(input); i++ {
		_, ok := lookup[input[i]]
		if ok {
			lookup[input[i]] += 1
		} else {
			lookup[input[i]] = 1
		}
	}
	return card{
		value:  input,
		lookup: lookup,
	}
}

func (c card) getType() int {
	if c.isFiveOfAKind() {
		return 5
	} else if c.isFourOfAKind() {
		return 4
	} else if c.isThreeOfAKind() {
		return 3
	} else if c.isTwoPair() {
		return 2
	} else if c.isOnePair() {
		return 1
	} else if c.isHighCard() {
		return 0
	} else {
		fmt.Printf("debug: unable to determine %q type\n", c.value)
		return -1
	}
}

func (c card) isFiveOfAKind() bool {
	return len(c.lookup) == 1
}

func (c card) isFourOfAKind() bool {
	return len(c.lookup) == 2 &&
		(c.lookup[c.value[0]] == 1 || c.lookup[c.value[0]] == 4)
}

func (c card) isFullHouse() bool {
	return len(c.lookup) == 2 &&
		(c.lookup[c.value[0]] == 3 || c.lookup[c.value[0]] == 2)
}

func (c card) isThreeOfAKind() bool {
	if len(c.lookup) != 3 {
		return false
	}
	for _, v := range c.lookup {
		if v == 2 {
			return false
		}
	}
	return true
}

func (c card) isTwoPair() bool {
	if c.isFiveOfAKind() || c.isFourOfAKind() || c.isThreeOfAKind() {
		return false
	}

	pairOne := false
	for _, v := range c.lookup {
		if v == 2 {
			if !pairOne {
				pairOne = true
			} else {
				return true
			}
		}
	}
	return false
}

func (c card) isOnePair() bool {
	pairCounts := 0
	for _, v := range c.lookup {
		if v == 2 {
			pairCounts++
		}
	}
	return pairCounts == 1
}

func (c card) isHighCard() bool {
	return len(c.lookup) == 5
}

func NewDay07Solution(sessionCookie string) *Day07 {
	rows, _ := util.GetCachedRows("https://adventofcode.com/2023/day/7/input", "7", ".txt", sessionCookie)
	return &Day07{
		rows: rows,
	}
}

func (s *Day07) PrintPart1() {
	items := parsePt1(s.rows)
	fmt.Println(calculateWinnings(items))
}

func parsePt1(rows []string) []item {
	items := make([]item, len(rows))
	for i, row := range rows {
		tokens := strings.Split(row, " ")
		bid, _ := strconv.Atoi(tokens[1])

		items[i] = item{
			hand: newCard(tokens[0]),
			bid:  bid,
		}
	}
	return items
}

func sortPt1(items []item) []item {
	sort.Slice(items, func(i, j int) bool {
		iType, jType := items[i].hand.getType(), items[j].hand.getType()
		if iType == jType {

			iValue, jValue := -1, -1
			for k := 0; k < 5; k++ {
				for idx, face := range cardFaces {
					if items[i].hand.value[k] == byte(face) {
						iValue = len(cardFaces) - idx
					}
					if items[j].hand.value[k] == byte(face) {
						jValue = len(cardFaces) - idx
					}
				}
				if iValue == jValue {
					continue
				} else {
					break
				}

			}
			return iValue < jValue
		} else {
			return iType < jType
		}
	})
	return items
}

func calculateWinnings(items []item) int {
	winnings := items[0].bid
	for i := 1; i < len(items); i++ {

		add := items[i].bid * (i + 1)
		winnings += add
	}
	return winnings
}

func (s *Day07) PrintPart2() {
}
