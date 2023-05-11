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
func (dList *DList) Len() int {
	return dList.len
}

// LeftEdge() - метод получения элемента из левого края двусвязного списка.
func (dList *DList) LeftEdge() *DListItem {
	return dList.leftEdge
}

// RightEdge() - метод получения элемента из правого края двусвязного списка.
func (dList *DList) RightEdge() *DListItem {
	return dList.rightEdge
}

// PushToLeftEdge(value interface{}) - метод добавления значения в левый край двусвязного списка.
func (dList *DList) PushToLeftEdge(value interface{}) *DListItem {
	item := &DListItem{
		value:          value,
		leftNeighbour:  nil,
		rightNeighbour: nil,
	}
	if dList.len == 0 {
		dList.leftEdge = item
		dList.rightEdge = item
	} else {
		item.rightNeighbour = dList.LeftEdge()
		dList.LeftEdge().leftNeighbour = item
		dList.leftEdge = item
	}
	dList.len++
	return item
}

// PushToRightEdge(value interface{}) - метод добавления значения в правый край двусвязного списка.
func (dList *DList) PushToRightEdge(value interface{}) *DListItem {
	item := &DListItem{
		value:          value,
		leftNeighbour:  nil,
		rightNeighbour: nil,
	}
	if dList.len == 0 {
		dList.leftEdge = item
		dList.rightEdge = item
	} else {
		item.leftNeighbour = dList.RightEdge()
		dList.RightEdge().rightNeighbour = item
		dList.rightEdge = item
	}
	dList.len++
	return item
}

// Contains(item *DListItem) - метод проверки наличия элемента в списке.
func (dList *DList) Contains(item *DListItem) bool {
	if (dList.LeftEdge() == item) || // Это левый элемент
		(dList.RightEdge() == item) || // Это правый элемент
		(item.LeftNeighbour() != nil && item.LeftNeighbour().RightNeighbour() == item &&
			item.RightNeighbour() != nil && item.RightNeighbour().LeftNeighbour() == item) { // Соседи ссылаются на него
		return true
	}
	return false
}

// Remove(item *DListItem) - метод удаления элемента из двусвязного списка.
func (dList *DList) Remove(item *DListItem) (*DListItem, error) {
	if !dList.Contains(item) {
		return nil, fmt.Errorf("it seems that item %s is not in the dList", item)
	}
	if dList.LeftEdge() == item && dList.RightEdge() == item {
		dList.leftEdge = nil
		dList.rightEdge = nil
	} else if dList.LeftEdge() == item {
		item.RightNeighbour().leftNeighbour = nil
		dList.leftEdge = item.RightNeighbour()
	} else if dList.RightEdge() == item {
		item.LeftNeighbour().rightNeighbour = nil
		dList.rightEdge = item.LeftNeighbour()
	} else {
		item.LeftNeighbour().rightNeighbour = item.RightNeighbour()
		item.RightNeighbour().leftNeighbour = item.LeftNeighbour()
	}
	dList.len--
	return item, nil
}

// SwapItems(x, y *DListItem) - метод перестановки местами элементов двусвязного списка.
func (dList *DList) SwapItems(x, y *DListItem) error {
	if !dList.Contains(x) {
		return fmt.Errorf("it seems that item %s is not in the dList", x)
	}
	if !dList.Contains(y) {
		return fmt.Errorf("it seems that item %s is not in the dList", y)
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
			dList.rightEdge = x
		}
		if yLeft != nil {
			yLeft.rightNeighbour = x
		} else {
			dList.leftEdge = x
		}
		if xRight != nil {
			xRight.leftNeighbour = y
		} else {
			dList.rightEdge = y
		}
		if xLeft != nil {
			xLeft.rightNeighbour = y
		} else {
			dList.leftEdge = y
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
			dList.rightEdge = x
		}
		if xLeft != nil {
			xLeft.rightNeighbour = y
		} else {
			dList.leftEdge = y
		}
	} else if yRight == x {
		x.leftNeighbour = yLeft
		y.rightNeighbour = xRight
		x.rightNeighbour = y
		y.leftNeighbour = x
		if xRight != nil {
			xRight.leftNeighbour = y
		} else {
			dList.rightEdge = y
		}
		if yLeft != nil {
			yLeft.rightNeighbour = x
		} else {
			dList.leftEdge = x
		}
	}
	return nil
}

// MoveToFront(item *DListItem) - метод перемещения элемента в начало двусвязного списка.
func (dList *DList) MoveToLeftEdge(item *DListItem) error {
	if !dList.Contains(item) {
		return fmt.Errorf("it seems that item %s is not in the dList", item)
	}
	if item.LeftNeighbour() != nil {
		item.LeftNeighbour().rightNeighbour = item.RightNeighbour()
	} else {
		dList.leftEdge = item.RightNeighbour()
	}
	if item.RightNeighbour() != nil {
		item.RightNeighbour().leftNeighbour = item.LeftNeighbour()
	} else {
		dList.rightEdge = item.LeftNeighbour()
	}

	item.leftNeighbour = nil
	item.rightNeighbour = dList.LeftEdge()
	if dList.LeftEdge() != nil {
		dList.LeftEdge().leftNeighbour = item
		dList.leftEdge = item
	} else {
		dList.leftEdge = item
		dList.rightEdge = item
	}
	return nil
}

// GetByIndex(i int) - метод получения i-того слева элемента двусвязного списка.
// Реализовано ИСКЛЮЧИТЕЛЬНО для демонстрации использования интерфейсных методов sort.Sort().
func (dList *DList) GetByIndex(i int) (*DListItem, error) {
	if i >= dList.Len() {
		return nil, fmt.Errorf("index is out of range")
	}
	item := dList.LeftEdge()
	for j := 0; j < i; j++ {
		item = item.RightNeighbour()
	}
	return item, nil
}

// Swap(i, j int) - метод перестановки i-того и j-того слева элементов двусвязного списка.
// Реализовано ИСКЛЮЧИТЕЛЬНО для демонстрации использования интерфейсных методов sort.Sort().
func (dList *DList) Swap(i, j int) {
	iItem, _ := dList.GetByIndex(i)
	jItem, _ := dList.GetByIndex(j)
	dList.SwapItems(iItem, jItem)
}

// Less(i, j int) - метод сравнения убывания i-того и j-того слева элементов двусвязного списка.
// Реализовано ИСКЛЮЧИТЕЛЬНО для демонстрации использования интерфейсных методов sort.Sort()
// ПРИ УСЛОВИИ хранения INT-значений в поле value.
func (dList *DList) Less(i, j int) bool {
	iItem, _ := dList.GetByIndex(i)
	jItem, _ := dList.GetByIndex(j)
	return iItem.Value().(int) < jItem.Value().(int)
}

// NewList() - функция инициализации нового двусвязного списка.
func NewDList() DLister {
	return new(DList)
}
