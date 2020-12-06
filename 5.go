package main

import (
	"fmt"
	"strings"
	"os"
	"bufio"
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
		fmt.Print(r)
		if r == "F" {
			s.row = s.row - subtrahend
		}
		subtrahend = subtrahend / 2
		fmt.Print(s.row, "\n")
	}
	subtrahend = (s.col / 2) + 1
	for _, c := range specAry[7:] {
		fmt.Print(c)
		if c == "L" {
			s.col = s.col - subtrahend
		}
		subtrahend = subtrahend / 2
		fmt.Print(s.col, "\n")
	}
	return s
}

func highestSeatId(path string) int {
	file, _ := os.Open(path)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	highestId := 0
	for scanner.Scan() {
		spec := scanner.Text()
		s := spec2seat(spec)
		if s.Id() > highestId {
			highestId = s.Id()
		}

	}
	return highestId
}

func main() {
	path := "5.txt"
	fmt.Println(highestSeatId(path))
	// spec := "FBFBBFFRLR"
	// spec := "BFFFBBFRRR"
	// spec := "FFFBBBFRRR"
	// spec := "BBFFBBFRLL"
	// s := spec2seat(spec)
	// fmt.Println(s, s.Id())
}
