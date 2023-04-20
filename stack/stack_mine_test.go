package main

import (
	"testing"
)

// go test -v -bench=. dataset.go stack_mine.go stack_mine_test.go
func BenchmarkSimplestMine10Values(b *testing.B) {
	stack := StackMine{}
	for _, num := range array10 {
		stack.Push(num)
	}
	for stack.len > 0 {
		stack.Pop()
	}
}

func BenchmarkSimplestMine1000Values(b *testing.B) {
	stack := StackMine{}
	for _, num := range array1000 {
		stack.Push(num)
	}
	for stack.len > 0 {
		stack.Pop()
	}
}

func BenchmarkSimplestMine100000Values(b *testing.B) {
	stack := StackMine{}
	for _, num := range array100000 {
		stack.Push(num)
	}
	for stack.len > 0 {
		stack.Pop()
	}
}
