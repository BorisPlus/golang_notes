package main

type stackPointedSpecial struct {
	items []*int
}

func CreateStackPointedSpecial() stackPointedSpecial {
	stack := stackPointedSpecial{}
	stack.Push(nil)
	return stack
}

func (stack *stackPointedSpecial) Push(value *int) {
	stack.items = append(stack.items, value)
}

func (stack *stackPointedSpecial) Pop() *int {
	stack_items_len := len(stack.items)
	filo := stack.items[stack_items_len-1]
	if filo == nil {
		return nil
	}
	stack.items = stack.items[:stack_items_len-1]
	return filo
}
