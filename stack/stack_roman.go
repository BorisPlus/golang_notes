package main

// https://go.dev/play/p/_ZwCVgiT4Dt
type StackRoman struct {
	items []int
	len   uint
}

func (stack *StackRoman) Push(i int) {
	stack.items = append(stack.items, i)
	stack.len++
}

func (stack *StackRoman) Pop() int {
	stack_items_len := len(stack.items)
	filo := stack.items[stack_items_len-1]
	stack.items = stack.items[:stack_items_len-1]
	stack.len--
	return filo
}
