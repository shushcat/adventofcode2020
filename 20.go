package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type tile struct {
	x, y int
	id   int
	pat  [][]string
}
type stack map[int]tile

func (t tile) Corner(ts stack) bool {
	// compare 1951 and 2311 at first
	fmt.Println("entered")
	printTile(t)
	matches := 0
	// t2 := ts[2311]
	for _, t2 := range ts {
		if t3, ok, side := fitTile(t, t2); ok {
			fmt.Println(t3.id, "matches on side", side)
			matches += 1
		}
	}
	fmt.Println(matches)
	fmt.Println(len(ts))
	// printTile(t2)
	// }
	// for id, t2 := range ts {
	// }
	// t2 = rotateTileRight(t2)
	// printTile(ts[2311])
	// fmt.Println()
	// printTile(t2)
	// fmt.Println()
	return false
}

func findCorner(ts stack) tile {
	matches := 0
	corner := tile{0, 0, 0, [][]string{}}
	fmt.Println(len(ts))
	for _, t1 := range ts {
		for _, t2 := range ts {
			if _, ok, _ := fitTile(t1, t2); ok {
				matches += 1
			}
		}
		if matches == 2 {
			corner = tile{t1.x, t1.y, t1.id, t1.pat}
			break
		}
	}

	// fmt.Println(matches)
	// fmt.Println(len(ts))
	fmt.Println(corner.id)
	printTile(corner)
	return corner
}

func readTiles(path string) stack {
	b, _ := ioutil.ReadFile(path)
	input := string(b)
	t := tile{0, 0, 0, [][]string{}}
	ts := stack{}
	for _, l := range strings.Split(input, "\n") {
		if strings.HasPrefix(l, "Tile") {
			l = strings.TrimPrefix(l, "Tile ")
			l = strings.TrimSuffix(l, ":")
			t.id, _ = strconv.Atoi(l)
		} else if strings.HasPrefix(l, ".") || strings.HasPrefix(l, "#") {
			t.pat = append(t.pat, strings.Split(l, ""))
		}
		if l == "" {
			ts[t.id] = t
			t = tile{}
		}
	}
	return ts
}

func printTile(t tile) {
	for row := 0; row < len(t.pat); row++ {
		fmt.Println(t.pat[row])
	}
}

func rotateTileRight(t tile) tile {
	t2 := tile{t.x, t.y, t.id, [][]string{}}
	for row := 9; row >= 0; row-- {
		t2.pat = append(t2.pat, []string{})
		for col := 0; col < 10; col++ {
			t2.pat[9-row] = append(t2.pat[9-row], t.pat[9-col][9-row])
		}
	}
	return t2
}

func flipTileRight(t tile) tile {
	t2 := tile{t.x, t.y, t.id, [][]string{}}
	for row := 0; row < 10; row++ {
		t2.pat = append(t2.pat, []string{})
		for col := 9; col >= 0; col-- {
			t2.pat[row] = append(t2.pat[row], t.pat[row][col])
		}
	}
	return t2
}

// Returns a string holding the characters of a given edge, joined in order
// from left to right and top to bottom.
func getEdge(t tile, side int) string {
	var edge []string
	switch side {
	case 0:
		edge = t.pat[0]
	case 1:
		for i := 0; i < 10; i++ {
			edge = append(edge, t.pat[i][9])
		}
	case 2:
		edge = t.pat[9]
	case 3:
		for i := 0; i < 10; i++ {
			edge = append(edge, t.pat[i][0])
		}
	}
	return strings.Join(edge, "")
}

// The `side` corresponds to the edges of t1 and falls in the range (0, 3).
// For a given value of `side`, the return value is the result of an attempted
// match with the opposite side of t2.
func matchEdge(t1 tile, t2 tile, side int) bool {
	var t1Edge, t2Edge string
	match := false
	t1Edge, t2Edge = getEdge(t1, side), getEdge(t2, (side+2)%4)
	match = t1Edge == t2Edge
	return match
}

func fitTile(t1 tile, t2 tile) (tile, bool, int) {
	match := false
	if t1.id == t2.id {
		return t2, false, -1
	}
	for flip := 0; flip < 2; flip++ {
		for rot := 0; rot < 4; rot++ {
			for edge := 0; edge < 4; edge++ {
				match = matchEdge(t1, t2, edge)
				if match {
					switch edge {
					case 0:
						t2.x, t2.y = t1.x, t1.y + 1
					case 1:
						t2.x, t2.y = t1.x + 1, t1.y
					case 2:
						t2.x, t2.y = t1.x, t1.y - 1
					case 3:
						t2.x, t2.y = t1.x - 1, t1.y + 1
					}
					return t2, true, edge
				}
			}
			t2 = rotateTileRight(t2)
		}
		t2 = flipTileRight(t2)
	}
	return t2, false, -1
}

func cornerProduct(ts stack) int {
	prod := 1
	disqualified := map[int]bool{}
	for _, t1 := range ts {
		if !disqualified[t1.id] {
			matches := 0
			for _, t2 := range ts {
				if _, ok, _ := fitTile(t1, t2); ok {
					matches += 1
				}
			}
			if matches == 2 {
				prod = prod * t1.id
			} else {
				disqualified[t1.id] = true
			}
		}
	}
	return prod
}

func main() {
	fmt.Println("======TEST======")
	path := "20_small.txt"
	// path := "20.txt"
	ts := readTiles(path)
	// for _, t := range ts {
		// fmt.Println(t.id, t.x, t.y)
	// }
	findCorner(ts)
	// matched := 0
	// fmt.Println(matched)
	// ts[1951].Corner(ts)
	// ts[2311].Corner(ts)
	// fmt.Println(Part 1: cornerProduct(ts))
	// fitTile(ts[1951], ts[2311])
	// fitTile(ts[1951], ts[2729])
	// ts[1951].Corner(ts)
}
