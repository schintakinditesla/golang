package main

import (
	"fmt"
)

func main() {
	a := 43

	fmt.Println(a)  //43
	fmt.Println(&a) // hex address

	var b *int = &a

	fmt.Println(b)  // hex address
	fmt.Println(*b) //43
	fmt.Println(&b)

	*b = 42
	fmt.Println(a)

}
