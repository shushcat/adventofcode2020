package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type rule struct {
	name   string
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
		for i := 0; i < len(rule.ranges)-1; i += 2 {
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

func parseInput(path string) ([]ticket, []rule) {
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
	return tickets, rules
}

func sumScanError(tickets []ticket, rules []rule) int {
	sumErr := 0
	for _, t := range tickets {
		t = validateTicket(t, rules)
		sumErr += t.valid
	}
	return sumErr
}

func discardInvalidNearby(tickets []ticket, rules []rule) []ticket {
	ts := []ticket{}
	for _, t := range tickets[1:] {
		t = validateTicket(t, rules)
		if t.valid == 0 {
			ts = append(ts, t)
		}
	}
	return ts
}

func valueField(tickets []ticket, rules []rule) [][]int {
	validTickets := discardInvalidNearby(tickets, rules)
	valField := [][]int{}
	for i, r := range rules {
		// Rules and fields.
		valField = append(valField, []int{})
		for field := 0; field < len(rules); field++ {
			valField[i] = append(valField[i], 0)
			validFields := 0
			validRuleRange := mergeRanges([]rule{r})
			for _, t := range validTickets {
				fieldInRange := validRuleRange[t.vals[field]]
				if fieldInRange {
					validFields += 1
				}
			}
			if validFields == len(validTickets) {
				valField[i][field] = 1
			}
		}
	}
	return valField
}

func ticketFieldNames(tickets []ticket, rules []rule) map[int]string {
	// valField ≡ rules x fields
	valField := valueField(tickets, rules)
	fieldMap := map[int]string{}
	for r := 0; r < len(valField); r++ {
		slots := 0
		field := 0
		for f := 0; f < len(valField[r]); f++ {
			if valField[r][f] == 1 {
				slots += 1
				field = f
			}
		}
		if slots == 1 {
			fieldMap[field] = rules[r].name
			for i := 0; i < len(valField); i++ {
				valField[i][field] = 0
			}
			r = -1
		}
	}
	return fieldMap
}

func multDepartureFields(tickets []ticket, rules []rule) int {
	fieldMap := ticketFieldNames(tickets, rules)
	t := tickets[0]
	var depFields []int
	for i, f := range fieldMap {
		if ok, _ := regexp.MatchString("^departure", f); ok {
			depFields = append(depFields, t.vals[i])
		}
	}
	prod := 1
	for _, f := range depFields {
		prod = prod * f
	}
	return prod
}

func main() {
	// path := "16_small.txt"
	// path := "16_small2.txt"
	path := "16.txt"
	tickets, rules := parseInput(path)
	fmt.Println("Part 1:", sumScanError(tickets, rules))
	fmt.Println("Part 2:", multDepartureFields(tickets, rules))
}
