package main

import (
	"testing"
)

// go test -v -bench=. dataset.go stack_poped_nil.go stack_poped_nil_test.go stack_roman.go stack_roman_10000000_test.go
func BenchmarkStackPopedNil10000000Values(b *testing.B) {
	stack := StackPopedNil{}
	for _, num := range array10000000 {
		stack.Push(num)
	}
	for stack.Pop() != nil {
	}
}
