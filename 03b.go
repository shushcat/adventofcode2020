// Following a slope of right-3, down-1 and starting in the upper left, how many trees would one encounter?

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// forest := "3_small.txt"
	forest := "3.txt"
	slopes := [5][2]int{{1,1},{3,1},{5,1},{7,1},{1,2}}
	prod, trees := 1, 0
	for _, slope := range slopes {
		trees = countTrees(forest, slope)
		prod = prod * trees
		fmt.Println("there are", trees, "trees on", slope)
	}
	fmt.Println(prod)
}

func countTrees(forest string, slope [2]int) int {
	file, _ := os.Open(forest)
	defer file.Close()
	var line []string
	linr, pos, trees := 0, 0, 0
	scanner := bufio.NewScanner(file)
	fmt.Println("Checking slope:", slope)
	for scanner.Scan() {
		line = strings.Split(scanner.Text(), "")
		if linr % slope[1] == 0 {
			if line[pos] == "#" {
				trees += 1
				fmt.Println(line, "has tree at", pos)
			} else {
				fmt.Println(line, "is treeless at", pos)
			}
			pos = (pos + slope[0]) % len(line)
		} else {
			fmt.Println(line, "was skipped")
		}
		linr += 1
	}
	return trees
}
