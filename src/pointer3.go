package main

import "fmt"

func zero(z *int) {
	fmt.Printf("printing z address inside function: \n%p\n", z)
	*z = 0
}

func main() {
	x := 5
	fmt.Println("Before zero function")
	fmt.Println(x)
	fmt.Printf("x address: %p\n", &x)
	zero(&x)
	fmt.Println("After zero function")
	fmt.Println(x)
}
