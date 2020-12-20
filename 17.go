package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type coord [4]int
type cubes map[coord]bool

func (c coord) AdjacentCubes() cubes {
	offs := offsets(c)
	step := pow3(len(c))
	adjCubes := cubes{}
	for n := 0; n < step; n++ {
		off := []int{}
		adj := coord{}
		for d := n; d < len(offs); d = d + step {
			off = append(off, offs[d])
		}
		for d := 0; d < len(c); d++ {
			adj[d] = c[d] + off[d]
		}
		adjCubes[adj] = false
	}
	delete(adjCubes, c)
	return adjCubes
}

func (c coord) LiveAdjacent(cbs cubes) cubes {
	liveAdj := cubes{}
	adj := c.AdjacentCubes()
	for c, _ := range adj {
		if cbs[c] == true {
			liveAdj[c] = true
		}
	}
	return liveAdj
}

func (c coord) TickCoord(cbs cubes) bool {
	numLiveAdj := len(c.LiveAdjacent(cbs))
	if (cbs[c] == true) && !(numLiveAdj == 2 || numLiveAdj == 3) {
		return false
	} else if (cbs[c] == false) && (numLiveAdj == 3) {
		return true
	}
	if cbs[c] == true {
		return true
	} else {
		return false
	}
}

func pow3(n int) int {
	pow := 1
	for i := 0; i < n; i++ {
		pow = pow * 3
	}
	return pow
}

func offsets(c coord) []int {
	dim := len(c)
	off := [3]int{-1, 0, 1}
	adj := pow3(dim)
	offsets := []int{}
	cycle := adj / 3
	d := 0
	for d < dim {
		for o := 0; o < 3; o++ {
			for n := 0; n < cycle; n++ {
				offsets = append(offsets, off[o])
			}
		}
		if len(offsets)%adj == 0 {
			cycle = cycle / 3
			d += 1
		}
	}
	return offsets
}

func numActiveCubes(cbs cubes) int {
	n := 0
	for _, c := range cbs {
		if c == true {
			n += 1
		}
	}
	return n
}

func parseInput(path string) cubes {
	file, _ := os.Open(path)
	defer file.Close()
	cbs := cubes{}
	scanner := bufio.NewScanner(file)
	y := 0
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		for x := 0; x < len(line); x++ {
			if line[x] == "#" {
				cbs[coord{y, x}] = true
			}
		}
		y++
	}
	return cbs
}

func tickCubes(cbs cubes, ticks int) cubes {
	tCbs := cubes{}
	for tick := 0; tick < ticks; tick++ {
		for c, _ := range cbs {
			adj := c.AdjacentCubes()
			tCbs[c] = c.TickCoord(cbs)
			for a, _ := range adj {
				tCbs[a] = a.TickCoord(cbs)
			}
		}
		for c, _ := range tCbs {
			if tCbs[c] == false {
				delete(tCbs, c)
			}
		}
		cbs = cubes{}
		for k, v := range tCbs {
			cbs[k] = v
		}
	}
	return tCbs
}

func main() {
	// c0 := coord{3, 4}
	// c0 := coord{3, 4, 5}
	// c0.LiveAdjacent()
	// fmt.Println(len(c0.AdjacentCubes()))
	// c0 := coord{1, 2, 1}
	// c0.LiveAdjacent(init)
	// fmt.Println(c1.TickCoord(init))
	// path := "17_small.txt"
	path := "17.txt"
	cbs := parseInput(path)
	t1 := tickCubes(cbs, 6)
	fmt.Println(t1, len(t1))
	// c1 := coord{2, 2, 0, 1}
	// fmt.Println(len(c1.AdjacentCubes()))
	// fmt.Println(cbs)
}
