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
	h, t int
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
	for i := 0; i < len(ia)-1; i++ {
		d1, d2 := ia[i], ia[i+1]
		cs.m[d1] = d2
	}
	cs.setTail()
	if size > 9 {
		cs.m[cs.t] = 10
		for i:= 10; i<size;i++{
			cs.m[i]=i+1
		}
		cs.t = size
		cs.m[size] = 0
	}
	return cs
}

func (cg *cups) Play(rounds int) {
	for {
		// Take three cups
		cup1 := cg.m[cg.h]
		cup2 := cg.m[cup1]
		cup3 := cg.m[cup2]
		after := cg.m[cup3]

		destination := setDes(cg.h)

		// Remove the three cups
		cg.m[cg.h] = after

		// Insert them after the destination
		fmt.Println(cg.m)
		oldDestValue := cg.m[destination]
		cg.m[destination] = cup1
		cg.m[cup3] = oldDestValue

		cg.h = after

	}
}

func (cs *cups) step() {
	head, tail := cs.h, cs.t
	cur, des := head, head-1
	hBeg, hEnd := cs.m[cur], cs.m[cs.m[cs.m[cur]]]
	cur, cs.h = cs.m[hEnd], cs.m[hEnd]
	for {
		if cur == des {
			if cur != tail {
				cs.m[hEnd] = cs.m[des]
				cs.m[cs.t] = head
			} else {
				cs.m[hEnd] = head
			}
			cs.t = head
			cs.m[des] = hBeg
			cs.m[head] = 0
			break
		} else if cur == tail {
			cur = cs.m[hEnd]
			des = setDes(des)
		} else {
			cur = cs.m[cur]
		}
		// rounds--
		// if rounds == 0 {
		// break
		// }
	}
}

func setDes(n int) int {
	des := 0
	if n > 1 {
		des = n - 1
	} else {
		des = 1000000
	}
	return des
}

func (cs *cups) nSteps(n int) {
	for i := 0; i < n; i++ {
		fmt.Println("Step", i+1)
		cs.step()
	}
}

func (cs *cups) setTail() {
	for k := 1; k < len(cs.m); k++ {
		if cs.m[k] == 0 {
			cs.t = k
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
	for {
		str = str + strconv.FormatInt(int64(cur), 10)
		if cur == cs.t {
			break
		}
		cur = cs.m[cur]
	}
	a := strings.Split(str, "1")
	str = a[1] + a[0]
	return str
}

func p2Prod(cs *cups) int {
	n1 := cs.m[1]
	n2 := cs.m[n1]
	prod := n1 * n2
	fmt.Println(n1, "*", n2)
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
	input := "389125467" // Sample
	// input := "289154673" // Sample
	// input := "837419265" // Sample
	// input := "398254716" // Personal
	// cs1 := populate(input)
	// printCups(cs1)
	// cs1.nSteps(100)
	// fmt.Println("Part 1:", p1Str(cs1))

	// cs2 := bigPopulate(input)

	cs1 := populate(input, 100)
	// cs1.step()
	fmt.Println(cs1)
	printCups(cs1)
	// cs2 := populate(input)
	// cs2.Play(1)
	// fmt.Println(cs2)

	fmt.Println("done")
	// fmt.Println(cs2, len(cs2.m), cs2.h, cs2.t, cs2.m[1:20])
	// cs2.nSteps(10_000_000)
	// cs2.nSteps(100)
	// fmt.Println(cs2.h, cs2.t)
	// fmt.Println(len(cs2.m))

	// fmt.Println(cs2.m[1])
	// printCups(cs1)
	// fmt.Println(cs2, cs2.head(), cs2.tail())
	// nSteps(cs2, 10)
	// fmt.Println(cs2[1])
}
