package main

import (
	"fmt"
)

func handshakeKey(pub1 int, pub2 int) int {
	tran, subn, loop := 1, 7, 0
	for tran != pub1 {
		tran = (subn * tran) % 20201227
		loop++
	}
	tran = 1
	for i := 0; i < loop; i++ {
		tran = (pub2 * tran) % 20201227
	}
	return tran
}

func main() {
	// pub1, pub2 := 5764801, 17807724		// Sample
	pub1, pub2 := 18499292, 8790390 // Personal
	fmt.Print(handshakeKey(pub1, pub2), "...and done!\n")
}
