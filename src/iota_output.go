package main

import "fmt"

const (
	A = iota //0
	B = iota //1
	C = iota //2
)

const (
	D = iota //0
	E        //1
	F        //2
)

func main() {
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println("F in a different cosnt declaration scope using iota: ", F)
}

// iota keeps incrementing itself when used more than once in the same scope within the parenthesis.

/* above iota declaration is similar to below
const(
	A = iota //0
	B
	C
)*/
