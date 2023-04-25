package main

import (
	"math/rand"
)

func FunctionWithName() (func() int, *func() int) {
	anonymousFunction := func() int {
		return rand.Intn(10)
	}
	return anonymousFunction, &anonymousFunction
}
