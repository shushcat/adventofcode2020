package main

import (
	"fmt"
	"strings"
	"os"
	"bufio"
	"sort"
)

type seat struct {
	row, col int
}

func (s seat) Id() int {
	seatId := ((s.row * 8) + s.col)
	return seatId
}

func spec2seat(spec string) seat {
	specAry := strings.Split(spec, "")
	s := seat{127, 7}
	subtrahend := (s.row / 2) + 1
	for _, r := range specAry[:7] {
		if r == "F" {
			s.row = s.row - subtrahend
		}
		subtrahend = subtrahend / 2
	}
	subtrahend = (s.col / 2) + 1
	for _, c := range specAry[7:] {
		if c == "L" {
			s.col = s.col - subtrahend
		}
		subtrahend = subtrahend / 2
	}
	return s
}

func collectSeatIds(path string) []int {
	file, _ := os.Open(path)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var ids []int
	for scanner.Scan() {
		spec := scanner.Text()
		s := spec2seat(spec)
		ids = append(ids, s.Id())
	}
	sort.Ints(ids)
	return ids
}

func findSeatingGap(path string) int {
	ids := collectSeatIds(path)
	gap := 0
	for i, _ := range ids[:(len(ids)-1)] {
		if (ids[i+1] - ids[i]) > 1 {
			fmt.Println("There's a gap between", ids[i], "and", ids[i+1])
			gap = ids[i] + 1
		}
	}
	return gap
}

func main() {
	path := "5.txt"
	fmt.Println(findSeatingGap(path))
}
