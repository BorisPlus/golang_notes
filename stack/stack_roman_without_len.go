package main

type StackRomanWithoutLen struct {
	items []int
}

func (stack *StackRomanWithoutLen) Push(i int) {
	stack.items = append(stack.items, i)
}

func (stack *StackRomanWithoutLen) Pop() (filo int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()
	stack_items_len := len(stack.items)
	filo = stack.items[stack_items_len-1]
	stack.items = stack.items[:stack_items_len-1]
	return
}
