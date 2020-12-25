package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type allergins map[string][]string
type foods map[string]int

func (f foods) Tally(fs []string) foods {
	for i:=0; i<len(fs); i++ {
		f[fs[i]] = f[fs[i]] + 1
	}
	return f
}

func pruneAllergins(as allergins) allergins {
	for a1, f1 := range as {
		if len(f1) == 1 {
			for a2, f2 := range as {
				if len(f2) > 1 {
					as[a2] = complement(as[a2], as[a1])
					as = pruneAllergins(as)
				}
			}
		}
	}
	return as

}

func readFoods(path string) (allergins, foods) {
	bytes, _ := ioutil.ReadFile(path)
	input := string(bytes)
	as := make(map[string][]string)
	fs := foods{}
	for _, l := range strings.Split(input, "\n") {
		if l != "" {
			fa := strings.Split(l, " (contains ")
			fa[1] = strings.TrimSuffix(fa[1], ")")
			fa[1] = strings.ReplaceAll(fa[1], ",", "")
			foodAry := strings.Split(fa[0], " ")
			// fs = union(fs, foodAry)
			fs = fs.Tally(foodAry)
			allerginAry := strings.Split(fa[1], " ")
			for _, a := range allerginAry {
				if foods, ok := as[a]; ok {
					// Add ingredients as set intersection
					// with the ingredients that have
					// already been added for this
					// allergin.
					as[a] = intersection(foods, foodAry)
				} else {
					// If the allergin hasn't been added to
					// the map yet, then add it along with
					// its ingredients.
					as[a] = foodAry
				}
			}
		}
	}
	return as, fs
}

func rm(a []string, i int) []string {
	a[i] = a[len(a)-1]
	return a[:len(a)-1]
}

func intersection(a1 []string, a2 []string) []string {
	a1Has := make(map[string]bool)
	var in []string
	for i:=0;i<len(a1);i++{
		a1Has[a1[i]]=true
	}
	for i:=0;i<len(a2);i++{
		if _, ok := a1Has[a2[i]]; ok {
			in = append(in, a2[i])
		}
	}
	return in
}

func union(a1 []string, a2 []string) []string {
	a1Has := make(map[string]bool)
	var un []string
	for i:=0;i<len(a1);i++{
		a1Has[a1[i]]=true
		un = append(un, a1[i])
	}
	for i:=0;i<len(a2);i++{
		if _, ok := a1Has[a2[i]]; !ok {
			un = append(un, a2[i])
		}
	}
	return un
}

func complement(a1 []string, a2 []string) []string {
	a2Has := make(map[string]bool)
	for i:=0;i<len(a2);i++{
		a2Has[a2[i]] = true
	}
	var a3 []string
	for i:=0;i<len(a1);i++{
		if !a2Has[a1[i]] {
			a3 = append(a3, a1[i])
		}
	}
	return a3
}

func tests() {
	s1 := []string{"a","b","c","d","e","f","g","h"}
	s2 := []string{"e","c","z","u"}
	fmt.Println(intersection(s1,s2))
	fmt.Println(union(s1,s2))
	// fmt.Println(rm(s1, 2))
	s3 := []string{"self"}
	s4 := []string{"self"}
	fmt.Println(intersection(s3,s4))
	fmt.Println(complement(s1, s2))
	s5 := []string{"a","f","e","c","h", "w"}
	fmt.Println(complement(s1, s5))
}

func numHypoallergenic(fs foods, as allergins) int {
	hypo := foods{}
	for k, v := range fs {
		hypo[k] = v
	}
	for _, v := range as {
		delete(hypo, v[0])
	}
	num := 0
	for _, v := range hypo {
		num += v
	}
	return num
}

func main() {
	// path := "21_small.txt"
	path := "21.txt"
	as, fs := readFoods(path)
	as = pruneAllergins(as)
	fmt.Println("Part 1:", numHypoallergenic(fs, as))
}
