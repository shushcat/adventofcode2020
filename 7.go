package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type rule struct {
	name string
	held map[string]int
}

// func (bag rule) containedByV1(rs []rule) int {
// 	canContain := 0
// 	for i, _ := range rs {
// 		_, contains := searchContents(rs[i], rs, bag.name, 1, false)
// 		if contains {
// 			canContain += 1
// 		}
// 	}
// 	return canContain
// }

func (r rule) containedBy(rs []rule) int {
	cb := 0
	for _, v := range rs  {
		for k, _ := range v.held {
			fmt.Println(k,v)
			if k == r.name {
				return 1
			} else if k == "none" {
				return 0
			}
			// if i, ok := keyIndex(v.name, rs); ok {
				// cb += (r.containedBy(rs[i:]))
			// }
		}
	}
	return cb
}

func (r rule) numHeld(rs []rule) int {
	nh := 0
	for k, _ := range r.held {
		nh += r.held[k]
		if i, ok := keyIndex(k, rs); ok {
			nh += (rs[i].numHeld(rs) * r.held[k])
		}
	}
	return nh
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

func keyIndex(key string, rs []rule) (int, bool) {
	for i, v := range rs {
		if v.name == key {
			return i, true
		}
	}
	return 0, false
}

func ruleContains(r rule, bag string) bool {
	if _, ok := r.held[bag]; ok {
		return true
	}
	return false
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

func main() {
	// path := "7.txt"
	path := "7_small.txt"
	rs := readRules(path)
	i, _ := keyIndex("shiny gold", rs)
	fmt.Println("Part 1: ", rs[i].containedBy(rs))
	// fmt.Println("Part 1: ", rs[i].containedByV1(rs))
	fmt.Println("Part 2: ", rs[i].numHeld(rs))
}
