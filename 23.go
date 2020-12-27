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

func populate(s string) map[int]int{
	ia := strings.Split(s, "")
	cs := make(map[int]int, len(ia))
	for i := 0; i< len(ia)-1;i++{
		d1, _ := strconv.Atoi(ia[i])
		d2, _ := strconv.Atoi(ia[i+1])
		cs[d1]=d2
	}
	return cs
}

func bigPopulate(s string) map[int]int {
	ia := strings.Split(s, "")
	cs := make(map[int]int, len(ia))
	for i := 0; i< len(ia)-1;i++{
		d1, _ := strconv.Atoi(ia[i])
		d2, _ := strconv.Atoi(ia[i+1])
		cs[d1]=d2
	}
	for i:=10; i<=1000000;i++{
		cs[i]=i
	}
	return cs
}

func step(cs map[int]int) map[int]int {
	head, tail := head(cs), tail(cs)
	cur := head
	des := setDes(cur)
	hBeg, hEnd := cs[cur], cs[cs[cs[cur]]]
	cur = cs[hEnd]
	for {
		if cur == des {
			if cur != tail {
				cs[hEnd] = cs[des]
				cs[tail] = head
			} else {
				cs[hEnd] = head
			}
			cs[des] = hBeg
			delete(cs, head)
			break
		} else if cur == tail {
			cur = cs[hEnd]
			des = setDes(des)
		} else {
			cur = cs[cur]
		}
	}
	return cs
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

func nSteps(cs map[int]int, n int) {
	for i := 0; i < n; i++ {
		fmt.Println("Step", i+1)
		cs = step(cs)
	}
}

func head(cs map[int]int) int {
	refd := make(map[int]bool)
	head := 0
	for _, v := range cs {
		refd[v] = true
	}
	for k, _ := range cs {
		if !(refd[k]) {
			head = k
			break
		}
	}
	return head
}

func tail(cs map[int]int) int {
	refd := make(map[int]bool)
	tail := 0
	for k, _ := range cs {
		refd[k] = true
	}
	for _, v := range cs {
		if !(refd[v]) {
			tail = v
			break
		}
	}
	return tail
}

func printCups(cs map[int]int) {
	cur := head(cs)
	for i:=0;i<=len(cs);i++ {
		fmt.Print(cur)
		cur = cs[cur]
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

func p1Str(cs map[int]int) string {
	var str string
	cur, tail := head(cs), tail(cs)
	for {
		str = str + strconv.FormatInt(int64(cur), 10)
		if cur == tail {
			break
		}
		cur = cs[cur]
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
	input := "389125467" // Sample
	// input := "289154673" // Sample
	// input := "837419265" // Sample
	// input := "398254716" // Personal
	cs1 := populate(input)
	nSteps(cs1, 100)
	fmt.Println("Part 1:", p1Str(cs1))	// -> 45798623
	cs2 := bigPopulate(input)
	nSteps(cs2, 10)
	fmt.Println(cs2[1])
}
