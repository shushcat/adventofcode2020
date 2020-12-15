package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

type ship struct {
	heading string
	nsPos, ewPos int
}

type waypoint struct {
	nsPos, ewPos int
}

type instruction struct {
	dir string
	dis int
}

func parseIns(path string) []instruction {
	ins := []instruction{}
	file, _ := os.Open(path)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var in instruction
		line := scanner.Text()
		in.dir = line[:1]
		in.dis, _ = strconv.Atoi(line[1:])
		ins = append(ins, in)
	}
	return ins
}

func wrappingIndex(a []string, i int) int {
	aLen := len(a)
	j := (i % aLen)
	if j < 0 {
		j = aLen + j
	}
	return j
}

func changeShipHeading(s ship, in instruction) ship {
	headings := []string{"N", "E", "S", "W"}
	if in.dir == "L" {
		// Reverse sign for L rotations
		in.dis = -in.dis
	}
	in.dis = (in.dis / 90)
	sHeadingIndex := 0
	for k, v := range headings {
		if s.heading == v {
			sHeadingIndex = k
		}
	}
	s.heading = headings[wrappingIndex(headings, in.dis + sHeadingIndex)]
	return s
}

func moveShip(s ship, in instruction) ship {
	switch in.dir {
	case "L", "R":
		s = changeShipHeading(s, in)
	case "N":
		s.nsPos += in.dis
	case "S":
		s.nsPos -= in.dis
	case "E":
		s.ewPos += in.dis
	case "W":
		s.ewPos -= in.dis
	case "F":
		in.dir = s.heading
		s = moveShip(s, in)
	}
	return s
}

func rotateWaypoint(w waypoint, in instruction) waypoint {
	if in.dir == "L" {
		in.dis = (360 - in.dis)/90
	} else {
		in.dis = in.dis/90
	}
	for i:= 0; i < in.dis; i++ {
		ns, ew := w.nsPos, w.ewPos
		w.nsPos, w.ewPos = -ew, ns
	}
	return w
}

func moveWaypoint(w waypoint, in instruction) waypoint {
	switch in.dir {
	case "L", "R":
		w = rotateWaypoint(w, in)
	case "N":
		w.nsPos += in.dis
	case "S":
		w.nsPos -= in.dis
	case "E":
		w.ewPos += in.dis
	case "W":
		w.ewPos -= in.dis
	}
	return w
}

func navigate(s ship, ins []instruction) ship {
	for _, in := range ins {
		s = moveShip(s, in)
	}
	return s
}

func approachWaypoint(s ship, w waypoint, in instruction) ship {
	for i := 0; i < in.dis; i++ {
		s.nsPos = s.nsPos + w.nsPos
		s.ewPos = s.ewPos + w.ewPos
	}
	return s
}

func navigateByWaypoint(s ship, ins []instruction) ship {
	w := waypoint{1, 10}
	for _, in := range ins {
		if in.dir == "F" {
			s = approachWaypoint(s, w, in)
		} else {
			w = moveWaypoint(w, in)
		}
	}
	return s
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func manhattanDistance(s ship) int {
	return abs(s.nsPos) + abs(s.ewPos)

}

func main() {
	// path := "12_small.txt"
	path := "12.txt"
	instructions := parseIns(path)
	s := ship{"E", 0, 0}
	s = navigate(s, instructions)
	fmt.Println("Part 1:", manhattanDistance(s))
	s = ship{"E", 0, 0}
	s = navigateByWaypoint(s, instructions)
	fmt.Println("Part 2:", manhattanDistance(s)) // want 30761
}
