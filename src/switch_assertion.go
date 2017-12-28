package main

import "fmt"

type Contact struct {
	greeting string
	name     string
}

// normally we switch on value of variable
// go allows you to switch on type of variable

func SwitchOnType(x interface{}) {
	switch x.(type) { // this is an assert that "x if of this particular type"
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case Contact:
		fmt.Println("contact")
	default:
		fmt.Println("unknown")
	}
}

func main() {
	SwitchOnType(7)
	SwitchOnType("Savitha")
	var t = Contact{"Good to see you", "Tim"}
	SwitchOnType(t)
}
