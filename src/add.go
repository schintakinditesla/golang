package main

import (
	"fmt"
)

func add(x int, y int) int { // same as func add(x, y int) (int){ .... }

	return x + y
}

func main() {
	addition := add(10, 20)
	fmt.Printf("%v\n", addition)
	// also can use fmt.Println(add(42,13)) directly
}
