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

func (r rule) numHeld(rs []rule) int {
	nh := 0
	for k, _ := range r.held {
		nh += r.held[k]
		if i, ok := keyIndex(k, rs); ok {
			nh += (rs[i].numHeld(rs) * r.held[k])
		}
	}
	fmt.Print(r, ".numHeld() == ", nh, "\n")
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

func main() {
	// path := "7_small.txt"
	// path := "7_small2.txt"
	path := "7.txt"
	rs := readRules(path)
	var r rule
	for k, v := range rs {
		if v.name == "shiny gold" {
			 fmt.Println(k)
			 r = v
		}
	}
	fmt.Println(r, "holds", r.numHeld(rs))
}
