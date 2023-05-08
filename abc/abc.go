package main

import "fmt"

func main() {
	one := 1
	two := 2
	three := 3
	a := []*int{&one, &two, &three}
	for _, item := range a {
		fmt.Println("OK", item, *item)
	}

	fmt.Println()
	two = *new(int)
	for _, item := range a {
		fmt.Println("OK", item, *item)
	}

	fmt.Println()
	a[1] = nil
	for _, item := range a {
		fmt.Println("OK", item)
	}
}
