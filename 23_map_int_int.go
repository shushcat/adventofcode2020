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
	m    map[int]int
}

func populate(s string) *cups {
	ia := strings.Split(s, "")
	cs := &cups{}
	cs.m = make(map[int]int, len(ia))
	for i := 0; i < len(ia)-1; i++ {
		d1, _ := strconv.Atoi(ia[i])
		d2, _ := strconv.Atoi(ia[i+1])
		cs.m[d1] = d2
	}
	cs.setHead()
	cs.setTail()
	return cs
}

func bigPopulate(s string) *cups {
	ia := strings.Split(s, "")
	cs := &cups{}
	cs.h, cs.t = 0, 0
	cs.m = make(map[int]int, len(ia))
	for i := 0; i < len(ia)-1; i++ {
		d1, _ := strconv.Atoi(ia[i])
		d2, _ := strconv.Atoi(ia[i+1])
		cs.m[d1] = d2
	}
	for i := 10; i <= 1000000; i++ {
		cs.m[i] = i
	}
	cs.setHead()
	cs.setTail()
	return cs
}

func (cs *cups) step() {
	head, tail := cs.h, cs.t
	cur := head
	des := head - 1
	hBeg, hEnd := cs.m[cur], cs.m[cs.m[cs.m[cur]]]
	cur = cs.m[hEnd]
	cs.h = cs.m[hEnd]
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
			delete(cs.m, head)
			break
		} else if cur == tail {
			cur = cs.m[hEnd]
			des = setDes(des)
		} else {
			cur = cs.m[cur]
		}
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

func (cs *cups) setHead() int {
	refd := make(map[int]bool)
	for _, v := range cs.m {
		if _, ok := cs.m[v]; ok {
			refd[v] = true
		}
	}
	for k, _ := range cs.m {
		if !(refd[k]) {
			cs.h = k
			break
		}
	}
	return cs.h
}

func (cs *cups) setTail() {
	for _, v := range cs.m {
		if _, ok := cs.m[v]; !ok {
			cs.t = v
			break
		}
	}
}

func printCups(cs *cups) {
	cur := cs.h
	for i := 0; i <= len(cs.m); i++ {
		fmt.Print(cur)
		cur = cs.m[cur]
	}
	fmt.Print("\n")
}

func has(a []int, n int) bool {
	for i := 0; i < len(a); i++ {
		if n == a[i] {
			return true
		}
	}
	return false
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
	// input := "389125467" // Sample
	// input := "289154673" // Sample
	// input := "837419265" // Sample
	input := "398254716" // Personal
	cs1 := populate(input)
	cs1.nSteps(100)
	fmt.Println("Part 1:", p1Str(cs1))
	cs2 := bigPopulate(input)
	cs2.nSteps(100)
	// fmt.Println(cs2.m[1])
	// printCups(cs1)
	// fmt.Println(cs2, cs2.head(), cs2.tail())
	// nSteps(cs2, 10)
	// fmt.Println(cs2[1])
}
