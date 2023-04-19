package main

import (
	"testing"
)

// go test -v -bench=. dataset.go stack_roman.go stack_roman_test.go
func BenchmarkSimplestRoman10Values(b *testing.B) {
	stack := StackRoman{}
	for num := range array10 {
		stack.Push(num)
	}
	for stack.len > 0 {
		stack.Pop()
	}
}

func BenchmarkSimplestRoman1000Values(b *testing.B) {
	stack := StackRoman{}
	for num := range array1000 {
		stack.Push(num)
	}
	for stack.len > 0 {
		stack.Pop()
	}
}

func BenchmarkSimplestRoman100000Values(b *testing.B) {
	stack := StackRoman{}
	for num := range array100000 {
		stack.Push(num)
	}
	for stack.len > 0 {
		stack.Pop()
	}
}
