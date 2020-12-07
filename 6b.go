package main

import (
	"fmt"
	"os"
	"strings"
	"bufio"
)

type boardingGroup struct {
	ppl int
	ans map[string]int
}

func (grp boardingGroup) AppendLine(ans []string) boardingGroup {
	grp.ppl += 1
	for _, a := range ans {
		grp.ans[a] = grp.ans[a] + 1
	}
	return grp
}

func (grp boardingGroup) UnanimousAns() []string {
	uans := []string{}
	for a, _ := range grp.ans {
		if grp.ans[a] == grp.ppl {
			uans = append(uans, a)
		}
	}
	return uans
}

func sumUnanimousAnswers(path string) int {
	file, _ := os.Open(path)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	grp := boardingGroup{0, make(map[string]int)}
	unanimousSum := 0
	for scanner.Scan() {
		line := scanner.Text()
		if !(line == "") {
			lineAnsrs := strings.Split(line, "")
			grp = grp.AppendLine(lineAnsrs)
		} else {
			unanimousSum += len(grp.UnanimousAns())
			fmt.Println(grp, "has", len(grp.UnanimousAns()), "unanimous answers")
			grp = boardingGroup{0, make(map[string]int)}
		}
	}
	return unanimousSum
}

func main() {
	path := "6.txt"
	fmt.Println(sumUnanimousAnswers(path))
}
