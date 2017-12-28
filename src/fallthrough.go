package main

import "fmt"

func main() {
	switch "Marcus" {
	case "Tim":
		fmt.Println("Hi TIm")
	case "Jenny":
		fmt.Println("Hi Jenny")
	case "Marcus":
		fmt.Println("Hi Marcus")
		fallthrough
	case "Medhi":
		fmt.Println("Hi Medhi")
		fallthrough
	case "Julian":
		fmt.Println("Hi Julian")
	case "Sushant":
		fmt.Println("Hi Sushant")
	}
}

// Simple switch and fall through statements
