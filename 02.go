package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	path := "2.txt"
	validCount := 0
	for _, line := range lineAry(path) {
		if validateLine(line) {
			validCount += 1
		}
	}
	fmt.Println(validCount)
}

func validateLine(line string) bool {
	occurRule, occurStr, password := parseLine(line)
	letterCount := 0
	for _, letter := range strings.Split(password, "") {
		if letter == occurStr {
			letterCount += 1
		}
	}
	if letterCount >= occurRule[0] && letterCount <= occurRule[1] {
		return true
	} else {
		return false
	}
}

func parseLine(line string) (occurRule [2]int, occurStr string, password string) {
	fields := strings.Split(line, " ")
	occurRule[0], _ = strconv.Atoi(strings.Split(fields[0], "-")[0])
	occurRule[1], _ = strconv.Atoi(strings.Split(fields[0], "-")[1])
	occurStr = strings.Split(fields[1], "")[0]
	password = fields[2]
	return occurRule, occurStr, password
}

func lineAry(path string) []string {
	file, _ := os.Open(path)
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
