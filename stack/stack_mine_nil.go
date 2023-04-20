package main

type StackMineNil struct {
	items []int
}

func (stack *StackMineNil) Push(i int) {
	stack.items = append(stack.items, i)
}

func (stack *StackMineNil) Pop() (filo int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
			filo = 0
		}
	}()
	stack_items_len := len(stack.items)
	filo = stack.items[stack_items_len-1]
	stack.items = stack.items[:stack_items_len-1]
	return
}
