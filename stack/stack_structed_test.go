package main

import (
	"testing"
)

// go test -v -bench=. dataset.go stack_structed.go stack_structed_test.go stack_roman.go stack_roman_10000000_test.go
func BenchmarkStackStructed10000000Values(b *testing.B) {
	stack := StackStructed{}
	for _, num := range array10000000 {
		stack.Push(num)
	}
	_, origin := stack.Pop()
	for origin != nil {
		_, origin = stack.Pop()
	}
}
