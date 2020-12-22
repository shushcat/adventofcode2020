package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"regexp"
)

type rules map[string]string
type messages []string

func parseInput(path string) (rules, messages) {
	b, _ := ioutil.ReadFile(path)
	input := string(b)
	rs := rules{}
	ms := messages{}
	for _, l := range strings.Split(input, "\n") {
		if l != "" {
			ru, _ := regexp.Compile("\\d+")
			spec, _ := regexp.Compile(":\\s.*$")
			if ru.MatchString(l){
				rs[ru.FindString(l)] = spec.FindString(l)[2:]
			} else {
				ms = append(ms, l)
			}
		}
	}
	return rs, ms
}

func wrapOrRule(s string) string {
	if ok, _ := regexp.MatchString(".*\\|.*", s); ok {
		s = " ( " + s + " ) "
	}
	return s
}

func recursiveRuleReplace(rs rules, rval string) string {
	tokens := strings.Split(rval, " ")
	for j, t := range tokens {
		if ok, _ := regexp.MatchString("[0-9]+", t); ok {
			tokens[j] = wrapOrRule(rs[t])
		}
	}
	s := strings.Join(tokens, " ")
	if ok, _ := regexp.MatchString(".*[0-9]+.*", s); ok {
		s = recursiveRuleReplace(rs, s)
	}
	s = strings.ReplaceAll(s, "\"", "")
	s = strings.ReplaceAll(s, "  ", " ")
	return s
}

func rule2Regex(rs rules, r string) *regexp.Regexp {
	r = recursiveRuleReplace(rs, r)
	r = strings.ReplaceAll(r, " ", "")
	r = "^" + r + "$"
	rx := regexp.MustCompile(r)
	return rx
}

func sumMatches(rx *regexp.Regexp, ms messages) int {
	sum := 0
	for _, m := range ms {
		if rx.MatchString(m) {
			sum += 1
		}
	}
	return sum
}

func main() {
	// path := "19_small.txt"
	path := "19.txt"
	rs, ms := parseInput(path)
	rx := rule2Regex(rs, rs["0"])
	fmt.Println("Part 1:", sumMatches(rx, ms))
}
