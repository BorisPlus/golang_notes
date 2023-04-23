package main

type Item struct {
	value int
	next  *Item
}

type StackStructed struct {
	origin *Item
}

func (stack *StackStructed) Push(value int) {
	willBePushed := Item{value: value, next: stack.origin}
	stack.origin = &willBePushed
}

func (stack *StackStructed) Pop() (int, *Item) { //
	poped := stack.origin.value
	stack.origin = stack.origin.next
	return poped, stack.origin 
}
