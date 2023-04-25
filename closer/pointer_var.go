package main

import (
	"fmt"
)

func main() {
	var i int
	fmt.Printf("  i = %p\n", i)  // Default value of var `i` is 0
	fmt.Printf(" &i = %p\n", &i) // Address of var `i` store
	var k *int
	fmt.Printf("  k = %p\n", k)  // Default value is 0x0
	fmt.Printf(" &k = %p\n", &k) // Address of var `k` with type point to int
}
