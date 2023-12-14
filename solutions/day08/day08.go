package day08

import (
	"fmt"

	"github.com/tomaskul/advent-of-code-23/util"
)

type Day08 struct {
	rows []string
}

type node struct {
	left  string
	right string
}

type direction int

const (
	left direction = iota
	right
)

func NewDay08Solution(sessionCookie string) *Day08 {
	rows, _ := util.GetCachedRows("https://adventofcode.com/2023/day/8/input", "8", ".txt", sessionCookie)
	return &Day08{
		rows: rows,
	}
}

func (s *Day08) PrintPart1() {
	fmt.Println(s.calcDistance())
}

func (s *Day08) calcDistance() int {
	directions := parseDirections(s.rows[0])
	nodes := parseNodes(s.rows[2:])

	cur := "AAA"
	distance := 0
	for i := 0; ; i = (i + 1) % len(directions) {
		if cur == "ZZZ" {
			return distance
		}
		dir := directions[i]
		if dir == left {
			cur = nodes[cur].left
		} else {
			cur = nodes[cur].right
		}
		distance++
	}
}

func parseDirections(s string) []direction {
	var instructions []direction
	for _, c := range s {
		switch c {
		case 'L':
			instructions = append(instructions, left)
		case 'R':
			instructions = append(instructions, right)
		}
	}
	return instructions
}

func parseNodes(lines []string) map[string]node {
	m := make(map[string]node, len(lines))
	for _, s := range lines {
		node, name := parseNode(s)
		m[name] = node
	}
	return m
}

func parseNode(s string) (node, string) {
	//TJS = (LFP, HKT)
	//0123456789012345
	name := s[:3]
	return node{
		left:  s[7:10],
		right: s[12:15],
	}, name
}

func (s *Day08) PrintPart2() {
}
