package main

import (
	"fmt"
	"os"
	"strings"
	"bufio"
)

type boardingGroup struct {
	ansrs []string
}

func (grp boardingGroup) UniqAppend(ansrs []string) boardingGroup {
	for _, ans := range ansrs {
		grp.ansrs = append(grp.ansrs, ans)
	}
	grp.ansrs = uniq(grp.ansrs)
	return grp
}

func uniq(ansrs []string) []string {
	m := make (map[string]int)
	var u []string
	for _, n := range ansrs {
		if (m[n] == 0) {
			u = append(u, n)
		}
		m[n] += 1
	}
	return u
}

func sumAnswers(path string) int {
	file, _ := os.Open(path)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var grp boardingGroup
	ansSum := 0
	for scanner.Scan() {
		line := scanner.Text()
		if !(line == "") {
			lineAnsrs := strings.Split(line, "")
			grp = grp.UniqAppend(lineAnsrs)
		} else {
			fmt.Println(grp, "has", len(grp.ansrs), "answers")
			ansSum += len(grp.ansrs)
			grp = boardingGroup{}
		}
	}
	return ansSum
}

func main() {
	// slice := []int{1,2,3,4,5,5,5,6,7,4}
	path := "6.txt"
	fmt.Println(sumAnswers(path))
	// slice := []string{"a", "a", "b", "c", "d", "d", "d"}
	// fmt.Println(uniq(slice))
}
