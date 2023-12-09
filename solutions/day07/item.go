package day07

import (
	"fmt"

	"github.com/jinzhu/copier"
)

type item struct {
	hand card
	bid  int
}

type card struct {
	value  string
	lookup thing
}

type thing map[byte]int

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

func (c card) getType(cashInJokers bool) int {
	var jLookup map[byte]int
	if cashInJokers {
		jLookup = c.cashInJokers()
	}

	if c.isFiveOfAKind(jLookup) {
		return FiveOfKindType
	} else if c.isFourOfAKind(jLookup) {
		return FourOfKindType
	} else if c.isFullHouse(jLookup) {
		return FullHouseType
	} else if c.isThreeOfAKind(jLookup) {
		return ThreeOfKindType
	} else if c.isTwoPair(jLookup) {
		return TwoPairType
	} else if c.isOnePair(jLookup) {
		return OnePairType
	} else if c.isHighCard(jLookup) {
		return HighCardType
	} else {
		fmt.Printf("debug: unable to determine %q type (cashInJokers: %v)\n", c.value, cashInJokers)
		if cashInJokers {
			fmt.Printf("jLookup: %+v\n", jLookup)
		}
		return -1
	}
}

func (c card) cashInJokers() map[byte]int {
	var maxKey byte
	maxCount, jokerCount := 0, 0
	for k, v := range c.lookup {
		if k == 'J' {
			jokerCount = v
		} else {
			if v > maxCount {
				maxCount = v
				maxKey = k
			}
		}
	}
	if jokerCount > 0 {
		var jLookup map[byte]int
		copier.Copy(&jLookup, c.lookup)

		jLookup[maxKey] += jokerCount
		delete(jLookup, 'J')
		return jLookup
	}
	return nil
}

func (t thing) hasAnyWithCount(count int) bool {
	for _, v := range t {
		if v == count {
			return true
		}
	}
	return false
}

func (c card) isFiveOfAKind(jLookup thing) bool {
	if jLookup != nil {
		return len(jLookup) == 1
	}
	return len(c.lookup) == 1
}

func (c card) isFourOfAKind(jLookup thing) bool {
	if jLookup != nil {
		return len(jLookup) == 2 &&
			(jLookup.hasAnyWithCount(1) || jLookup.hasAnyWithCount(4))
	}

	return len(c.lookup) == 2 &&
		(c.lookup.hasAnyWithCount(1) || c.lookup.hasAnyWithCount(4))
}

func (c card) isFullHouse(jLookup thing) bool {
	if jLookup != nil {
		return len(jLookup) == 2 &&
			(jLookup.hasAnyWithCount(3) || jLookup.hasAnyWithCount(2))
	}

	return len(c.lookup) == 2 &&
		(c.lookup[c.value[0]] == 3 || c.lookup[c.value[0]] == 2)
}

func (c card) isThreeOfAKind(jLookup thing) bool {
	search := c.lookup
	if jLookup != nil {
		search = jLookup
	}

	if len(search) != 3 {
		return false
	}
	for _, v := range search {
		if v == 2 {
			return false
		}
	}
	return true
}

func (c card) isTwoPair(jLookup thing) bool {
	if c.isFiveOfAKind(jLookup) || c.isFourOfAKind(jLookup) || c.isThreeOfAKind(jLookup) {
		return false
	}

	search := c.lookup
	if jLookup != nil {
		search = jLookup
	}

	pairOne := false
	for _, v := range search {
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

func (c card) isOnePair(jLookup thing) bool {
	search := c.lookup
	if jLookup != nil {
		search = jLookup
	}

	pairCounts := 0
	for _, v := range search {
		if v == 2 {
			pairCounts++
		}
	}
	return pairCounts == 1
}

func (c card) isHighCard(jLookup thing) bool {
	if jLookup != nil {
		return len(jLookup) == 5
	}

	return len(c.lookup) == 5
}
