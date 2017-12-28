package main

import "fmt"

func main() {
	i := 0
	for {
		fmt.Println(i)
		if i >= 10 {
			break
		}
		i++
	}
}

//using infinite for loop and break condition using if and break statement.
