package day15

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tomaskul/advent-of-code-23/util"
)

type Day15 struct {
	rows []string
}

func NewDay15Solution(sessionCookie string) *Day15 {
	rows, _ := util.GetCachedRows("https://adventofcode.com/2023/day/15/input", "15", ".txt", sessionCookie)
	return &Day15{
		rows: rows,
	}
}

func (s *Day15) PrintPart1() {
	fmt.Println(hashSumOfTokens(strings.Split(s.rows[0], ",")))
}

func hashSumOfTokens(values []string) int {
	hashValues := make([]int, 0)
	for _, v := range values {
		hashValues = append(hashValues, hash(v))
	}
	return util.Sum(hashValues)
}

func hash(value string) int {
	var currentValue int
	for _, v := range value {
		currentValue += int(v)
		currentValue *= 17
		currentValue %= 256
	}
	return currentValue
}

type lens struct {
	label string
	fl    int
}

func (s *Day15) PrintPart2() {
	hashmap := make(map[int][]lens, 256)

	//tokens := strings.Split("rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7", ",")
	tokens := strings.Split(s.rows[0], ",")
	for _, token := range tokens {

		if token[len(token)-1] == '-' {
			// remove lens if present.
			label := token[:len(token)-1]
			boxId := hash(label)
			lenses, ok := hashmap[boxId]
			if !ok {
				continue
			}
			index, found := findLensByLabel(lenses, label)
			if !found {
				continue
			}
			hashmap[boxId] = removeIndex(lenses, index)

		} else if strings.Contains(token, "=") {
			instructions := strings.Split(token, "=")
			boxId := hash(instructions[0])
			focalLength, _ := strconv.Atoi(instructions[1])

			newLens := lens{
				label: instructions[0],
				fl:    focalLength,
			}

			lenses, ok := hashmap[boxId]
			if !ok {
				// add
				hashmap[boxId] = []lens{newLens}
			} else {
				//append or replace
				index, found := findLensByLabel(lenses, newLens.label)

				if found {
					hashmap[boxId][index] = newLens
				} else {
					hashmap[boxId] = append(hashmap[boxId], newLens)
				}
			}

		}
	}
	fmt.Println(calculateFocussingPower(hashmap))
}

func findLensByLabel(lenses []lens, label string) (int, bool) {
	for i, l := range lenses {
		if l.label == label {
			return i, true
		}
	}
	return -1, false
}

func removeIndex(s []lens, index int) []lens {
	ret := make([]lens, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func calculateFocussingPower(hashmap map[int][]lens) int {
	var result int
	for boxId, lenses := range hashmap {
		for slotId, lens := range lenses {
			val := (boxId + 1) * lens.fl * (slotId + 1)
			result += val
		}
	}
	return result
}
