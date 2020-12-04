package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	product := p1("1.txt")
	fmt.Println(product)
}

func p1(path string) int {
	intAry := lines2intAry(path)
	aryLen := len(intAry)
	for i := 0; i < aryLen; i++ {
		for j := i+1; j <  aryLen; j++ {
			for k := j+1; k <  aryLen; k++ {
				sum := intAry[i]+intAry[j]+intAry[k]
				if sum == 2020 {
					fmt.Println(intAry[i], "+", intAry[j], "+", intAry[k])
					return intAry[i]*intAry[j]*intAry[k]
				}
			}
		}
	}
	return 0
}

func lines2intAry(path string) []int {
	file, _ := os.Open(path)
	nextInt := 0
	defer file.Close()
	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		nextInt, _ = strconv.Atoi(scanner.Text())
		lines = append(lines, nextInt)
	}
	return lines
}
