package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func parseData(path string) []int {
	file, _ := os.Open(path)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	data := []int{}
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		data = append(data, num)
	}
	return data
}

func validateNums(preambleLength int, data []int) (int, bool) {
	if preambleLength > len(data) {
		return 0, false
	}
	for i, next := range data[preambleLength:] {
		preSum := false
		for j, v1 := range data[i:i+preambleLength-1] {
			for _, v2 := range data[j+i:i+preambleLength] {
				if (v1+v2 == next) {
					preSum = true
				}
			}
		}
		if preSum == false {
			return next, false
		}
	}
	return 0, true
}

func sumSlice(s []int) int {
	sum := 0
	for _, n := range s {
		sum += n
	}
	return sum
}

func min(a []int) int {
	m := a[0]
	for _, v := range a {
		if v < m {
			m = v
		}
	}
	return m
}

func max(a []int) int {
	m := 0
	for _, v := range a {
		if v > m {
			m = v
		}
	}
	return m
}

func sumContiguousSumMinMax(target int, data []int) int {
	for i, j := 1, 0; i < len(data); i++ {
		slsm := sumSlice(data[j:i])
		if slsm > target {
			j += 1
			i = j + 1
		} else if slsm == target {
			return min(data[j:i]) + max(data[j:i])
		}
	}
	return 0
}

func main() {
	path := "9.txt"
	data := parseData(path)
	invalid, _ := validateNums(25, data)
	fmt.Println("Part 1:", invalid)
	fmt.Println("Part 2:", sumContiguousSumMinMax(invalid, data))
}
