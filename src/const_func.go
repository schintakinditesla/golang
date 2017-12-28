package main

import "fmt"

const (
	p string = "checking const"
	r bool   = true
) //multiple constant declaration

func main() {
	const q = 42

	fmt.Println("p - ", p)
	fmt.Println("q - ", q)
	fmt.Println("r - ", r)
}

//checking how constants work in golang
