package main

type StackPopedNil struct {
	items []int
	// Больше нет поля `len`
}

func (stack *StackPopedNil) Push(value int) {
	stack.items = append(stack.items, value)
}

func (stack *StackPopedNil) Pop() *int { // Изменился возвращаемый тип
	stack_items_len := len(stack.items)
	if stack_items_len == 0 {
		return nil // Возвращает nil указатель
	}
	filo := stack.items[stack_items_len-1]
	stack.items = stack.items[:stack_items_len-1]
	return &filo // Возвращает указатель на элемент
}
