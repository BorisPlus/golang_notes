package main

type StackPointered struct {
	items []*int // Учет указателей
	len uint
}

func (stack *StackPointered) Push(value *int) { // Учет указателей
	stack.items = append(stack.items, value)
	stack.len++
}

func (stack *StackPointered) Pop() *int { // Возврат указателей
	stack_items_len := len(stack.items)
	filo := stack.items[stack_items_len-1]
	stack.items = stack.items[:stack_items_len-1]
	stack.len--
	return filo // Возврат указателя
}
