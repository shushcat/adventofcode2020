package main

import (
	"fmt"
)

type coord struct {
	x, y, z int
}

type cubes map[coord]bool

func main() {
	cs := make(cubes)
	cs[coord{0,0,0}] = true
	cs[coord{0,70,0}] = true
	fmt.Println(cs)
}
