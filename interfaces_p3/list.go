package main

import (
	"fmt"
)

// ListItem - структура элемента двусвязного списка.
type ListItem struct {
	value interface{}
	left  *ListItem
	right *ListItem
}

// Left() - получить стоящий слева элемент.
func (listItem *ListItem) Left() *ListItem {
	return listItem.left
}

// Right() - получить стоящий справа элемент.
func (listItem *ListItem) Right() *ListItem {
	return listItem.right
}

// Value() - получить значение из элемента.
func (listItem *ListItem) Value() interface{} {
	return listItem.value
}

// Eq(x, y ListItem) - элементs равны тогда и только тогда, когда это один и тот же элемент.
func Eq(x, y ListItem) bool {
	return &x == &y
}

// Lister - интерфейс двусвязного списка.
type Lister interface {
	Len() int
	LeftEnd() *ListItem
	RightEnd() *ListItem
	PushToLeftEnd(value interface{}) *ListItem
	PushToRightEnd(value interface{}) *ListItem
	Remove(item *ListItem) (*ListItem, error)
	Swap(x, y *ListItem) error
}

// List - структура двусвязного списка.
type List struct {
	len      int
	leftEnd  *ListItem
	rightEnd *ListItem
}

// Len() - получить длину двусвязного списка.
func (list *List) Len() int {
	return list.len
}

// LeftEnd() - получить элемент из левого края двусвязного списка.
func (list *List) LeftEnd() *ListItem {
	return list.leftEnd
}

// RightEnd() - получить элемент из правого края двусвязного списка.
func (list *List) RightEnd() *ListItem {
	return list.rightEnd
}

// PushToLeftEnd() - добавить значение в левый край двусвязного списка.
func (list *List) PushToLeftEnd(value interface{}) *ListItem {
	item := &ListItem{
		value: value,
		left:  nil,
		right: nil,
	}
	if list.len == 0 {
		list.leftEnd = item
		list.rightEnd = item
	} else {
		item.right = list.leftEnd
		list.leftEnd.left = item
		list.leftEnd = item
	}
	list.len++
	return item
}

// PushToRightEnd(any interface{}) - добавить значение в правый край двусвязного списка.
func (list *List) PushToRightEnd(value interface{}) *ListItem {
	item := &ListItem{
		value: value,
		left:  nil,
		right: nil,
	}
	if list.len == 0 {
		list.leftEnd = item
		list.rightEnd = item
	} else {
		item.left = list.rightEnd
		list.rightEnd.right = item
		list.rightEnd = item
	}
	list.len++
	return item
}

// Contains(item *ListItem) - проверить, есть ли элемент в списке.
func (list *List) Contains(item *ListItem) bool {
	if (list.leftEnd == item) || // Это левый элемент
		(list.rightEnd == item) || // Это правый элемент
		(item.left != nil && item.left.right == item && item.right != nil && item.right.left == item) { // Соседи ссылаются на него
		return true
	}
	return false
}

// Remove(item *ListItem) - удалить элемент из двусвязного списка.
func (list *List) Remove(item *ListItem) (*ListItem, error) {
	if !list.Contains(item) {
		return nil, fmt.Errorf("it seems that item %s is not in the list", item)
	}
	if list.leftEnd == item && list.rightEnd == item {
		list.leftEnd = nil
		list.rightEnd = nil
	} else if list.leftEnd == item {
		item.right.left = nil
		list.leftEnd = item.right
	} else if list.rightEnd == item {
		item.left.right = nil
		list.rightEnd = item.left
	} else {
		item.left.right = item.right
		item.right.left = item.left
	}
	list.len--
	return item, nil
}

func (list *List) Swap(x, y *ListItem) error {
	if !list.Contains(x) {
		return fmt.Errorf("it seems that item %s is not in the list", x)
	}
	if !list.Contains(y) {
		return fmt.Errorf("it seems that item %s is not in the list", y)
	}
	if x == y {
		return nil
	}
	xLeft := x.left
	xRight := x.right
	yLeft := y.left
	yRight := y.right

	if xLeft != y && xRight != y {
		//
		y.right = xRight
		y.left = xLeft
		x.right = yRight
		x.left = yLeft
		//
		if yRight != nil {
			yRight.left = x
		} else {
			list.rightEnd = x
		}
		if yLeft != nil {
			yLeft.right = x
		} else {
			list.leftEnd = x
		}
		if xRight != nil {
			xRight.left = y
		} else {
			list.rightEnd = y
		}
		if xLeft != nil {
			xLeft.right = y
		} else {
			list.leftEnd = y
		}
		//
	} else if xRight == y {
		y.left = xLeft
		x.right = yRight
		y.right = x
		x.left = y
		if yRight != nil {
			yRight.left = x
		} else {
			list.rightEnd = x
		}
		if xLeft != nil {
			xLeft.right = y
		} else {
			list.leftEnd = y
		}
	} else if yRight == x {
		x.left = yLeft
		y.right = xRight
		x.right = y
		y.left = x
		if xRight != nil {
			xRight.left = y
		} else {
			list.rightEnd = y
		}
		if yLeft != nil {
			yLeft.right = x
		} else {
			list.leftEnd = x
		}
	}

	// y.right = xRight.right.left
	// x.left = yLeft.left.right

	// if xRight.right != nil {
	// 	y.right = xRight.right.left
	// 	y.left = xLeft
	// } else {
	// 	y.right = xRight
	// 	y.left = xLeft.left.right
	// }

	// if yLeft.left != nil {
	// 	x.left = yLeft.left.right
	// 	x.right = yRight
	// } else {
	// 	x.left = yLeft
	// 	x.right = yRight.right.left
	// }

	// if yLeft.left != nil{
	// 	x.left = yLeft.left.right
	// } else if yLeft.right != nil {
	// 	x.left = yLeft.right.left
	// }

	// xLeft.right = y
	// yRight.left = x
	// if x.left != nil {
	// 	x.left.right = y
	// }
	// x.right = y.right

	// if y.right != nil {
	// 	y.right.left = x
	// }
	// yLeft := y.left
	// y.left = x.left

	// if xRight != nil {
	// 	xRight.left = y
	// }

	// if xLeft != nil {
	// 	xLeft.right = y
	// }

	// if yLeft != nil {
	// 	yLeft.right = x
	// }
	// y.right = xRight

	// if x.left != nil {
	// 	x.left.right = y
	// }
	// x.right = y.right
	// xLeft, xRight, yLeft, yRight := , x.right, y.left, y.right
	// x.left, x.right, y.left, y.right = yLeft, yRight, xLeft, xRight
	return nil
}

func NewList() Lister {
	return new(List)
}
