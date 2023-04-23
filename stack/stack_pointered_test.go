package main

import (
	"testing"
)

// go test -v dataset.go stack_pointed.go stack_pointed_test.go
func TestStackPointered10Values(t *testing.T) {
	stack := StackPointered{}
	for _, num := range array10 {
		stack.Push(&num)
	}

	// Будет замыкание
	// for i := stack.len; i > 0; {
	// 	fmt.Println(i)
	// 	stack.Pop()
	// }

	// Yе будет замыкания
	// for ;stack.len > 0; {
	for stack.len > 0 {
		// fmt.Println(stack.len)
		stack.Pop()
	}
}

// go test -v -bench=. dataset.go stack_pointered.go stack_pointered_test.go stack_roman.go stack_roman_10000000_test.go
func BenchmarkStackPointered10000000Values(b *testing.B) {
	stack := StackPointered{}
	for _, num := range array10000000 {
		stack.Push(&num)
	}
	for stack.len > 0 {
		stack.Pop()
	}
}
