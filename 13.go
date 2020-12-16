package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type schedule struct {
	time   int
	busses []int
}

func parseSchedule(path string) schedule {
	file, _ := os.Open(path)
	defer file.Close()
	reader, sched := bufio.NewReader(file), schedule{}
	timeStr, _ := reader.ReadString('\n')
	sched.time, _ = strconv.Atoi(strings.TrimSpace(timeStr))
	busStr, _ := reader.ReadString('\n')
	busses := strings.Split(strings.TrimSpace(busStr), ",")
	for _, b := range busses {
		if b == "x" {
			sched.busses = append(sched.busses, 0)
		} else {
			v, _ := strconv.Atoi(b)
			sched.busses = append(sched.busses, v)
		}
	}
	return sched
}

func nextBus(s schedule) (busID int, waitTime int) {
	busID, waitTime = 0, s.time
	for _, b := range s.busses {
		if b != 0 {
			bTime := (s.time - (s.time % b)) + b
			if bTime-s.time < waitTime {
				busID, waitTime = b, bTime-s.time
			}
		}
	}
	return busID, waitTime
}

func aryPosProd(ary []int) int {
	prod := 1
	for i := 0; i < len(ary); i++ {
		if ary[i] != 0 {
			prod = prod * ary[i]
		}
	}
	return prod
}

func departureRunLen(s schedule, t int) int {
	runLen := 0
	for i, b := range s.busses {
		if b > 0 {
			if (t+i)%b != 0 {
				return runLen
			}
		}
		runLen += 1
	}
	return runLen
}

func departureRunSearch(s schedule) int {
	period, perSlcLen := 1, 1
	time := s.busses[0]
	for {
		runLen := departureRunLen(s, time)
		if runLen == len(s.busses) {
			return time
		} else if perSlcLen < runLen {
			period, perSlcLen = 1, runLen
			period = aryPosProd(s.busses[:perSlcLen])
		}
		time += period
	}
	return 0
}

func main() {
	// path := "13_small1.txt"	// 3417
	// path := "13_small2.txt"	// 754018
	// path := "13_small3.txt"	// 779210
	// path := "13_small4.txt"	// 1261476
	// path := "13_small5.txt"	// 1202161486
	// path := "13_small.txt"
	path := "13.txt"
	schedule := parseSchedule(path)
	nextBus, waitTime := nextBus(schedule)
	fmt.Println("Part 1:", nextBus*waitTime)
	schedule = parseSchedule(path)
	fmt.Println("Part 2:", departureRunSearch(schedule))
}
