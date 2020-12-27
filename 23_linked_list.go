package main

import (
	"fmt"
	"os"
	"log"
	"strconv"
	"strings"
	"runtime/pprof"
	"flag"
)

type cup struct {
	prev *cup
	val  int
	next *cup
}

type cups struct {
	head *cup
	tail *cup
}

func initCups() *cups { return &cups{} }

func (cs *cups) Populate(s string) {
	ia := str2IntAry(s)
	for i := 0; i < len(ia); i++ {
		cs.addCup(ia[i])
	}
}

func (cs *cups) bigPopulate(s string) {
	ia := str2IntAry(s)
	for i := 0; i < len(ia); i++ {
		cs.addCup(ia[i])
	}
	for val:=10; val<=1000000;val++{
		cs.addCup(val)
	}
}

func (cs *cups) addCup(n int) {
	c := &cup{val: n}
	if cs.head == nil {
		cs.head = c
	} else {
		cur := cs.tail
		cur.next = c
		c.prev = cs.tail
	}
	cs.tail = c
}

func (cs *cups) printCups() {
	cur := cs.head
	fmt.Println("==============Cups==============")
	fmt.Println(*cur)
	for cur.next != nil {
		cur = cur.next
		fmt.Println(*cur)
	}
	fmt.Println("================================")
}

func (cs *cups) step() {
	cur := cs.head
	des := setDes(cur.val)
	heldBeg, heldEnd := cur.next, ((cur.next).next).next
	cs.tail, cs.head = cur, heldEnd.next
	(cs.head).prev, (cs.tail).next = nil, nil
	cur = cs.head
	for {
		if cur.val == des {
			heldEnd.next = cur.next
			if cur.next != nil {
				(cur.next).prev = heldEnd
			}
			cur.next = heldBeg
			heldBeg.prev = cur
			break
		} else if cur.next == nil {
			cur = cs.head
			des = setDes(des)
		} else {
			cur = cur.next
		}
	}
	cur = cs.head
	for cur.next != nil {
		cur = cur.next
	}
	cur.next = cs.tail
	(cs.tail).prev = cur
}

func setDes(n int) int {
	des := 0
	if n > 1 {
		des = n - 1
	} else {
		des = 9
	}
	return des
}

func (cs *cups) nSteps(n int) {
	for i := 0; i < n; i++ {
		cs.step()
	}
}

func has(a []int, n int) bool {
	for i := 0; i < len(a); i++ {
		if n == a[i] {
			return true
		}
	}
	return false
}

func max(a []int) int {
	m := 0
	for i := 0; i < len(a); i++ {
		if a[i] > m {
			m = a[i]
		}
	}
	return m
}

func min(a []int) int {
	m := a[0]
	for i := 0; i < len(a); i++ {
		if a[i] < m {
			m = a[i]
		}
	}
	return m
}

func str2IntAry(s string) []int {
	sa := strings.Split(s, "")
	ia := make([]int, len(sa))
	for i := 0; i < len(sa); i++ {
		ia[i], _ = strconv.Atoi(sa[i])
	}
	return ia
}

func pickDes(rem []int, cur int) (int, int) {
	cur = cur - 1
	des := 0
	max, min := max(rem), min(rem)
	for des == 0 {
		if cur < min {
			des = max
		} else if has(rem, cur) {
			des = cur
		} else {
			cur = cur - 1
		}
	}
	i := 0
	for j := 0; j < len(rem); j++ {
		if rem[j] == des {
			i = j
			break
		}
	}
	return des, i
}

func p1Str(cs *cups) string {
	var str string
	cur := cs.head
	for {
		str = str + strconv.FormatInt(int64(cur.val), 10)
		if cur.next == nil {
			break
		}
		cur = cur.next
	}
	a := strings.Split(str, "1")
	str = a[1] + a[0]
	return str
}

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	steps := 10
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
		steps = 1000
	}
	input := "389125467" // Sample
	// input := "837419265" // Sample
	// input := "398254716" // Personal
	cs := initCups()
	// cs.Populate(input)
	cs.bigPopulate(input)
	cs.nSteps(steps)
	fmt.Println(cs.tail)
	// fmt.Println("Part 1:", p1Str(cs))	// -> 45798623
	// cs.printCups()
	// cs.Populate(input)
	// fmt.Println(input)
	// m1 := moveCups(input)
	// moveCups(input)
	// m1 := moveCups(input)
	// fmt.Println(m1)
	// m2 := moveCups(m1)
	// fmt.Println(m2)

}
