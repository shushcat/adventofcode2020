// Following a slope of right-3, down-1 and starting in the upper left, how many trees would one encounter?

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	path := "3.txt"
	fmt.Println(countTrees(path))
}

func countTrees(path string) int {
	file, _ := os.Open(path)
	defer file.Close()
	var line []string
	pos, trees := 0, 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = strings.Split(scanner.Text(), "")
		if line[pos] == "#" {
			trees += 1
			fmt.Println(line, "has tree at", pos)
		} else {
			fmt.Println(line, "is treeless at", pos)
		}
		pos = (pos + 3) % 31
	}
	return trees
}

// Read the file
// each line is in an array, split
// add 3 to current line index, then modulo line length
// check if new index if "#"; add to count if so
// go to the next line.
