package interfaces_p3

import (
	"fmt"
)

// ListItem - структура элемента двусвязного списка.
type ListItem struct {
	value          interface{}
	leftNeighbour  *ListItem
	rightNeighbour *ListItem
}

// LeftNeighbour() - метод получения стоящего слева элемента.
func (listItem *ListItem) LeftNeighbour() *ListItem {
	return listItem.leftNeighbour
}

// SetLeftNeighbour(item *ListItem) - метод присвоения стоящего слева элемента.
func (listItem *ListItem) SetLeftNeighbour(item *ListItem) {
	listItem.leftNeighbour = item
}

// RightNeighbour() - метод получения стоящего справа элемента.
func (listItem *ListItem) RightNeighbour() *ListItem {
	return listItem.rightNeighbour
}

// SetRightNeighbour(item *ListItem) - метод присвоения стоящего справа элемента.
func (listItem *ListItem) SetRightNeighbour(item *ListItem) {
	listItem.rightNeighbour = item
}

// Value() - метод получения значения из элемента.
func (listItem *ListItem) Value() interface{} {
	return listItem.value
}

// SetValue() - метод присвоения значения в элементе.
func (listItem *ListItem) SetValue(value interface{}) {
	listItem.value = value
}

// Eq(x, y ListItem) - элементы равны тогда и только тогда, когда это один и тот же элемент по памяти.
func Eq(x, y ListItem) bool {
	return &x == &y
}

// Lister - интерфейс двусвязного списка.
type Lister interface {
	Len() int
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

// Len() - метод получения длины двусвязного списка.
func (list *List) Len() int {
	return list.len
}

// LeftEdge() - метод получения элемента из левого края двусвязного списка.
func (list *List) LeftEdge() *ListItem {
	return list.leftEdge
}

// RightEdge() - метод получения элемента из правого края двусвязного списка.
func (list *List) RightEdge() *ListItem {
	return list.rightEdge
}

// PushToLeftEdge(value interface{}) - метод добавления значения в левый край двусвязного списка.
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

// PushToRightEdge(value interface{}) - метод добавления значения в правый край двусвязного списка.
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

// Contains(item *ListItem) - метод проверки наличия элемента в списке.
func (list *List) Contains(item *ListItem) bool {
	if (list.LeftEdge() == item) || // Это левый элемент
		(list.RightEdge() == item) || // Это правый элемент
		(item.LeftNeighbour() != nil && item.LeftNeighbour().RightNeighbour() == item &&
			item.RightNeighbour() != nil && item.RightNeighbour().LeftNeighbour() == item) { // Соседи ссылаются на него
		return true
	}
	return false
}

// Remove(item *ListItem) - метод удаления элемента из двусвязного списка.
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

// SwapItems(x, y *ListItem) - метод перестановки местами элементов двусвязного списка.
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

// MoveToFront(item *ListItem) - метод перемещения элемента в начало двусвязного списка.
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

// GetByIndex(i int) - метод получения i-того слева элемента двусвязного списка.
// Реализовано ИСКЛЮЧИТЕЛЬНО для демонстрации использования интерфейсных методов sort.Sort().
func (list *List) GetByIndex(i int) (*ListItem, error) {
	if i >= list.Len() {
		return nil, fmt.Errorf("index is out of range")
	}
	item := list.LeftEdge()
	for j := 0; j < i; j++ {
		item = item.RightNeighbour()
	}
	return item, nil
}

// Swap(i, j int) - метод перестановки i-того и j-того слева элементов двусвязного списка.
// Реализовано ИСКЛЮЧИТЕЛЬНО для демонстрации использования интерфейсных методов sort.Sort().
func (list *List) Swap(i, j int) {
	iItem, _ := list.GetByIndex(i)
	jItem, _ := list.GetByIndex(j)
	list.SwapItems(iItem, jItem)
}

// Less(i, j int) - метод сравнения убывания i-того и j-того слева элементов двусвязного списка.
// Реализовано ИСКЛЮЧИТЕЛЬНО для демонстрации использования интерфейсных методов sort.Sort()
// ПРИ УСЛОВИИ хранения INT-значений в поле value.
func (list *List) Less(i, j int) bool {
	iItem, _ := list.GetByIndex(i)
	jItem, _ := list.GetByIndex(j)
	return iItem.Value().(int) < jItem.Value().(int)
}

// NewList() - функция инициализации нового двусвязного списка.
func NewList() Lister {
	return new(List)
}
