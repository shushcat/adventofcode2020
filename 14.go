package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func setFloatingAdrs(s seasys, adr int, val int) seasys {
	badr := strings.Split(strconv.FormatInt(int64(adr), 2), "")
	for len(badr) < len(s.mask) {
		badr = append([]string{"0"}, badr...)
	}
	count := 0
	for _, v := range s.mask {
		if v == "X" {
			count += 1
		}
	}
	for pval := 0; pval <= ((1 << count) - 1); pval++ {
		bpval := strings.Split(strconv.FormatInt(int64(pval), 2), "")
		for len(bpval) < count {
			bpval = append([]string{"0"}, bpval...)
		}
		k := 0
		padr := []string{}
		for j, v := range s.mask {
			switch v {
			case "0":
				padr = append(padr, badr[j])
			case "1":
				padr  = append(padr, "1")
			case "X":
				padr = append(padr, bpval[k])
				k += 1
			}
		}
		a, _ := strconv.ParseInt(strings.Join(padr, ""), 2, 64)
		s.mem[int(a)] = val
	}
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
				s= setFloatingAdrs(s, adr, val)
			}
		}
	}
	return s
}

func main() {
	s := seasys{map[int]int{}, [36]string{}}
	path := "14.txt"
	s = runProgram(path, 1)
	fmt.Println("Part 1:", s.memSum())
	s = runProgram(path, 2)
	fmt.Println("Part 2:", s.memSum())
}
