package main

import (
	"fmt"
)

func main() {
	myfrnds := "Mar"

	switch {
	case len(myfrnds) == 2:
		fmt.Println("Len of myfrnds 2")
	case myfrnds == "Tim":
		fmt.Println("hi tim")
	case myfrnds == "jenny":
		fmt.Println("hi jenny")
	case myfrnds == "Marcus", myfrnds == "Medhi":
		fmt.Println("Hey Marcus, or Hey Medhi")
	case myfrnds == "julian":
		fmt.Println("hi julian")
	default:
		fmt.Println("nothing matched, this is default")
	}
}

// switch statement with expression that evaluates and makes switch statement work like if/else

//if nothing matched runs default.
