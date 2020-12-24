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
	pat  [10][10]string
}

type stack map[int]tile

func readTiles(path string) stack {
	b, _ := ioutil.ReadFile(path)
	input := string(b)
	t := tile{}
	ts := stack{}
	i := 0
	for _, l := range strings.Split(input, "\n") {
		if strings.HasPrefix(l, "Tile") {
			l = strings.TrimPrefix(l, "Tile ")
			l = strings.TrimSuffix(l, ":")
			t.id, _ = strconv.Atoi(l)
			i = 0
		} else if strings.HasPrefix(l, ".") || strings.HasPrefix(l, "#") {
			cs := strings.Split(l, "")
			for j := 0; j < 10; j++ {
				t.pat[i][j] = cs[j]
			}
			i++
		}
		if l == "" {
			ts[t.id] = t
			t = tile{}
		}
	}
	return ts
}

func printAryAry(a [][]string) {
	for row := 0; row < len(a); row++ {
		fmt.Println(a[row])
	}
}

func rotateTileRight(t tile) tile {
	t2 := tile{t.x, t.y, t.id, [10][10]string{}}
	for row := 9; row >= 0; row-- {
		for col := 0; col < 10; col++ {
			t2.pat[9-row][col] = t.pat[9-col][9-row]
		}
	}
	return t2
}

func rotateImgRight(img [][]string) [][]string {
	img2 := [][]string{}
	imgSideLen := len(img)
	img2 = initializeImage(img2, imgSideLen)
	for row := imgSideLen - 1; row >= 0; row-- {
		for col := 0; col < imgSideLen; col++ {
			img2[(imgSideLen-1)-row][col] = img[(imgSideLen-1)-col][(imgSideLen-1)-row]
		}
	}
	return img2
}

func flipTileRight(t tile) tile {
	t2 := tile{t.x, t.y, t.id, [10][10]string{}}
	for row := 0; row < 10; row++ {
		for col := 9; col >= 0; col-- {
			t2.pat[row][9-col] = t.pat[row][col]
		}
	}
	return t2
}

func flipImgRight(img [][]string) [][]string {
	img2 := [][]string{}
	imgSideLen := len(img)
	img2 = initializeImage(img, imgSideLen)
	for row := 0; row < imgSideLen; row++ {
		for col := 0; col < imgSideLen; col ++ {
			img2[row][col] = img[row][(imgSideLen-1)-col]
		}
	}
	return img2
}

// Returns a string holding the characters of a given edge, joined in order
// from left to right and top to bottom.
func getSide(t tile, side int) [10]string {
	var sideChars [10]string
	switch side {
	case 0:
		sideChars = t.pat[0]
	case 1:
		for i := 0; i < 10; i++ {
			sideChars[i] = t.pat[i][9]
		}
	case 2:
		sideChars = t.pat[9]
	case 3:
		for i := 0; i < 10; i++ {
			sideChars[i] = t.pat[i][0]
		}
	}
	return sideChars
}

// The `side` corresponds to the edges of t1 and falls in the range (0, 3).
// For a given value of `side`, the return value is the result of an attempted
// match with the opposite side of t2.
func matchSide(t1 tile, t2 tile, side int) bool {
	var t1Side, t2Side [10]string
	match := false
	t1Side, t2Side = getSide(t1, side), getSide(t2, (side+2)%4)
	match = t1Side == t2Side
	return match
}

func fitTile(t1 tile, t2 tile) (tile, bool) {
	match := false
	if t1.id == t2.id {
		return t2, false
	}
	for flip := 0; flip < 2; flip++ {
		for rot := 0; rot < 4; rot++ {
			for side := 0; side < 4; side++ {
				match = matchSide(t1, t2, side)
				if match {
					switch side {
					case 0:
						t2.x, t2.y = t1.x, t1.y+1
					case 1:
						t2.x, t2.y = t1.x+1, t1.y
					case 2:
						t2.x, t2.y = t1.x, t1.y-1
					case 3:
						t2.x, t2.y = t1.x-1, t1.y
					}
					return t2, true
				}
			}
			t2 = rotateTileRight(t2)
		}
		t2 = flipTileRight(t2)
	}
	return t2, false
}

func cornerProduct(layout [][]int) int {
	nw := layout[0][0]
	ne := layout[0][len(layout)-1]
	sw := layout[len(layout)-1][0]
	se := layout[len(layout)-1][len(layout)-1]
	prod := nw * ne * sw * se
	return prod
}

func tileChain(t1 tile, chained stack, ts stack) stack {
	chained[t1.id] = t1
	if len(chained) != len(ts) {
		for _, t2 := range ts {
			if t3, ok := fitTile(t1, t2); ok {
				t1 = t3
				chained[t1.id] = t1
				break
			}
		}
		chained = tileChain(t1, chained, ts)
	}
	return chained
}

func seedTile(ts stack) tile {
	t1 := tile{}
	for _, t2 := range ts {
		t1 = t2
		break
	}
	return t1
}

