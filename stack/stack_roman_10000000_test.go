package main

import (
	"testing"
)

// go test -v -bench=. dataset.go stack_roman.go stack_roman_10000000_test.go
func BenchmarkRoman10000000Values(b *testing.B) {
	stack := StackRoman{}
	for _, num := range array10000000 {
		stack.Push(num)
	}
	for stack.len > 0 {
		stack.Pop()
	}
}
