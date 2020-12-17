package main

import (
	"fmt"
)

type turns struct {
	nth  int
	prev int
	nums map[int]int
}

func (ts turns) GenNext() turns {
	ts.nth += 1
	prev := ts.prev
	if _, ok := ts.nums[prev]; ok {
		if ts.nums[prev] != ts.nth-1 {
			ts.prev = (ts.nth - 1) - ts.nums[prev]
			ts.nums[prev] = ts.nth - 1
		} else if ts.nums[prev] == ts.nth-1 {
			ts.prev = 0
		}
	} else {
		ts.prev = 0
		ts.nums[prev] = ts.nth - 1
	}
	return ts
}

func (ts turns) NthSpoken(nth int) int {
	if nth <= ts.nth {
		for k, v := range ts.nums {
			if v == nth {
				return k
			}
		}
	}
	for ts.nth < nth {
		ts = ts.GenNext()
	}
	return ts.prev
}

func newGame(startingNums []int) turns {
	ts := turns{0, 0, map[int]int{}}
	for k, v := range startingNums {
		ts.nth += 1
		ts.nums[v] = k + 1
		ts.prev = v
	}
	return ts
}

func main() {
	ts := turns{0, 0, map[int]int{}}
	ts = newGame([]int{15, 12, 0, 14, 3, 1})
	fmt.Println("Part 1:", ts.NthSpoken(2020))
	ts = newGame([]int{15, 12, 0, 14, 3, 1})
	fmt.Println("Part 2:", ts.NthSpoken(30_000_000))
}
