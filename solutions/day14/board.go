package day14

import (
	"fmt"
	"strings"
)

type board struct {
	moveable   item
	immovable  item
	emptySpace item
	data       [][]item
}

type item rune

// insertAt inserts specified item at given coordinates.
func (b *board) insertAt(y, x int, value item) {
	b.data[y][x] = value
}

// moveNorth attempts to move an item at specified coordinates north. Returns
// true if moved, otherwise false.
func (b *board) moveNorth(y, x int) bool {
	if b.data[y][x] != b.moveable {
		return false
	}
	if y-1 < 0 {
		return false
	}

	itemNorth := b.data[y-1][x]
	if itemNorth != b.emptySpace {
		return false
	}

	b.data[y][x], b.data[y-1][x] = b.data[y-1][x], b.data[y][x]
	return true
}

// moveWest attempts to move an item at specified coordinates west. Returns
// true if moved, otherwise false.
func (b *board) moveWest(y, x int) bool {
	if b.data[y][x] != b.moveable {
		return false
	}
	if x-1 < 0 {
		return false
	}

	itemWest := b.data[y][x-1]
	if itemWest != b.emptySpace {
		return false
	}

	b.data[y][x], b.data[y][x-1] = b.data[y][x-1], b.data[y][x]
	return true
}

// moveSouth attempts to move an item at specified coordinates south. Returns
// true if moved, otherwise false.
func (b *board) moveSouth(y, x int) bool {
	if b.data[y][x] != b.moveable {
		return false
	}
	if y+1 >= len(b.data) {
		return false
	}

	itemSouth := b.data[y+1][x]
	if itemSouth != b.emptySpace {
		return false
	}

	b.data[y][x], b.data[y+1][x] = b.data[y+1][x], b.data[y][x]
	return true
}

// moveEast attempts to move an item at specified coordinates east. Returns
// true if moved, otherwise false.
func (b *board) moveEast(y, x int) bool {
	if b.data[y][x] != b.moveable {
		return false
	}
	if x+1 >= len(b.data[y]) {
		return false
	}

	itemEast := b.data[y][x+1]
	if itemEast != b.emptySpace {
		return false
	}

	b.data[y][x], b.data[y][x+1] = b.data[y][x+1], b.data[y][x]
	return true
}

func (b *board) String() string {
	rows := make([]string, len(b.data))
	for y, row := range b.data {
		rows[y] = ""
		for _, v := range row {
			rows[y] = fmt.Sprintf("%s,'%s'", rows[y], string(v))
		}
	}
	return strings.Join(rows, "\n")
}

func newBoard(input []string) *board {
	result := board{
		moveable:   item('O'),
		immovable:  item('#'),
		emptySpace: item('.'),
		data:       make([][]item, len(input)),
	}
	for y, row := range input {
		result.data[y] = make([]item, len(row))
		for x, value := range row {
			result.insertAt(y, x, item(value))
		}
	}

	return &result
}

func (b *board) TiltBoardNorth() {
	b.TiltBoard(b.moveNorth)
}

func (b *board) SpinCycle(cycles int) {
	for i := 0; i < cycles; i++ {

		b.TiltBoard(b.moveNorth)
		b.TiltBoard(b.moveWest)
		b.TiltBoard(b.moveSouth)
		b.TiltBoard(b.moveEast)
	}
}

func (b *board) TiltBoard(itemDirectionalMoveFunc func(int, int) bool) {
	movementsPerformed := true
	for movementsPerformed {

		movementsPerformed = false
		for y, row := range b.data {
			for x, _ := range row {
				moved := itemDirectionalMoveFunc(y, x)

				if !movementsPerformed && moved {
					movementsPerformed = moved
				}
			}
		}
	}
}
