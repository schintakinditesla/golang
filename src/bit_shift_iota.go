package main

import (
	"fmt"
)

const (
	_  = iota             //0  blank identifier assigned with iota first time so 0 value
	KB = 1 << (iota * 10) // same as 1 << (1 * 10) iota used 2nd time
	MB = 1 << (iota * 10) // same as 1 << (2 * 10) because iota used third time
)

func main() {
	fmt.Println("Binary\t\t\tDecimal")
	fmt.Printf("%b\t\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t", MB)
	fmt.Printf("%d\n", MB)
}
