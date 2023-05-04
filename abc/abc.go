package main

import "fmt"

func main() {
	a := [3]float64{0, 1, 2}
	b := [3]float64{0, 1, 2}

	if a == b {
		fmt.Println("OK")
	} else {
		fmt.Println("ERR")

	}
}
