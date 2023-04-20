package main

import (
	"testing"
)

// go test -v -bench=. dataset.go stack_roman.go stack_roman_test.go
func BenchmarkSimplestRomanWithoutLen10Values(b *testing.B) {
	stack := StackRomanWithoutLen{}
	for _, num := range array10 {
		stack.Push(num)
	}
	var err error
	for err == nil {
		_, err = stack.Pop(); 
	}
}

func BenchmarkSimplestRomanWithoutLen1000Values(b *testing.B) {
	stack := StackRomanWithoutLen{}
	for _, num := range array1000 {
		stack.Push(num)
	}
	var err error
	for err == nil {
		_, err = stack.Pop(); 
	}
}

func BenchmarkSimplestRomanWithoutLen100000Values(b *testing.B) {
	stack := StackRomanWithoutLen{}
	for _, num := range array100000 {
		stack.Push(num)
	}
	var err error
	for err == nil {
		_, err = stack.Pop(); 
	}
}
