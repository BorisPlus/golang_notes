package main

import (
	"fmt"
)

// ListItem - структура элемента двусвязного списка.
type ListItem struct {
	value          interface{}
	leftNeighbour  *ListItem
	rightNeighbour *ListItem
}

// LeftNeighbour() - получить стоящий слева элемент.
func (listItem *ListItem) LeftNeighbour() *ListItem {
	return listItem.leftNeighbour
}

// SetLeftNeighbour(item *ListItem) - установить стоящий слева элемент.
func (listItem *ListItem) SetLeftNeighbour(item *ListItem) {
	listItem.leftNeighbour = item
}

// RightNeighbour() - получить стоящий справа элемент.
func (listItem *ListItem) RightNeighbour() *ListItem {
	return listItem.rightNeighbour
}

// SetRightNeighbour(item *ListItem) - установить стоящий слева элемент.
func (listItem *ListItem) SetRightNeighbour(item *ListItem) {
	listItem.rightNeighbour = item
}

// Value() - получить значение из элемента.
func (listItem *ListItem) Value() interface{} {
	return listItem.value
}

// Eq(x, y ListItem) - элементы равны тогда и только тогда, когда это один и тот же элемент по памяти.
func Eq(x, y ListItem) bool {
	return &x == &y
}

// Lister - интерфейс двусвязного списка.
type Lister interface {
	Length() int
	LeftEdge() *ListItem
	RightEdge() *ListItem
	PushToLeftEdge(value interface{}) *ListItem
	PushToRightEdge(value interface{}) *ListItem
	Remove(item *ListItem) (*ListItem, error)
	SwapItems(x, y *ListItem) error
}

// List - структура двусвязного списка.
type List struct {
	len       int
	leftEdge  *ListItem
	rightEdge *ListItem
}

// Length() - получить длину двусвязного списка.
func (list *List) Length() int {
	return list.len
}

// LeftEdge() - получить элемент из левого края двусвязного списка.
func (list *List) LeftEdge() *ListItem {
	return list.leftEdge
}

// RightEdge() - получить элемент из правого края двусвязного списка.
func (list *List) RightEdge() *ListItem {
	return list.rightEdge
}

// PushToLeftEdge(value interface{}) - добавить значение в левый край двусвязного списка.
func (list *List) PushToLeftEdge(value interface{}) *ListItem {
	item := &ListItem{
		value:          value,
		leftNeighbour:  nil,
		rightNeighbour: nil,
	}
	if list.len == 0 {
		list.leftEdge = item
		list.rightEdge = item
	} else {
		item.rightNeighbour = list.LeftEdge()
		list.LeftEdge().leftNeighbour = item
		list.leftEdge = item
	}
	list.len++
	return item
}

// PushToRightEdge(value interface{}) - добавить значение в правый край двусвязного списка.
func (list *List) PushToRightEdge(value interface{}) *ListItem {
	item := &ListItem{
		value:          value,
		leftNeighbour:  nil,
		rightNeighbour: nil,
	}
	if list.len == 0 {
		list.leftEdge = item
		list.rightEdge = item
	} else {
		item.leftNeighbour = list.RightEdge()
		list.RightEdge().rightNeighbour = item
		list.rightEdge = item
	}
	list.len++
	return item
}

// Contains(item *ListItem) - проверить, есть ли элемент в списке.
func (list *List) Contains(item *ListItem) bool {
	if (list.LeftEdge() == item) || // Это левый элемент
		(list.RightEdge() == item) || // Это правый элемент
		(item.LeftNeighbour() != nil && item.LeftNeighbour().RightNeighbour() == item &&
			item.RightNeighbour() != nil && item.RightNeighbour().LeftNeighbour() == item) { // Соседи ссылаются на него
		return true
	}
	return false
}

// Remove(item *ListItem) - удалить элемент из двусвязного списка.
func (list *List) Remove(item *ListItem) (*ListItem, error) {
	if !list.Contains(item) {
		return nil, fmt.Errorf("it seems that item %s is not in the list", item)
	}
	if list.LeftEdge() == item && list.RightEdge() == item {
		list.leftEdge = nil
		list.rightEdge = nil
	} else if list.LeftEdge() == item {
		item.RightNeighbour().leftNeighbour = nil
		list.leftEdge = item.RightNeighbour()
	} else if list.RightEdge() == item {
		item.LeftNeighbour().rightNeighbour = nil
		list.rightEdge = item.LeftNeighbour()
	} else {
		item.LeftNeighbour().rightNeighbour = item.RightNeighbour()
		item.RightNeighbour().leftNeighbour = item.LeftNeighbour()
	}
	list.len--
	return item, nil
}

// SwapItems(x, y *ListItem) - поменять элементы двусвязного списка местами.
func (list *List) SwapItems(x, y *ListItem) error {
	if !list.Contains(x) {
		return fmt.Errorf("it seems that item %s is not in the list", x)
	}
	if !list.Contains(y) {
		return fmt.Errorf("it seems that item %s is not in the list", y)
	}
	if x == y {
		return nil
	}
	xLeft := x.LeftNeighbour()
	xRight := x.RightNeighbour()
	yLeft := y.LeftNeighbour()
	yRight := y.RightNeighbour()

	if xLeft != y && xRight != y {
		//
		y.rightNeighbour = xRight
		y.leftNeighbour = xLeft
		x.rightNeighbour = yRight
		x.leftNeighbour = yLeft
		//
		if yRight != nil {
			yRight.leftNeighbour = x
		} else {
			list.rightEdge = x
		}
		if yLeft != nil {
			yLeft.rightNeighbour = x
		} else {
			list.leftEdge = x
		}
		if xRight != nil {
			xRight.leftNeighbour = y
		} else {
			list.rightEdge = y
		}
		if xLeft != nil {
			xLeft.rightNeighbour = y
		} else {
			list.leftEdge = y
		}
		//
	} else if xRight == y {
		y.leftNeighbour = xLeft
		x.rightNeighbour = yRight
		y.rightNeighbour = x
		x.leftNeighbour = y
		if yRight != nil {
			yRight.leftNeighbour = x
		} else {
			list.rightEdge = x
		}
		if xLeft != nil {
			xLeft.rightNeighbour = y
		} else {
			list.leftEdge = y
		}
	} else if yRight == x {
		x.leftNeighbour = yLeft
		y.rightNeighbour = xRight
		x.rightNeighbour = y
		y.leftNeighbour = x
		if xRight != nil {
			xRight.leftNeighbour = y
		} else {
			list.rightEdge = y
		}
		if yLeft != nil {
			yLeft.rightNeighbour = x
		} else {
			list.leftEdge = x
		}
	}
	return nil
}

// MoveToFront(item *ListItem) - переместить элемент в начало двусвязного списка.
func (list *List) MoveToLeftEdge(item *ListItem) error {
	if !list.Contains(item) {
		return fmt.Errorf("it seems that item %s is not in the list", item)
	}
	if item.LeftNeighbour() != nil {
		item.LeftNeighbour().rightNeighbour = item.RightNeighbour()
	} else {
		list.leftEdge = item.RightNeighbour()
	}
	if item.RightNeighbour() != nil {
		item.RightNeighbour().leftNeighbour = item.LeftNeighbour()
	} else {
		list.rightEdge = item.LeftNeighbour()
	}

	item.leftNeighbour = nil
	item.rightNeighbour = list.LeftEdge()
	if list.LeftEdge() != nil {
		list.LeftEdge().leftNeighbour = item
		list.leftEdge = item
	} else {
		list.leftEdge = item
		list.rightEdge = item
	}
	return nil
}

func NewList() Lister {
	return new(List)
}
