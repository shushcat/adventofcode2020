package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type rule struct {
	name string
	ranges []int
}

type ticket struct {
	vals  []int
	valid int
}

func parseRule(line string) rule {
	r := rule{}
	f := strings.Split(line, ": ")
	f = append(f[:1], strings.Split(strings.Join(f[1:], ""), " or ")...)
	r.name = f[0]
	for _, v := range f[1:] {
		rStrs := strings.Split(v, "-")
		rBeg, _ := strconv.Atoi(rStrs[0])
		rEnd, _ := strconv.Atoi(rStrs[1])
		r.ranges = append(r.ranges, []int{rBeg, rEnd}...)
	}
	return r
}

func parseTicket(line string) ticket {
	t := ticket{[]int{}, -1}
	f := strings.Split(line, ",")
	for i := 0; i < len(f); i++ {
		valInt, _ := strconv.Atoi(f[i])
		t.vals = append(t.vals, valInt)
	}
	return t
}

func mergeRanges(rs []rule) []bool {
	merged := []bool{}
	for _, rule := range rs {
		// Make sure `merged' is big enough
		for len(merged) <= rule.ranges[len(rule.ranges)-1] {
			merged = append(merged, false)
		}
		for i := 0; i < len(rule.ranges) - 1; i += 2 {
			l, h := rule.ranges[i], rule.ranges[i+1]
			for j := l; j <= h; j++ {
				merged[j] = true
			}
		}
	}
	return merged
}

func validateTicket(t ticket, rs []rule) ticket {
	mrs := mergeRanges(rs)
	for _, n := range t.vals {
		if n > len(mrs) {
			t.valid = n
			return t
		} else if mrs[n] != true {
			t.valid = n
			return t
		}
	}
	t.valid = 0
	return t
}

func parseInput(path string) ([]rule, []ticket) {
	file, _ := os.Open(path)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	parsing := "rules"
	rules := []rule{}
	tickets := []ticket{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "your ticket:" || line == "nearby tickets:" {
			parsing = "tickets"
		} else if line != "" {
			switch parsing {
			case "rules":
				rules = append(rules, parseRule(line))
			case "tickets":
				tickets = append(tickets, parseTicket(line))
			}
		}
	}
	return rules, tickets
}

func ticketScanError(tickets []ticket, rules []rule) int {
	sumErr := 0
	for _, t := range tickets[1:] {
		t = validateTicket(t, rules)
		sumErr += t.valid
	}
	return sumErr
}

func main() {
	// path := "16_small.txt"	// -> 71
	path := "16.txt"	// -> 19240
	rules, tickets := parseInput(path)
	fmt.Println("Part 1:", ticketScanError(tickets, rules))
	// tic := validateTicket(tickets[4], rules)
	// fmt.Println(tic)
}
