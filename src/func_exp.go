package main

import (
	"fmt"
)

func main() {
	x := 0
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())

}

//anonymous function -  function with no name also called as function expression which means assigning an anonymous function to a name and can use the name to call the function

//x is not a package level scope its narrowed down to curly braces level.
