package main

import (
	"fmt"
	"strings"
	"os"
	"bufio"
	"reflect"
)

func seatLayout(path string) [][]string {
	file, _ := os.Open(path)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var layout [][]string
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		layout = append(layout, line)
	}
	return layout
}

func occupied(layout [][]string, row int, seat int) int {
	if layout[row][seat] == "#" {
		return 1
	} else {
		return 0
	}
}

func numSurrounding(layout [][]string, row int, seat int) int {
	surrounding := 0
	if row > 0 { // N
		surrounding += occupied(layout, row-1, seat)
		if seat > 0 { // NW
			surrounding += occupied(layout, row-1, seat-1)
		}
		if seat < len(layout[row]) - 1 { // NE
			surrounding += occupied(layout, row-1, seat+1)
		}
	}
	if row < len(layout) - 1{ // S
		surrounding += occupied(layout, row+1, seat)
		if seat > 0 { // SW
			surrounding += occupied(layout, row+1, seat-1)
		}
		if seat < len(layout[row]) -1 {
			surrounding += occupied(layout, row+1, seat+1)
		}
	}
	if seat > 0{ // W
		surrounding += occupied(layout, row, seat-1)
	}
	if seat < len(layout[row]) - 1{ // E
		surrounding += occupied(layout, row, seat+1)
	}
	return surrounding
}

func peepBeam(layout [][]string, row int, seat int, i int, j int) int {
	for {
		row, seat = row + i, seat + j
		if row < 0 || seat < 0 || row > len(layout) - 1 || seat > len(layout[row]) - 1 {
			return 0
		} else if layout[row][seat] == "#" {
			return 1
		} else if layout[row][seat] == "L" {
			return 0
		}
	}
}

func printLayout(layout [][]string) {
	for _, v := range layout {
		fmt.Println(v)
	}
}

func tickSeat(layout [][]string, row int, seat int, part int) string {
	surrounding := 0
	if part == 1 {
		surrounding = numSurrounding(layout, row, seat)
	} else if part == 2 {
		surrounding = numPeepable(layout, row, seat)
	}
	if layout[row][seat] == "L" {
		if surrounding == 0 {
			return "#"
		} else {
			return "L"
		}
	} else if layout[row][seat] == "#" {
		if (surrounding > 3 && part == 1) || (surrounding > 4 && part == 2) {
			return "L"
		} else {
			return "#"
		}
	}
	return "."
}

func tickUniverse(layout [][]string, part int) [][]string {
	var tick [][]string
	for row := 0; row < len(layout); row++ {
		tick = append(tick, []string{})
		for seat := 0; seat < len(layout[row]); seat++ {
			tick[row] = append(tick[row], tickSeat(layout, row, seat, part))
		}
	}
	return tick
}

func runDownUniverse(layout [][]string, part int) ([][]string, int) {
	tick := layout
	lastTick := tick
	entropyDistance := 0
	for i := 0; ; i++ {
		entropyDistance += 1
		tick = tickUniverse(tick, part)
		if reflect.DeepEqual(tick, lastTick) {
			break
		}
		lastTick = tick
	}
	return tick, entropyDistance - 1
}

func occupiedSeats(layout [][]string) int {
	occupied := 0
	for row := 0; row < len(layout); row++ {
		for seat := 0; seat < len(layout[row]); seat++ {
			if layout[row][seat] == "#" {
				occupied += 1
			}
		}
	}
	return occupied
}

func numPeepable(layout [][]string, row int, seat int) int {
	peeped := 0
	for i := -1; i <= 1; i++ {
		for j:= -1; j <= 1; j++ {
			if !(j == 0 && i == 0) {
				peeped = peeped + peepBeam(layout, row, seat, i, j)
			}
		}
	}
	return peeped
}

func main() {
	path := "11.txt"
	layout := seatLayout(path)
	surround, _ := runDownUniverse(layout, 1)
	fmt.Println("Part 1:", occupiedSeats(surround))
	peeped, _ := runDownUniverse(layout, 2)
	fmt.Println("Part 2:", occupiedSeats(peeped))
}
