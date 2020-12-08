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

// acc: add to Accumulator
// jmp: adds to instruction index; +1 -> next, -1 -> prev, &c
// nop: impotent
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

// func execInstruction(instruction) bool {
// }


func runInstructions(instructions []instruction) int {
	trace := make([]bool, len(instructions))
	accum := 0
	for i := 0; i < len(instructions); {
		fmt.Println(i, instructions[i], trace[i])
		if trace[i] == true {
			break
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
	fmt.Println(trace)
	return accum
}

func main() {
	// path := "8_small.txt"
	path := "8.txt"
	instructions := parseInstructions(path)
	fmt.Println(runInstructions(instructions))
	// fmt.Println(instructions)
}
