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

var cardFacesPt1 = []rune{'A', 'K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2'}
var cardFacesPt2 = []rune{'A', 'K', 'Q', 'T', '9', '8', '7', '6', '5', '4', '3', '2', 'J'}

const (
	HighCardType = iota
	OnePairType
	TwoPairType
	ThreeOfKindType
	FullHouseType
	FourOfKindType
	FiveOfKindType
)

func NewDay07Solution(sessionCookie string) *Day07 {
	rows, _ := util.GetCachedRows("https://adventofcode.com/2023/day/7/input", "7", ".txt", sessionCookie)
	return &Day07{
		rows: rows,
	}
}

func (s *Day07) PrintPart1() {
	fmt.Println(calculateWinnings(sortPt1(parse(s.rows))))
}

func parse(rows []string) []item {
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
		iType, jType := items[i].hand.getType(false), items[j].hand.getType(false)
		if iType == jType {

			iValue, jValue := -1, -1
			for k := 0; k < 5; k++ {
				for idx, face := range cardFacesPt1 {
					//cardFaces are in descending value.
					if items[i].hand.value[k] == byte(face) {
						iValue = len(cardFacesPt1) - idx
					}
					if items[j].hand.value[k] == byte(face) {
						jValue = len(cardFacesPt1) - idx
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

func sortPt2(items []item) []item {
	sort.Slice(items, func(i, j int) bool {
		iType, jType := items[i].hand.getType(true), items[j].hand.getType(true)
		if iType == jType {
			iValue, jValue := -1, -1
			for k := 0; k < 5; k++ {
				for idx, face := range cardFacesPt2 {
					//cardFaces are in descending value.
					if items[i].hand.value[k] == byte(face) {
						iValue = len(cardFacesPt2) - idx
					}
					if items[j].hand.value[k] == byte(face) {
						jValue = len(cardFacesPt2) - idx
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
	fmt.Println(calculateWinnings(sortPt2(parse(s.rows))))
}
