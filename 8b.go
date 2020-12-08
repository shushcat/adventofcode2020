package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
	"bufio"
)

type instruction struct {
	opr string
	arg int
}

func parseInstructions(path string) []instruction {
	file, _ := os.Open(path)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	instructions := []instruction{}
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		i := instruction{}
		i.opr = line[0]
		i.arg, _ = strconv.Atoi(line[1])
		instructions = append(instructions, i)
	}
	return instructions
}

func runInstructions(instructions []instruction) (int, bool) {
	trace := make([]bool, len(instructions))
	accum := 0
	for i := 0; i < len(instructions); {
		if trace[i] == true {
			return accum, false
		}
		trace[i] = true
		switch instructions[i].opr {
		case "acc":
			accum += instructions[i].arg
			i += 1
		case "jmp":
			i = i + instructions[i].arg
		case "nop":
			i += 1
		}
	}
	return accum, true
}

func mutInstructions(instructions []instruction) int {
	mutIns := make([]instruction, len(instructions))
	for i := 0; i < len(instructions); i++ {
		copy(mutIns, instructions)
		switch mutIns[i].opr {
		case "jmp":
			mutIns[i].opr = "nop"
		case "nop":
			mutIns[i].opr = "jmp"
		}
		if acc, ok := runInstructions(mutIns); ok {
			return acc
		}
	}
	return 0
}

func main() {
	// path := "8_small.txt"
	path := "8.txt"
	instructions := parseInstructions(path)
	fmt.Println(mutInstructions(instructions))
}
