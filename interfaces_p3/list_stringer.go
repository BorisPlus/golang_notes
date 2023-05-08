package main

import (
	"fmt"
)

// String - наглядное представление значения элемента двусвязного списка.
//
// Например,
//
//	-------------------             -------------------
//	Item:  0xc00002e400             Item: 0xc00002e400
//	-------------------             -------------------
//	Value: 30                или    Value: 30
//	Left:  0xc00002e3c0             Left:  0x0
//	Right: 0xc00002e440             Right: 0x0
//	-------------------             -------------------
func (listItem *ListItem) String() string {
	template := `
-------------------
Item:  %p
-------------------
Value: %v
Left:  %p
Right: %p
-------------------`
	return fmt.Sprintf(template, listItem, listItem.Value(), listItem.Left(), listItem.Right())
}

// String - наглядное представление всего двусвязного списка.
//
// Например,
//
// - пустой список:
//
//	(nil:0x0)
//	   ^|
//	   L|R
//	    |v
//	(nil:0x0)
//
// - список из двух элементов:
//
//	 (nil:0x0)
//	    ^|
//	    L|R
//	     |v
//	-------------------
//	Item:  0xc00002e3a0 <--------┐
//	-------------------          |
//	Value: 2                     |
//	Left:  0x0                   |
//	Right: 0xc00002e380  >>>-----|---┐ Right 0xc00002e380
//	-------------------          |   | ссылается на
//	    ^|                       |   | блок 0xc00002e380
//	    L|R                      |   |
//	     |v                      |   |
//	-------------------          |   |
//	Item:  0xc00002e380  <-----------┘
//	-------------------          | Left 0xc00002e3a0
//	Value: 1                     | ссылается на
//	Left:  0xc00002e3a0  >>>-----┘ блок 0xc00002e3a0
//	Right: 0x0
//	-------------------
//	    ^|
//	    L|R
//	     |v
//	 (nil:0x0)
func (list *List) String() string {
	result := ""
	Nill := `
 (nil:0x0)`
	delimiter := `
    ^|
    L|R
     |v`
	result += Nill
	result += delimiter
	for i := list.LeftEnd(); i != nil; i = i.Right() {
		result += i.String()
		result += delimiter
	}
	result += Nill
	return result
}
