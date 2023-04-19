package main

import (
	"testing"
)

// go test -v -bench=. dataset.go stack_konstantin.go stack_konstantin_test.go
func BenchmarkSimplestKonstantin10Values(b *testing.B) {

	stack := StackKonstantin{}
	for num := range array10 {
		stack.Push(num)
	}
	for stack.len > 0 {
		stack.Pop()
	}
}

func BenchmarkSimplestKonstantin1000Values(b *testing.B) {

	stack := StackKonstantin{}
	for num := range array1000 {
		stack.Push(num)
	}
	for stack.len > 0 {
		stack.Pop()
	}
}
func BenchmarkSimplestKonstantin100000Values(b *testing.B) {

	stack := StackKonstantin{}
	for num := range array100000 {
		stack.Push(num)
	}
	for stack.len > 0 {
		stack.Pop()
	}
}
