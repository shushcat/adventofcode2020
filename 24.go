/*
Hexagonal tilings.

Given a list of hexagonal tiles, what is the final pattern they will assume
following a series of flips?  Each tile is black on one side and  white on the
other, and is uniquely specified from the reference tile at the center of the
board.  The specification for each tile is a string composed from the
concatenated abbreviations for the directions east, southeast, southwest, west,
northwest, and northeast; respectively: "e", "se", "sw", "w", "nw", "ne".  For
example, "esenee" and "nsweew" are both valid specifications.

TODO Confer with the resource at
https://www.redblobgames.com/grids/hexagons/#coordinates, and review the Reddit
thread for this problem.

*/

package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"flag"
	"os"
	"log"
	"runtime/pprof"
)

type tiles map[[2]int]bool

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func (ts tiles) cleanup() {
	for k, v := range ts {
		if v == false {
			delete(ts, k)
		}
	}
}

func parseLine(line string) [2]int {
	chars := strings.Split(line, "")
	var coord [2]int
	for i := 0; i < len(chars); i++ {
		switch chars[i] {
		case "e":
			coord[0], coord[1] = coord[0]+1, coord[1]+-1
		case "w":
			coord[0], coord[1] = coord[0]+-1, coord[1]+1
		case "n":
			i++
			if chars[i] == "e" {
				coord[0], coord[1] = coord[0]+1, coord[1]+0
			} else {
				coord[0], coord[1] = coord[0]+0, coord[1]+1
			}
		case "s":
			i++
			if chars[i] == "e" {
				coord[0], coord[1] = coord[0]+0, coord[1]+-1
			} else {
				coord[0], coord[1] = coord[0]+-1, coord[1]+0
			}
		}
	}
	return coord
}

func (ts tiles) flipTile(coord [2]int) {
	if ts[coord] == true {
		ts[coord] = false
	} else {
		ts[coord] = true
	}
}

func (ts tiles) adjacent(coord [2]int) [6][2]int {
	dirs := [6][2]int{[2]int{1, -1}, [2]int{-1, 1},
		[2]int{1, 0}, [2]int{0, 1},
		[2]int{0, -1}, [2]int{-1, 0}}
	var adj [6][2]int
	for i, dir := range dirs {
		adj[i] = [2]int{coord[0] + dir[0], coord[1] + dir[1]}
	}
	return adj
}

func (ts tiles) adjacentBlack(coord [2]int) int {
	count := 0
	for _, adj := range ts.adjacent(coord) {
		if ts[adj] {
			count += 1
		}
	}
	return count
}

func (ts tiles) readTiles(path string) {
	bytes, _ := ioutil.ReadFile(path)
	lines := strings.Split(string(bytes), "\n")
	for i := 0; i < len(lines); i++ {
		if lines[i] != "" {
			ts.flipTile(parseLine(lines[i]))
		}
	}
}

func (ts tiles) stepTiles(days int) {
	for {
		flipQueue := tiles{}
		ts.cleanup()
		// Queue tiles to be flipped.
		for t1, _ := range ts {
			adjacentTiles := ts.adjacent(t1)
			adjacentBlack := ts.adjacentBlack(t1)
			if adjacentBlack == 0 || adjacentBlack > 2 {
				flipQueue[t1] = false
			}
			for _, t2 := range adjacentTiles {
				adjacentBlack = ts.adjacentBlack(t2)
				if ts[t2] == false && adjacentBlack == 2 {
					flipQueue[t2] = true
				}
			}
		}
		for k, v := range flipQueue {
			ts[k] = v
		}
		days--
		if days == 0 {
			break
		}
	}
}

func test() {
	m := make(map[[2]int]bool)
	fmt.Println(m)
	m[[2]int{1, 2}] = true
	fmt.Println(m)

}

func (ts tiles) sumBlack() int {
	count := 0
	for _, v := range ts {
		if v == true {
			count += 1
		}
	}
	return count
}

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	// input := "24_small.txt"
	input := "24.txt"

	ts := tiles{}
	ts.readTiles(input)
	fmt.Println("Part 1:", ts.sumBlack())
	ts.stepTiles(100)
	fmt.Println("Part 2:", ts.sumBlack())
}


