package main

import "fmt"

func main() {
	switch "Jenny" {
	case "Tim", "Jenny":
		fmt.Println("Hi Tim or err Jenny")
	case "Marcus", "Medhi":
		fmt.Println("Both names start with M")
	case "Julian", "Sushant":
		fmt.Println("Hi Julian/Sushant")
	}
}

//using multiple switch case - checks if atlest one of them is true
