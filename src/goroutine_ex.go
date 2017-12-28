package main

import (
	"fmt"
	"time"
)

func f(from string) {
	//	a := makeTimestamp()
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
		a := makeTimestamp()
		fmt.Printf("%d\n", a)
	}
}

func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func main() {
	f("direct")

	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
