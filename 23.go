package main

import (
	"fmt"
	"strings"
)

func aHas(a []string, elem string) bool {
	for i:=0;i<len(a);i++{
		if elem == a[i] {
			return true
		}
	}
	return false
}

func moveCups(in []string) {
	fmt.Println(in)
	cur, held, rem := in[0], in[1:4], in[4:]
	next := cur-1

	fmt.Println(cur)
	fmt.Println(held)
	fmt.Println(rem)

}

func main() {
	input := strings.Split("389125467", "")		// Sample
	// input := strings.Split("398254716", "")	// Personal
	moveCups(input)
}
