package main

import "fmt"

func main() {
	for i := 50; i < 140; i++ {
		fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))
	}
	foo := 'a'
	foo3 := "a"

	fmt.Println(foo)
	fmt.Println("%T\n", foo)
	fmt.Println(foo3)
	fmt.Printf("%T\n", foo3)
}

// %T - represents type
