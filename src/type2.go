package main

import (
	"fmt"
	"time"
)

func timeMap(y interface{}) {
	z, ok := y.(map[string]interface{})
	if ok {
		z["updated_at"] = time.Now()
	}
}

func main() {
	foo := map[string]interface{}{
		"Matt": 42,
	}
	timeMap(foo)
	fmt.Println(foo)
}

//timemap function takes a value and can be asserted as a map of interfaces{} keyed by strings.