func normalizeTileCoordinates(ts stack) (stack, int, int) {
	xMin, yMin := 0, 0
	for _, t := range ts {
		if t.x < xMin {
			xMin = t.x
		}
		if t.y < yMin {
			yMin = t.y
		}
	}
	xNorm, yNorm := -xMin, -yMin
	normStack := stack{}
	for k, t := range ts {
		normStack[k] = tile{t.x + xNorm, t.y + yNorm, t.id, t.pat}
	}
	xMax, yMax := 0, 0
	for _, t := range normStack {
		if t.x > xMax {
			xMax = t.x
		}
		if t.y > yMax {
			yMax = t.y
		}
	}
	return normStack, xMax, yMax
}

func layoutTiles(ts stack) [][]int {
	tc := stack{}
	seed := seedTile(ts)
	tc = tileChain(seed, tc, ts)
	normalized, xMax, yMax := normalizeTileCoordinates(tc)
	layout := [][]int{}
	for y := 0; y <= yMax; y++ {
		layout = append(layout, []int{})
		for x := 0; x <= xMax; x++ {
			layout[y] = append(layout[y], 0)
		}
	}
	for _, t := range normalized {
		layout[t.y][t.x] = t.id
	}
	return layout
}

func joinTiles(ts stack, layout [][]int) [][]string {
	image := [][]string{}
	return image
}

func printLayout(layout [][]int) {
	for i := 0; i < len(layout); i++ {
		fmt.Println(layout[i])
	}
}

func shrinkTile(t tile) [][]string {
	t2 := [][]string{}
	for row := 1; row < 9; row++ {
		t2 = append(t2, []string{})
		for col := 1; col < 9; col++ {
			t2[row-1] = append(t2[row-1], t.pat[row][col])
		}
	}
	return t2
}

// Accepts an image, a tile shrunk to 8 by 8, and the row and column of that
// tile in the tile layout.
func insertTile(img [][]string, st [][]string, lRow int, lCol int) [][]string {
	for row := 0; row < 8; row++ {
		for col := 0; col < 8; col++ {
			iRow := row + (lRow * 8)
			iCol := col + (lCol * 8)
			img[iRow][iCol] = st[row][col]
		}
	}
	return img
}

func initializeImage(img [][]string, imgSideLen int) [][]string {
	for row := 0; row < imgSideLen; row++ {
		img = append(img, []string{})
		for col := 0; col < imgSideLen; col++ {
			img[row] = append(img[row], "x")
		}
	}
	return img
}

func assembleImage(layout [][]int, ts stack) [][]string {
	layoutSideLen := len(layout)
	imgSideLen := layoutSideLen * 8
	img := [][]string{}
	// Initialize image.
	img = initializeImage(img, imgSideLen)
	// Adjoin shrunken tiles in the image.
	for row := 0; row < layoutSideLen; row++ {
		for col := 0; col < layoutSideLen; col++ {
			tileID := layout[row][col]
			insertTile(img, shrinkTile(ts[tileID]), row, col)
		}
	}
	return img
}

func patAtPoint(img [][]string, row int, col int) bool {
	fmt.Println(row, col)
	pat := [15]string{img[row][col], img[row+1][col+1], img[row+1][col+4], img[row][col+5], img[row][col+6], img[row+1][col+7], img[row+1][col+10], img[row][col+11], img[row][col+12], img[row+1][col+13], img[row+1][col+16], img[row][col+17], img[row-1][col+18], img[row][col+18], img[row][col+19]}
	match := [15]string{"#", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#"} == pat
	return match
}

func numPats(img [][]string) int {
	num := 0
	for row := 1; row < len(img)-2; row++ {
		for col := 0; col < len(img) - 21; col++ {
			patAtPoint(img, row, col)
		}
	}
	return num
}

func wherePatAt(img [][]string) int {
	num := 0
	for flip := 0; flip < 2; flip++ {
		for rot := 0; rot < 4; rot++ {
			// num = numPats(img)
			// if num > 0 {
				// return num
			// }
			fmt.Println("rotate")
			img = rotateImgRight(img)
		}
		fmt.Println("flip")
		img = flipImgRight(img)
	}
	return num
}

func main() {
	fmt.Println("============TEST============")
	path := "20_small.txt"
	// path := "20.txt"
	ts := readTiles(path)
	// fmt.Println(ts)
	layout := layoutTiles(ts)
	img := assembleImage(layout, ts)
	wherePatAt(img)
	// fmt.Println(layout)		//0
	// fmt.Println(numPats(img))
	// img = rotateImgRight(img)	//1
	// fmt.Println(numPats(img))
	// img = rotateImgRight(img)	//2
	// fmt.Println(numPats(img))
	// img = rotateImgRight(img)	//3
	// fmt.Println(numPats(img))
	// img = rotateImgRight(img)	//0
	// fmt.Println("flip")
	// img = flipImgRight(img)
	// fmt.Println(numPats(img))
	// fmt.Println("Part 1:", cornerProduct(layout))
}
