package main

// https://go.dev/play/p/j2AVLXLh4Tt
type StackKonstantin struct {
	items []int
	len   uint
}

func (stack *StackKonstantin) Push(i int) {
	stack.items = append([]int{i}, stack.items...)
	stack.len++
}

func (stack *StackKonstantin) Pop() int {
	filo := stack.items[0]
	stack.items = stack.items[1:]
	stack.len--
	return filo
}
