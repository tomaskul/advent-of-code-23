package util

import (
	"fmt"
	"testing"
)

const SessionCookie = ""

func Test_GetRows_NoEmptyRows(t *testing.T) {
	t.Skip("Specify sessionCookie.")
	testCases := map[string]string{
		"Day 4": "https://adventofcode.com/2023/day/4/input",
		"Day 5": "https://adventofcode.com/2023/day/5/input",
	}

	for name, url := range testCases {
		t.Run(name, func(t *testing.T) {
			rows := GetRows(url, SessionCookie)
			for i, v := range rows {
				if v == "" {
					t.Errorf("Unexpected empty row at index: %d (length: %d)", i, len(rows))
				}
			}
		})
	}
}

func Test_GetRowsV2(t *testing.T) {
	t.Skip("Specify sessionCookie.")

	t.Run("manual test", func(t *testing.T) {
		rows, err := GetCachedRows("https://adventofcode.com/2023/day/3/input", "3", ".txt", "")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		fmt.Printf("rows[:3]: %v\n", rows[:3])
	})
}
