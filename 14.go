package main

import (
	"fmt"
	"strconv"
	"strings"
	"os"
	"bufio"
)

type seasys struct {
	mem  map[int]int
	mask [36]string
}

func (s seasys) memSum() int {
	sum := 0
	for _, n := range s.mem {
		sum += n
	}
	return sum
}

func setMask(s seasys, mask string) seasys {
	for i, c := range strings.Split(mask, "") {
		s.mask[i] = c
	}
	return s
}

func setAdr(s seasys, adr int, val int) seasys {
	bval := strings.Split(strconv.FormatInt(int64(val), 2), "")
	for i := 36 - len(bval); i > 0; i-- {
		bval = append([]string{"0"}, bval...)
	}
	for i := 0; i < 36; i++ {
		if s.mask[i] != "X" {
			bval[i] = s.mask[i]
		}
	}
	if v, err := strconv.ParseInt(strings.Join(bval, ""), 2, 64); err == nil {
		s.mem[adr] = int(v)
	}
	return s
}

func (s seasys) addBinAdr(binadr []string, val int) {
	if a, err := strconv.ParseInt(strings.Join(binadr, ""), 2, 64); err == nil {
		s.mem[int(a)] = val
	}
}

func setFloatingAdrs(s seasys, badr []string, val int) seasys {
	fmt.Println("Adding", badr, "variants")
	ladr := make([]string, len(badr))
	for i, v := range(badr) {
		ladr[i] = v
	}
	for i := 0; i < 36; i++ {
		if badr[i] == "X" {
			ladr[i] = "0"
		}
	}
	s.addBinAdr(ladr, val)
	fmt.Println(s.mem)
	for i := 0; i < 36; i++ {
		if badr[i] == "X" {
			// fmt.Println(s.mem)
			ladr[i] = "1"
			s.addBinAdr(ladr, val)
		}
	}
	return s
}

func setAdr2(s seasys, adr int, val int) seasys {
	badr := strings.Split(strconv.FormatInt(int64(adr), 2), "")
	// Prepend to badr array
	for i := 36 - len(badr); i > 0; i-- {
		badr = append([]string{"0"}, badr...)
	}
	for i := 0; i < 36; i++ {
		if (s.mask[i] == "1") || (s.mask[i] == "X") {
			badr[i] = s.mask[i]
		}
	}
	s = setFloatingAdrs(s, badr, val)
	return s
}

func runProgram(path string, version int) seasys {
	file, _ := os.Open(path)
	defer file.Close()
	s := seasys{map[int]int{}, [36]string{}}
	scanner := bufio.NewScanner(file)
	var instruction []string
	for scanner.Scan() {
		instruction = strings.Split(scanner.Text(), " = ")
		switch instruction[0][:4] {
		case "mask":
			s = setMask(s, instruction[1])
		case "mem[":
			instruction[0] = strings.TrimPrefix(instruction[0], "mem[")
			instruction[0] = strings.TrimSuffix(instruction[0], "]")
			adr, _ := strconv.Atoi(instruction[0])
			val, _ := strconv.Atoi(instruction[1])
			switch version {
			case 1:
				s = setAdr(s, adr, val)
			case 2:
				s = setAdr2(s, adr, val)
			}
		}
	}
	return s
}

// func permuteSmall(a []string) []string {
// 	fmt.Print(a, " -> ")
// 	b := a
// 	for i := 0; i < len(a); i++ {
// 		if a[i] == "X" {
// 			b[i] = "1"
// 			permuteSmall(b)
// 			b := a
// 			b[i] = "0"
// 			permuteSmall(b)
// 		}
// 	}
// 	fmt.Print(b, "\n")
// 	return b
// }

func permuteSmall(a []string) {
	count := 1
	for _, v := range a {
		if v == "X" {
			count += 1
		}
	}
	for i := 0; i <= count; i++ {
		bval := strings.Split(strconv.FormatInt(int64(i), 2), "")
		fmt.Println(bval)
	}
}

func main() {
	// path := "14_small.txt"
	// path := "14_small2.txt"
	// s := seasys{}
	// path := "14.txt"
	// s = runProgram(path, 1)
	// fmt.Println("Part 1:", s.memSum())
	// s = runProgram(path, 2)
	// fmt.Println(s)
	small := []string{"X", "0", "X"}
	permuteSmall(small)
}
