package dlist

import (
	"fmt"
)

// DListItem - структура элемента двусвязного списка.
type DListItem struct {
	value          interface{}
	leftNeighbour  *DListItem
	rightNeighbour *DListItem
}

// LeftNeighbour() - метод получения стоящего слева элемента.
func (dlistItem *DListItem) LeftNeighbour() *DListItem {
	return dlistItem.leftNeighbour
}

// SetLeftNeighbour(item *DListItem) - метод присвоения стоящего слева элемента.
func (dlistItem *DListItem) SetLeftNeighbour(item *DListItem) {
	dlistItem.leftNeighbour = item
}

// RightNeighbour() - метод получения стоящего справа элемента.
func (dlistItem *DListItem) RightNeighbour() *DListItem {
	return dlistItem.rightNeighbour
}

// SetRightNeighbour(item *DListItem) - метод присвоения стоящего справа элемента.
func (dlistItem *DListItem) SetRightNeighbour(item *DListItem) {
	dlistItem.rightNeighbour = item
}

// Value() - метод получения значения из элемента.
func (dlistItem *DListItem) Value() interface{} {
	return dlistItem.value
}

// SetValue() - метод присвоения значения в элементе.
func (dlistItem *DListItem) SetValue(value interface{}) {
	dlistItem.value = value
}

// Eq(x, y DListItem) - элементы равны тогда и только тогда, когда это один и тот же элемент по памяти.
func Eq(x, y DListItem) bool {
	return &x == &y
}

// DLister - интерфейс двусвязного списка.
type DLister interface {
	Len() int
	LeftEdge() *DListItem
	RightEdge() *DListItem
	PushToLeftEdge(value interface{}) *DListItem
	PushToRightEdge(value interface{}) *DListItem
	Remove(item *DListItem) (*DListItem, error)
	SwapItems(x, y *DListItem) error
}

// DList - структура двусвязного списка.
type DList struct {
	len       int
	leftEdge  *DListItem
	rightEdge *DListItem
}

// Len() - метод получения длины двусвязного списка.
func (dlist *DList) Len() int {
	return dlist.len
}

// LeftEdge() - метод получения элемента из левого края двусвязного списка.
func (dlist *DList) LeftEdge() *DListItem {
	return dlist.leftEdge
}

// RightEdge() - метод получения элемента из правого края двусвязного списка.
func (dlist *DList) RightEdge() *DListItem {
	return dlist.rightEdge
}

// PushToLeftEdge(value interface{}) - метод добавления значения в левый край двусвязного списка.
func (dlist *DList) PushToLeftEdge(value interface{}) *DListItem {
	item := &DListItem{
		value:          value,
		leftNeighbour:  nil,
		rightNeighbour: nil,
	}
	if dlist.len == 0 {
		dlist.leftEdge = item
		dlist.rightEdge = item
	} else {
		item.rightNeighbour = dlist.LeftEdge()
		dlist.LeftEdge().leftNeighbour = item
		dlist.leftEdge = item
	}
	dlist.len++
	return item
}

// PushToRightEdge(value interface{}) - метод добавления значения в правый край двусвязного списка.
func (dlist *DList) PushToRightEdge(value interface{}) *DListItem {
	item := &DListItem{
		value:          value,
		leftNeighbour:  nil,
		rightNeighbour: nil,
	}
	if dlist.len == 0 {
		dlist.leftEdge = item
		dlist.rightEdge = item
	} else {
		item.leftNeighbour = dlist.RightEdge()
		dlist.RightEdge().rightNeighbour = item
		dlist.rightEdge = item
	}
	dlist.len++
	return item
}

// Contains(item *DListItem) - метод проверки наличия элемента в списке.
func (dlist *DList) Contains(item *DListItem) bool {
	if (dlist.LeftEdge() == item) || // Это левый элемент
		(dlist.RightEdge() == item) || // Это правый элемент
		(item.LeftNeighbour() != nil && item.LeftNeighbour().RightNeighbour() == item &&
			item.RightNeighbour() != nil && item.RightNeighbour().LeftNeighbour() == item) { // Соседи ссылаются на него
		return true
	}
	return false
}

