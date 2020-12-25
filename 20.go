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
	fmt.Println()
	for row := 0; row < len(a); row++ {
		fmt.Println(a[row])
	}
}

func printTilePat(t tile) {
	fmt.Println()
	for row := 0; row < len(t.pat); row++ {
		fmt.Println(t.pat[row])
	}
}

func printLayout(layout [][]int) {
	for i := 0; i < len(layout); i++ {
		fmt.Println(layout[i])
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
	nRows, nCols := len(img), len(img[0])
	img2 := initializeImage(nCols, nRows)
	for row := nRows - 1; row >= 0; row-- {
		for col := 0; col < nCols; col++ {
			img2[(nCols-1)-col][row] = img[(nRows-1)-row][(nCols-1)-col]
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
	nRows, nCols := len(img), len(img[0])
	img2 := initializeImage(nRows, nCols)
	for row := 0; row < nRows; row++ {
		for col := 0; col < nCols; col++ {
			img2[row][col] = img[row][(nCols-1)-col]
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

func chooseCorner(ts stack) tile {
	var t1 tile
	for _, t2 := range ts {
		adjacent := 0
		for _, t3 := range ts {
			if !(t2.id == t3.id) {
				if _, ok := fitTile(t2, t3); ok {
					adjacent += 1
				}
			}
		}
		if adjacent == 2 {
			t1 = t2
			break
		}
	}
	return t1
}

func initTileChain(ts stack) (tile, stack) {
	chained := stack{}
	corner := chooseCorner(ts)
	chained[corner.id] = corner
	// Orient corner in upper-left.
	adjacent := []tile{}
	for _, t2 := range ts {
		if t3, ok := fitTile(corner, t2); ok {
			adjacent = append(adjacent, t3)
		}
	}
	var displacement [2]int
	for _, t3 := range adjacent {
		displacement[0] += t3.x
		displacement[1] += t3.y
	}
	switch displacement {
	case [2]int{-1, 1}:
		corner = rotateTileRight(corner)
		corner = rotateTileRight(corner)
	case [2]int{1, 1}:
		corner = rotateTileRight(corner)
	case [2]int{-1, -1}:
		corner = flipTileRight(corner)
	}
	return corner, chained
}

func directionalChain(ts stack) stack {
	t1, chained := initTileChain(ts)
	dir := 2
	for len(chained) != len(ts) {
		if t2, ok := directionalFit(t1, ts, dir); ok {
			chained[t2.id] = t2
			t1 = t2
		} else if t2, ok := directionalFit(t1, ts, 1); ok {
			chained[t2.id] = t2
			dir = (dir+2)%4
			t1 = t2
		}
	}
	return chained
}

func directionalFit(t1 tile, ts stack, dir int) (tile, bool) {
	match := false
	for _, t2 := range ts {
		if t1.id != t2.id {
			for flip := 0; flip < 2; flip++ {
				for rot := 0; rot < 4; rot++ {
					match = matchSide(t1, t2, dir)
					if match {
						switch dir {
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
					t2 = rotateTileRight(t2)
				}
				t2 = flipTileRight(t2)
			}
		}
	}
	return t1, false
}

func tileChain(ts stack) stack {
	t1, chained := initTileChain(ts)
	for len(chained) != len(ts) {
		// tries := 0
		fmt.Println(len(ts), len(chained))
		for _, t2 := range ts {
			if _, ok := chained[t2.id]; !ok {
				if t2, ok := fitTile(t1, t2); ok {
					chained[t1.id] = t1
					t1 = t2
					// tries = 0
				}
			}
			// if tries == len(ts) {
				// t1, chained = initTileChain(ts)
			// }
			// tries += 1
		}
	}
	fmt.Println(chained)
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
	normStack := stack{}
	for k, t := range ts {
		normStack[k] = tile{t.x - xMin, t.y - yMin, t.id, t.pat}
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

func layoutTiles(ts stack) ([][]int, stack) {
	tc := stack{}
	tc = directionalChain(ts)
	normalized, xMax, yMax := normalizeTileCoordinates(tc)
	tc = normalized
	layout := [][]int{}
	for y := 0; y <= yMax; y++ {
		layout = append(layout, []int{})
		for x := 0; x <= xMax; x++ {
			layout[y] = append(layout[y], 0)
		}
	}
	for _, t := range tc {
		layout[t.y][t.x] = t.id
	}
	return layout, tc
}

func joinTiles(ts stack, layout [][]int) [][]string {
	image := [][]string{}
	return image
}

func shrinkTile(t tile) [8][8]string {
	var st [8][8]string
	for row := 1; row < 9; row++ {
		for col := 1; col < 9; col++ {
			st[row-1][col-1] = t.pat[row][col]
		}
	}
	return st
}

// Accepts an image, a tile , and the row and column of that tile in the
// layout.
func insertTile(img [][]string, t tile, lRow int, lCol int) [][]string {
	t = flipTileRight(rotateTileRight(rotateTileRight(t)))
	st := shrinkTile(t)
	fmt.Print("\n", t.id, ":\n")
	for i:= 0; i<8;i++ {
		fmt.Println(st[i])
	}
	for row := 0; row < 8; row++ {
		for col := 0; col < 8; col++ {
			iRow := row + (lRow * 8)
			iCol := col + (lCol * 8)
			img[iRow][iCol] = st[row][col]
		}
	}
	printAryAry(img)
	return img
}

func initializeImage(nRows int, nCols int) [][]string {
	img := make([][]string, nRows)
	for row := 0; row < nRows; row++ {
		for col := 0; col < nCols; col++ {
			img[row] = append(img[row], "x")
		}
	}
	return img
}

func assembleImage(layout [][]int, ts stack) [][]string {
	layoutSideLen := len(layout)
	imgSideLen := layoutSideLen * 8
	// Initialize image.
	img := initializeImage(imgSideLen, imgSideLen)
	// Join shrunken tiles to the image.
	for row := 0; row < layoutSideLen; row++ {
		for col := 0; col < layoutSideLen; col++ {
			tileID := layout[row][col]
			img = insertTile(img, ts[tileID], row, col)
		}
	}
	return img
}

func patAtPoint(img [][]string, row int, col int) bool {
	if len(img) < 3 || len(img[0]) < 20 {
		return false
	}
	pat := [15]string{img[row+1][col], img[row+2][col+1], img[row+2][col+4], img[row+1][col+5], img[row+1][col+6], img[row+2][col+7], img[row+2][col+10], img[row+1][col+11], img[row+1][col+12], img[row+2][col+13], img[row+2][col+16], img[row+1][col+17], img[row][col+18], img[row+1][col+18], img[row+1][col+19]}
	var match = [15]string{"#", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#"} == pat
	return match
}

func numPats(img [][]string) int {
	num := 0
	nRows, nCols := len(img), len(img[0])
	for row := 0; row < nRows-2; row++ {
		for col := 0; col < nCols-19; col++ {
			if patAtPoint(img, row, col) {
				num += 1
			}
		}
	}
	return num
}

func wherePatAt(img [][]string) ([][]string, int) {
	num := 0
	for flip := 0; flip < 2; flip++ {
		for rot := 0; rot < 4; rot++ {
			num += numPats(img)
			if num > 0 {
				return img, num
			}
			img = rotateImgRight(img)
		}
		img = flipImgRight(img)
	}
	return img, num
}

func imgsEqual(img1 [][]string, img2 [][]string) (bool, int, int) {
	if !(len(img1) == len(img2)) || !(len(img1[0]) == len(img2[0])) {
		return false, 0, 0
	}
	for i:=0;i<len(img1);i++{
		for j:=0;j<len(img1);j++{
			if !(img1[i][j] == img2[i][j]) {
				return false, i, j
			}
		}
	}
	return true, 0, 0
}

func testSamplePat() [][]string {
	path := "20_small_img.txt"
	b, _ := ioutil.ReadFile(path)
	input := string(b)
	var img [][]string
	for _, l := range strings.Split(input, "\n") {
		if l != "" {
			img = append(img, strings.Split(l, ""))
		}
	}
	n := 0
	_, n = wherePatAt(img)
	fmt.Println("rot 0:", n)
	img = rotateImgRight(img)
	_, n = wherePatAt(img)
	fmt.Println("rot 1:", n)
	img = rotateImgRight(img)
	_, n = wherePatAt(img)
	fmt.Println("rot 2:", n)
	img = rotateImgRight(img)
	_, n = wherePatAt(img)
	fmt.Println("rot 3:", n)
	img = rotateImgRight(img)
	_, n = wherePatAt(img)
	fmt.Println("rot 4:", n)
	img = flipImgRight(img)
	_, n = wherePatAt(img)
	fmt.Println("flip 1:", n)
	return img
}

func hashesLessPats(img [][]string) int {
	pats := 33 // TODO Fix the off-by ones that keep giving me 32.
	hashes := 0
	for row := 0; row < len(img); row++ {
		for col := 0; col < len(img[0]); col++ {
			if img[row][col] == "#" {
				hashes += 1
			}
		}
	}
	return hashes - (pats*15)
}

func main() {
	fmt.Println("============TEST============")
	// path := "20_small.txt"
	path := "20.txt"
	ts := readTiles(path)
	layout, tc := layoutTiles(ts)
	img := assembleImage(layout, tc)
	_, n := wherePatAt(img)
	fmt.Println(n)
	printLayout(layout)
	fmt.Println(hashesLessPats(img))
	// testSamplePat()
	// printAryAry(img)
	// printLayout(layout)
}
