package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// type rule map[string]map[string]int

type rule struct {
	name string
	held map[string]int
}

func parseRule(line string) rule {
	r := rule{"", make(map[string]int)}
	l := strings.Split(line[:len(line)-1], " bags contain ")
	l[1] = strings.ReplaceAll(strings.ReplaceAll(l[1], " bags", ""), " bag", "")
	if l[1] == "no other" {
		r.name = l[0]
		r.held = map[string]int{"none": 0}
		return r
	}
	held := make(map[string]int)
	for _, s := range strings.Split(l[1], ", ") {
		v, _ := strconv.Atoi(s[:1])
		k := s[2:]
		held[k] = v
	}
	r.name = l[0]
	r.held = held
	return r
}

func ruleContains(r rule, bag string) bool {
	if _, ok := r.held[bag]; ok {
		return true
	}
	return false
}

func readRules(path string) []rule {
	file, _ := os.Open(path)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var rs []rule
	for scanner.Scan() {
		rs = append(rs, parseRule(scanner.Text()))
	}
	return rs
}

func delRule(rs []rule, r rule) []rule {
	rs2 := []rule{}
	for _, v := range rs {
		if v.name != r.name {
			rs2 = append(rs2, v)
		}
	}
	return rs2
}

func keyIndex(key string, rs []rule) (int, bool) {
	for i, v := range rs {
		if v.name == key {
			return i, true
		}
	}
	return 0, false
}

func min(ary []int) int {
	m := ary[0]
	for _, n := range ary {
		if n < m {
			m = n
		}
	}
	return m
}

func searchContents(r rule, rs []rule, bag string, depth int, found bool) (int, bool) {
	if !ruleContains(r, bag) {
		rs = delRule(rs, r)
		for k, _ := range r.held {
			branchDepth := depth
			if i, ok := keyIndex(k, rs); ok {
				branchDepth, found = searchContents(rs[i], rs, bag, branchDepth+1, false)
			}
			if found == true {
				return branchDepth, true
			}
		}
	} else {
		found = true
	}
	return depth, found
}

func canContain(rs []rule, bag string) int {
	canContain := 0
	for i, _ := range rs {
		_, contains := searchContents(rs[i], rs, bag, 1, false)
		if contains {
			canContain += 1
		}
	}
	return canContain
}

func main() {
	// path := "7_small.txt"
	path := "7.txt"
	rs := readRules(path)
	fmt.Println(canContain(rs, "shiny gold"))
	// fmt.Println(searchContents(rs[0], rs, "muted yellow", 1, false))
	// fmt.Println(searchContents(rs[0], rs, "vibrant plum", 1, false))
	// fmt.Println(searchContents(rs[0], rs, "shiny gold", 1, false))
	// fmt.Println(searchContents(rs[4], rs, "light red", 1, false))
	// fmt.Println(containmentDistance(rs["muted yellow"], rs, "faded blue"))
	// for _, r := range rs {
	// fmt.Println("Does", r, "contain 'shiny gold'?", ruleContains(r, "shiny gold"))
	// }
}