// Remove(item *DListItem) - метод удаления элемента из двусвязного списка.
func (dlist *DList) Remove(item *DListItem) (*DListItem, error) {
	if !dlist.Contains(item) {
		return nil, fmt.Errorf("it seems that item %s is not in the dlist", item)
	}
	if dlist.LeftEdge() == item && dlist.RightEdge() == item {
		dlist.leftEdge = nil
		dlist.rightEdge = nil
	} else if dlist.LeftEdge() == item {
		item.RightNeighbour().leftNeighbour = nil
		dlist.leftEdge = item.RightNeighbour()
	} else if dlist.RightEdge() == item {
		item.LeftNeighbour().rightNeighbour = nil
		dlist.rightEdge = item.LeftNeighbour()
	} else {
		item.LeftNeighbour().rightNeighbour = item.RightNeighbour()
		item.RightNeighbour().leftNeighbour = item.LeftNeighbour()
	}
	dlist.len--
	return item, nil
}

// SwapItems(x, y *DListItem) - метод перестановки местами элементов двусвязного списка.
func (dlist *DList) SwapItems(x, y *DListItem) error {
	if !dlist.Contains(x) {
		return fmt.Errorf("it seems that item %s is not in the dlist", x)
	}
	if !dlist.Contains(y) {
		return fmt.Errorf("it seems that item %s is not in the dlist", y)
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
			dlist.rightEdge = x
		}
		if yLeft != nil {
			yLeft.rightNeighbour = x
		} else {
			dlist.leftEdge = x
		}
		if xRight != nil {
			xRight.leftNeighbour = y
		} else {
			dlist.rightEdge = y
		}
		if xLeft != nil {
			xLeft.rightNeighbour = y
		} else {
			dlist.leftEdge = y
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
			dlist.rightEdge = x
		}
		if xLeft != nil {
			xLeft.rightNeighbour = y
		} else {
			dlist.leftEdge = y
		}
	} else if yRight == x {
		x.leftNeighbour = yLeft
		y.rightNeighbour = xRight
		x.rightNeighbour = y
		y.leftNeighbour = x
		if xRight != nil {
			xRight.leftNeighbour = y
		} else {
			dlist.rightEdge = y
		}
		if yLeft != nil {
			yLeft.rightNeighbour = x
		} else {
			dlist.leftEdge = x
		}
	}
	return nil
}

// MoveToFront(item *DListItem) - метод перемещения элемента в начало двусвязного списка.
func (dlist *DList) MoveToLeftEdge(item *DListItem) error {
	if !dlist.Contains(item) {
		return fmt.Errorf("it seems that item %s is not in the dlist", item)
	}
	if item.LeftNeighbour() != nil {
		item.LeftNeighbour().rightNeighbour = item.RightNeighbour()
	} else {
		dlist.leftEdge = item.RightNeighbour()
	}
	if item.RightNeighbour() != nil {
		item.RightNeighbour().leftNeighbour = item.LeftNeighbour()
	} else {
		dlist.rightEdge = item.LeftNeighbour()
	}

	item.leftNeighbour = nil
	item.rightNeighbour = dlist.LeftEdge()
	if dlist.LeftEdge() != nil {
		dlist.LeftEdge().leftNeighbour = item
		dlist.leftEdge = item
	} else {
		dlist.leftEdge = item
		dlist.rightEdge = item
	}
	return nil
}

// GetByIndex(i int) - метод получения i-того слева элемента двусвязного списка.
// Реализовано ИСКЛЮЧИТЕЛЬНО для демонстрации использования интерфейсных методов sort.Sort().
func (dlist *DList) GetByIndex(i int) (*DListItem, error) {
	if i >= dlist.Len() {
		return nil, fmt.Errorf("index is out of range")
	}
	item := dlist.LeftEdge()
	for j := 0; j < i; j++ {
		item = item.RightNeighbour()
	}
	return item, nil
}

// Swap(i, j int) - метод перестановки i-того и j-того слева элементов двусвязного списка.
// Реализовано ИСКЛЮЧИТЕЛЬНО для демонстрации использования интерфейсных методов sort.Sort().
func (dlist *DList) Swap(i, j int) {
	iItem, _ := dlist.GetByIndex(i)
	jItem, _ := dlist.GetByIndex(j)
	dlist.SwapItems(iItem, jItem)
}

// Less(i, j int) - метод сравнения убывания i-того и j-того слева элементов двусвязного списка.
// Реализовано ИСКЛЮЧИТЕЛЬНО для демонстрации использования интерфейсных методов sort.Sort()
// ПРИ УСЛОВИИ хранения INT-значений в поле value.
func (dlist *DList) Less(i, j int) bool {
	iItem, _ := dlist.GetByIndex(i)
	jItem, _ := dlist.GetByIndex(j)
	return iItem.Value().(int) < jItem.Value().(int)
}

// NewList() - функция инициализации нового двусвязного списка.
func NewDList() DLister {
	return new(DList)
}
