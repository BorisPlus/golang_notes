package main

type StackMine struct {
	items []int
	len uint
}

func (stack *StackMine) Push(i int) {
	stack.items = append(stack.items, i)
	stack.len++
}

func (stack *StackMine) Pop() int {
	filo := stack.items[0]
	stack.items = stack.items[1:]
	stack.len--
	return filo
}

