package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"regexp"
)

type rules map[string]string
type messages []string
// TODO Come back to this problem if I need to learn about "recursive regular expressions", also known as "subexpressions".
// See the Ruby file for this problem for a non-stupid way.  The discussion under the comment at https://old.reddit.com/r/adventofcode/comments/kg1mro/2020_day_19_solutions/ggeajnr/ elaborates on match expression used in the Ruby program.
var mindDumb string = "42 31 | 42 42 31 31 | 42 42 42 31 31 31 | 42 42 42 42 31 31 31 31 | 42 42 42 42 42 31 31 31 31 31"
// solver[11] = "(?<r>#{solver[42]}\\g<r>?#{solver[31]})"
// For more on parsing, see https://en.wikipedia.org/wiki/CYK_algorithm.

// "Lua does not have proper regex, but it honestly never even crossed my mind
// that it was necessary for part 2. Recursively chomp off the prefix of the
// message and compare it to the permutations of rule 42. Do the same for the
// suffix and rule 31. If you end up with an empty string for your message, and
// you chomped at least 2 prefixes and 1 suffix, and you had at least 1 more
// prefix than suffixes, then you count it."

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

func recursiveRuleReplaceBadLoops(rs rules, rval string) string {
	tokens := strings.Split(rval, " ")
	for j, t := range tokens {
		if ok, _ := regexp.MatchString("[0-9]+", t); ok {
			if t == "8" {
				rs[t] = " ( 42 ) +"
			} else if t == "11" {
				rs[t] = mindDumb
			}
			tokens[j] = wrapOrRule(rs[t])
		}
	}
	s := strings.Join(tokens, " ")
	if ok, _ := regexp.MatchString(".*[0-9]+.*", s); ok {
		s = recursiveRuleReplaceBadLoops(rs, s)
	}
	s = strings.ReplaceAll(s, "\"", "")
	s = strings.ReplaceAll(s, "  ", " ")
	return s
}

func rule2Regex(rs rules, r string, part int) *regexp.Regexp {
	if part == 1 {
		r = recursiveRuleReplace(rs, r)
	} else if part == 2 {
		r = recursiveRuleReplaceBadLoops(rs, r)
	}
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

func printMatches(rx *regexp.Regexp, ms messages) {
	for _, m := range ms {
		if rx.MatchString(m) {
			fmt.Println(m)
		}
	}
}

func main() {
	// path := "19_small.txt"
	// path := "19_small2.txt"
	// path := "19_small2_loop.txt"
	path := "19.txt"
	rs, ms := parseInput(path)
	rx := rule2Regex(rs, rs["0"], 1)
	fmt.Println("Part 1:", sumMatches(rx, ms))
	// path2 := "19_small2_loop.txt"
	path2 := "19_loop.txt"
	rs, ms = parseInput(path2)
	rx = rule2Regex(rs, rs["0"], 2)
	fmt.Println("Part 2:", sumMatches(rx, ms))
	// r := regexp.MustCompile(mindDumb)
	// fmt.Println(r.Simplify())
}
