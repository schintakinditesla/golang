package main

import (
	"fmt"
)

func incr(p *int) int {
	*p++ // increments what p points to; does not change p
	fmt.Printf("Value of p inside incr: %d\n", *p)
	return *p
}

func main() {
	v := 1
	incr(&v)
	fmt.Println(incr(&v))
}
