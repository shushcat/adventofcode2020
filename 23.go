package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"strconv"
	"strings"
)

type cups struct {
	h int
	m    []int
	max  int
}

func str2Ints(s string) []int {
	sa := strings.Split(s, "")
	ia := make([]int, len(sa))
	for i := 0; i < len(ia); i++ {
		d, _ := strconv.Atoi(sa[i])
		ia[i] = d
	}
	return ia
}

func populate(input string, size int) *cups {
	ia := str2Ints(input)
	cs := &cups{}
	cs.m = make([]int, size+1)
	cs.h = ia[0]
	cs.max = 9
	for i := 0; i < len(ia)-1; i++ {
		d1, d2 := ia[i], ia[i+1]
		cs.m[d1] = d2
	}
	last := ia[len(ia)-1]
	if size > 9 {
		cs.m[last] = 10
		last = 10
		cs.max = size
		for i:= 10; i<size;i++{
			cs.m[i]=i+1
			last = i
		}
	}
	cs.m[last] = cs.h
	return cs
}

func (cs *cups) play(rounds int) {
	for {

		// Three cups
		cup1 := cs.m[cs.h]
		cup2 := cs.m[cup1]
		cup3 := cs.m[cup2]
		post := cs.m[cup3]

		dest := cs.h - 1
		if dest < 1 {
			dest = cs.max
		}
		for cup1 == dest || cup2 == dest || cup3 == dest {
		dest--
			if dest < 1 {
				dest = cs.max
			}
		}

		// Remove cups
		cs.m[cs.h] = post

		// Insert them
		oldDestVal := cs.m[dest]
		cs.m[dest] = cup1
		cs.m[cup3] = oldDestVal

		cs.h = post

		rounds--
		if rounds == 0 {
			break
		}
	}
}

func printCups(cs *cups) {
	cur := cs.h
	for i := 1; i <= len(cs.m)-1; i++ {
		fmt.Print(cur)
		cur = cs.m[cur]
	}
	fmt.Print("\n")
}

func p1Str(cs *cups) string {
	var str string
	cur := cs.h
	for i := 1; i <= len(cs.m)-1; i++ {
		d := strconv.FormatInt(int64(cur), 10)
		str = str + d
		cur = cs.m[cur]
	}
	sa := strings.Split(str, "1")
	return sa[1]+sa[0]
}

func p2Prod(cs *cups) int {
	n1 := cs.m[1]
	n2 := cs.m[n1]
	prod := n1 * n2
	return prod
}

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	fmt.Println("=============================")
	fmt.Println("============TEST=============")
	fmt.Println("=============================\n")
	input := "389125467" // Sample
	// input := "289154673" // Sample
	// input := "837419265" // Sample
	// input := "398254716" // Personal
	// cs1 := populate(input)
	// printCups(cs1)
	// cs1.nSteps(100)

	cs1 := populate(input, 9)
	cs1.play(100)
	fmt.Println("Part 1:", p1Str(cs1))

	cs2 := populate(input, 9)
	cs2.play(10_000_000)
	fmt.Println("Part 2:", p2Prod(cs2))

	fmt.Println(cs2.m[1])

	fmt.Println("done")
}
