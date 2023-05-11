# Интерфейсы. Двусвязный список

## Реализация

Интерфейсы и структуры:

<details>
<summary>см. "dlist.go":</summary>

```go
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

```

</details>

Наглядность:

<details>
<summary>см. "dlist_stringer.go":</summary>

```go
package dlist

import (
    "fmt"
)

// String() - представление элемента двусвязного списка.
//
// Например,
//
//    -------------------             -------------------
//     Item: 0xc00002e400              Item: 0xc00002e400
//    -------------------             -------------------
//    Value: 30                или    Value: 30
//     Left: 0xc00002e3c0              Left: 0x0
//    Right: 0xc00002e440             Right: 0x0
//    -------------------             -------------------
func (dlistItem *DListItem) String() string {
    template := `
-------------------
 Item: %p
-------------------
value: %v
 left: %p
right: %p
-------------------`
    return fmt.Sprintf(template, dlistItem, dlistItem.Value(), dlistItem.LeftNeighbour(), dlistItem.RightNeighbour())
}

// String() - представление двусвязного списка.
//
// Например,
//
// - пустой список:
//
//    (nil:0x0)
//       ^|
//       L|R
//        |v
//    (nil:0x0)
//
// - список из двух элементов:
//
//     (nil:0x0)
//        ^|
//        L|R
//         |v
//    -------------------
//     Item: 0xc00002e3a0 <--------┐
//    -------------------          |
//    Value: 2                     |
//     Left: 0x0                   |
//    Right: 0xc00002e380  >>>-----|---┐ Right 0xc00002e380
//    -------------------          |   | ссылается на
//        ^|                       |   | блок 0xc00002e380
//        L|R                      |   |
//         |v                      |   |
//    -------------------          |   |
//     Item: 0xc00002e380  <-----------┘
//    -------------------          | Left 0xc00002e3a0
//    Value: 1                     | ссылается на
//     Left: 0xc00002e3a0  >>>-----┘ блок 0xc00002e3a0
//    Right: 0x0
//    -------------------
//        ^|
//        L|R
//         |v
//     (nil:0x0)
func (dlist *DList) String() string {
    result := ""
    Nill := `
 (nil:0x0)`
    delimiter := `
    ^|
    L|R
     |v`
    result += Nill
    result += delimiter
    for i := dlist.LeftEdge(); i != nil; i = i.RightNeighbour() {
        result += i.String()
        result += delimiter
    }
    result += Nill
    return result
}

```

</details>

Тестирование:

<details>
<summary>см. "dlist_test.go":</summary>

```go
package dlist_test

import (
    "fmt"
    "sort"
    "testing"

    "github.com/stretchr/testify/require"

    dlist "github.com/BorisPlus/golang_notes/dlist" 
)

// go test -v dList.go list_string.go list_test.go

func ElementsLeftToRight(dList *dlist.DList) []int {
    elements := make([]int, 0, dList.Len())
    for i := dList.LeftEdge(); i != nil; i = i.RightNeighbour() {
        elements = append(elements, i.Value().(int))
    }
    return elements
}

func ElementsRightToLeft(dList *dlist.DList) []int {
    elements := make([]int, 0, dList.Len())
    for i := dList.RightEdge(); i != nil; i = i.LeftNeighbour() {
        elements = append([]int{i.Value().(int)}, elements...)
    }
    return elements
}

func TestDListSimple(t *testing.T) {
    t.Run("Zero-value DListItem test.", func(t *testing.T) {
        zeroValueItem := dlist.DListItem{}
        require.Nil(t, zeroValueItem.Value())
        require.Nil(t, zeroValueItem.LeftNeighbour())
        require.Nil(t, zeroValueItem.RightNeighbour())
        fmt.Println("\n[zero-value] is:", &zeroValueItem)
        fmt.Println()
    })

    t.Run("DListItem direct and vice versa referencies test.", func(t *testing.T) {
        fmt.Println("\n[1] <--> [2] <--> [3]")
        first := &dlist.DListItem{}
        first.SetValue(1)
        fmt.Println("\n[1] is:", first)

        second := &dlist.DListItem{}
        second.SetValue(2)
        second.SetLeftNeighbour(first)
        first.SetRightNeighbour(second)
        fmt.Println("\nadd [2]:", second)
        fmt.Println("\n[1] become:", first)

        third := dlist.DListItem{}
        third.SetValue(3)
        third.SetLeftNeighbour(second)
        second.SetRightNeighbour(&third)
        fmt.Println("\nadd [3]:", &third)
        fmt.Println("\n[2] become:", second)
        fmt.Println()

        fmt.Println("first.RightNeighbour().RightNeighbour().RightNeighbour() is nil. OK.")
        require.Nil(t, first.RightNeighbour().RightNeighbour().RightNeighbour())
        fmt.Println("first.RightNeighbour().RightNeighbour() is third. OK.")
        require.Equal(t, &third, first.RightNeighbour().RightNeighbour())
        fmt.Println("third.LeftNeighbour().LeftNeighbour().LeftNeighbour() is nil. OK.")
        require.Nil(t, third.LeftNeighbour().LeftNeighbour().LeftNeighbour())
        fmt.Println("third.LeftNeighbour().LeftNeighbour() is first. OK.")
        require.Equal(t, first, third.LeftNeighbour().LeftNeighbour())
        fmt.Println()
    })

    t.Run("Empty DList test.", func(t *testing.T) {
        dList := dlist.NewDList()
        require.Equal(t, 0, dList.Len())
        require.Nil(t, dList.LeftEdge())
        require.Nil(t, dList.RightEdge())
        fmt.Println("\nList is:\n", dList)
        fmt.Println()
    })

    t.Run("DList init test.", func(t *testing.T) {
        dList := dlist.NewDList()

        fmt.Println("\nList was:\n", dList)

        item := dList.PushToLeftEdge(1)
        fmt.Println("\nItem was:", item)
        require.Equal(t, item, dList.LeftEdge())
        require.Equal(t, item, dList.RightEdge())
        fmt.Println()
        fmt.Println("dList.LeftEdge() and dList.RightEdge() is item. OK.")
        fmt.Println("\nList become:\n", dList)
        fmt.Println()
    })

    t.Run("Little DList test.", func(t *testing.T) {
        dList := dlist.NewDList()
        fmt.Println("\nList was:\n", dList)

        itemFirst := dList.PushToLeftEdge(1) // [1]
        fmt.Println("\nItem [1] become:\n", itemFirst)

        leftEnd := dList.LeftEdge()
        fmt.Println("\nlist.LeftEdge() become:\n", leftEnd)
        rightEnd := dList.RightEdge()
        fmt.Println("\nlist.RightEdge() become:\n", rightEnd)

        itemSecond := dList.PushToLeftEdge(2) // [2]
        fmt.Println("\nItem [2] become:\n", itemSecond)

        leftEnd = dList.LeftEdge()
        fmt.Println("\nlist.LeftEdge() become:\n", leftEnd)
        rightEnd = dList.RightEdge()
        fmt.Println("\nlist.RightEdge() become:\n", rightEnd)
        fmt.Println("\nList become:\n", dList)
        require.Equal(t, itemSecond, dList.LeftEdge())
        require.Equal(t, itemFirst, dList.RightEdge())

        someItem, _ := dList.Remove(dList.LeftEdge())
        fmt.Println("\nWas removed:\n", someItem)
        require.Equal(t, itemSecond, someItem)

        fmt.Println("\nList become:\n", dList)

        require.Equal(t, dList.LeftEdge(), dList.RightEdge())
        require.Nil(t, dList.LeftEdge().LeftNeighbour())
        require.Nil(t, dList.LeftEdge().RightNeighbour())

        someItem, _ = dList.Remove(dList.LeftEdge())
        fmt.Println("\nWas removed:\n", someItem)
        require.Equal(t, itemFirst, someItem)

        fmt.Println("\nList become:\n", dList)

        require.Equal(t, dList.LeftEdge(), dList.RightEdge())
        require.Nil(t, dList.LeftEdge())
        require.Nil(t, dList.RightEdge())
    })
}

func TestDListComplex(t *testing.T) {
    t.Run("Сomplex DList processing test.", func(t *testing.T) {
        dList := dlist.NewDList()
        dList.PushToRightEdge(10) // [10]
        dList.PushToRightEdge(20) // [10, 20]
        dList.PushToRightEdge(30) // [10, 20, 30]
        require.Equal(t, []int{10, 20, 30}, ElementsLeftToRight(dList.(*dlist.DList)))
        require.Equal(t, []int{10, 20, 30}, ElementsRightToLeft(dList.(*dlist.DList)))
        fmt.Println("Forward stroke check for [10, 20, 30]. OK.")
        require.Equal(t, 3, dList.Len())
        middle := dList.LeftEdge().RightNeighbour()
        require.Equal(t, middle.Value(), 20)
        fmt.Printf("middle.Value() is %v. OK.\n", middle.Value())
        dList.Remove(middle)
        fmt.Printf("middle was removed. OK.\n")
        require.Equal(t, 2, dList.Len())
        require.Equal(t, []int{10, 30}, ElementsLeftToRight(dList.(*dlist.DList)))
        require.Equal(t, []int{10, 30}, ElementsRightToLeft(dList.(*dlist.DList)))
        fmt.Println("Forward stroke check for [10, 30]. OK.")
        for i, v := range [...]int{40, 50, 60, 70, 80} {
            if i%2 == 0 {
                dList.PushToLeftEdge(v)
            } else {
                dList.PushToRightEdge(v)
            }
        } // [80, 60, 40, 10, 30, 50, 70]
        fmt.Println("DList [10, 30] mixing values {40, 50, 60, 70, 80} with mod(2, index).")

        require.Equal(t, 7, dList.Len())
        fmt.Printf("dList.Len() is %v. OK.\n", dList.Len())
        require.Equal(t, 80, dList.LeftEdge().Value())
        fmt.Printf("dList.LeftEdge().Value() is %v. OK.\n", dList.LeftEdge().Value())
        require.Equal(t, 70, dList.RightEdge().Value())
        fmt.Printf("dList.RightEdge().Value() is %v. OK.\n", dList.RightEdge().Value())

        require.Equal(t, []int{80, 60, 40, 10, 30, 50, 70}, ElementsLeftToRight(dList.(*dlist.DList)))
        require.Equal(t, []int{80, 60, 40, 10, 30, 50, 70}, ElementsRightToLeft(dList.(*dlist.DList)))
        fmt.Println("Forward stroke check for [80, 60, 40, 10, 30, 50, 70]. OK.")

        rightEnd := dList.RightEdge()
        dList.Remove(rightEnd)
        dList.PushToRightEdge(rightEnd.Value())
        fmt.Println("Remove right and push back to right - check for [80, 60, 40, 10, 30, 50, 70]. OK.")
        leftEnd, _ := dList.Remove(dList.LeftEdge())
        dList.PushToLeftEdge(leftEnd.Value())
        fmt.Println("Remove left and push back to left - check for [80, 60, 40, 10, 30, 50, 70]. OK.")
        // - check for nil-refs of first and last
        require.Equal(t, dList.LeftEdge().Value(), 80)
        require.Equal(t, dList.RightEdge().Value(), 70)
        fmt.Println("Check for dList.LeftEdge() is 80 and dList.RightEdge() is 70. OK.")
        require.Nil(t, dList.LeftEdge().LeftNeighbour())
        require.Nil(t, dList.RightEdge().RightNeighbour())
        fmt.Println("Check for dList.LeftEdge().Left() and dList.RightEdge().Right() is nils. OK.")

    })
}

func TestDListSwap(t *testing.T) {
    t.Run("DList swap test.", func(t *testing.T) {
        dList := dlist.NewDList()
        one := dList.PushToRightEdge(10)   // [10]
        two := dList.PushToRightEdge(20)   // [10, 20]
        three := dList.PushToRightEdge(30) // [10, 20, 30]
        four := dList.PushToRightEdge(40)  // [10, 20, 30, 40]
        five := dList.PushToRightEdge(50)  // [10, 20, 30, 40, 50]
        require.Equal(t, []int{10, 20, 30, 40, 50}, ElementsLeftToRight(dList.(*dlist.DList)))
        require.Equal(t, []int{10, 20, 30, 40, 50}, ElementsRightToLeft(dList.(*dlist.DList)))
        fmt.Println("[10, 20, 30, 40, 50]. OK.")
        dList.SwapItems(two, four)
        fmt.Println("swap different element-pairs")
        require.Equal(t, []int{10, 40, 30, 20, 50}, ElementsLeftToRight(dList.(*dlist.DList)))
        require.Equal(t, []int{10, 40, 30, 20, 50}, ElementsRightToLeft(dList.(*dlist.DList)))
        fmt.Println("[10, 40, 30, 20, 50]. OK.")
        dList.SwapItems(two, four)
        require.Equal(t, []int{10, 20, 30, 40, 50}, ElementsLeftToRight(dList.(*dlist.DList)))
        require.Equal(t, []int{10, 20, 30, 40, 50}, ElementsRightToLeft(dList.(*dlist.DList)))
        fmt.Println("[10, 20, 30, 40, 50]. OK.")
        dList.SwapItems(one, five)
        require.Equal(t, []int{50, 20, 30, 40, 10}, ElementsLeftToRight(dList.(*dlist.DList)))
        require.Equal(t, []int{50, 20, 30, 40, 10}, ElementsRightToLeft(dList.(*dlist.DList)))
        fmt.Println("[50, 20, 30, 40, 10]. OK.")
        dList.SwapItems(one, five)
        require.Equal(t, []int{10, 20, 30, 40, 50}, ElementsLeftToRight(dList.(*dlist.DList)))
        require.Equal(t, []int{10, 20, 30, 40, 50}, ElementsRightToLeft(dList.(*dlist.DList)))
        fmt.Println("[10, 20, 30, 40, 50]. OK.")
        dList.SwapItems(one, three)
        require.Equal(t, []int{30, 20, 10, 40, 50}, ElementsLeftToRight(dList.(*dlist.DList)))
        require.Equal(t, []int{30, 20, 10, 40, 50}, ElementsRightToLeft(dList.(*dlist.DList)))
        fmt.Println("[30, 20, 10, 40, 50]. OK.")
        dList.SwapItems(one, three)
        require.Equal(t, []int{10, 20, 30, 40, 50}, ElementsLeftToRight(dList.(*dlist.DList)))
        require.Equal(t, []int{10, 20, 30, 40, 50}, ElementsRightToLeft(dList.(*dlist.DList)))
        fmt.Println("[10, 20, 30, 40, 50]. OK.")
        dList.SwapItems(five, two)
        require.Equal(t, []int{10, 50, 30, 40, 20}, ElementsLeftToRight(dList.(*dlist.DList)))
        require.Equal(t, []int{10, 50, 30, 40, 20}, ElementsRightToLeft(dList.(*dlist.DList)))
        fmt.Println("[10, 50, 30, 40, 20]. OK.")
        dList.SwapItems(two, five)
        require.Equal(t, []int{10, 20, 30, 40, 50}, ElementsLeftToRight(dList.(*dlist.DList)))
        require.Equal(t, []int{10, 20, 30, 40, 50}, ElementsRightToLeft(dList.(*dlist.DList)))
        fmt.Println("[10, 20, 30, 40, 50]. OK.")
        dList.SwapItems(two, three)
        require.Equal(t, []int{10, 30, 20, 40, 50}, ElementsLeftToRight(dList.(*dlist.DList)))
        require.Equal(t, []int{10, 30, 20, 40, 50}, ElementsRightToLeft(dList.(*dlist.DList)))
        fmt.Println("[10, 30, 20, 40, 50]. OK.")
        dList.SwapItems(two, three)
        require.Equal(t, []int{10, 20, 30, 40, 50}, ElementsLeftToRight(dList.(*dlist.DList)))
        require.Equal(t, []int{10, 20, 30, 40, 50}, ElementsRightToLeft(dList.(*dlist.DList)))
        fmt.Println("[10, 20, 30, 40, 50]. OK.")
        dList.SwapItems(four, three)
        require.Equal(t, []int{10, 20, 40, 30, 50}, ElementsLeftToRight(dList.(*dlist.DList)))
        require.Equal(t, []int{10, 20, 40, 30, 50}, ElementsRightToLeft(dList.(*dlist.DList)))
        fmt.Println("[10, 20, 40, 30, 50]. OK.")
        dList.SwapItems(five, three)
        require.Equal(t, []int{10, 20, 40, 50, 30}, ElementsLeftToRight(dList.(*dlist.DList)))
        require.Equal(t, []int{10, 20, 40, 50, 30}, ElementsRightToLeft(dList.(*dlist.DList)))
        fmt.Println("[10, 20, 40, 50, 30]. OK.")
        dList.SwapItems(one, two)
        require.Equal(t, []int{20, 10, 40, 50, 30}, ElementsLeftToRight(dList.(*dlist.DList)))
        require.Equal(t, []int{20, 10, 40, 50, 30}, ElementsRightToLeft(dList.(*dlist.DList)))
        fmt.Println("[20, 10, 40, 50, 30]. OK.")
        dList.SwapItems(one, two)
        require.Equal(t, []int{10, 20, 40, 50, 30}, ElementsLeftToRight(dList.(*dlist.DList)))
        require.Equal(t, []int{10, 20, 40, 50, 30}, ElementsRightToLeft(dList.(*dlist.DList)))
        fmt.Println("[10, 20, 40, 50, 30]. OK.")
        dList.SwapItems(three, five)
        require.Equal(t, []int{10, 20, 40, 30, 50}, ElementsLeftToRight(dList.(*dlist.DList)))
        require.Equal(t, []int{10, 20, 40, 30, 50}, ElementsRightToLeft(dList.(*dlist.DList)))
        fmt.Println("[10, 20, 40, 30, 50]. OK.")
        dList.SwapItems(three, four)
        require.Equal(t, []int{10, 20, 30, 40, 50}, ElementsLeftToRight(dList.(*dlist.DList)))
        require.Equal(t, []int{10, 20, 30, 40, 50}, ElementsRightToLeft(dList.(*dlist.DList)))
        fmt.Println("[10, 20, 30, 40, 50]. OK.")
    })
}


func TestDListSortInterface(t *testing.T) {

    t.Run("Let's sort Dlist.", func(t *testing.T) {
        dList := dlist.NewDList()
        dList.PushToRightEdge(10) // [10]
        dList.PushToRightEdge(30) // [10, 30]
        dList.PushToRightEdge(20) // [10, 30, 20]
        dList.PushToRightEdge(50) // [10, 30, 20, 50]
        dList.PushToRightEdge(40) // [10, 30, 20, 50, 40]
        sample := []int{10, 30, 20, 50, 40}
        require.Equal(t, sample, ElementsLeftToRight(dList.(*dlist.DList)))
        require.Equal(t, sample, ElementsRightToLeft(dList.(*dlist.DList)))
        fmt.Printf("\nTest dList before sort: %v. OK.\n", sample)
        fmt.Printf("\nList before sort with Stringer() formatting:\n%s\n", dList)
        sort.Sort(dList.(*dlist.DList))
        expected := []int{10, 20, 30, 40, 50}
        require.Equal(t, expected, ElementsLeftToRight(dList.(*dlist.DList)))
        require.Equal(t, expected, ElementsRightToLeft(dList.(*dlist.DList)))
        fmt.Printf("\nTest dList after sort: %v. OK.\n", expected)
        fmt.Printf("\nList after sort with Stringer() formatting:\n%s\n", dList)
    })
}

```

</details>

* TestDListSimple

```shell
go test -v -run TestDListSimple ./dlist.go ./dlist_stringer.go ./dlist_test.go  > dlist_test.go.simple.txt
```

<details>
<summary>см. лог "TestDListSimple":</summary>

```text
=== RUN   TestDListSimple
=== RUN   TestDListSimple/Zero-value_DListItem_test.

[zero-value] is: 
-------------------
 Item: 0xc00002e300
-------------------
value: <nil>
 left: 0x0
right: 0x0
-------------------

=== RUN   TestDListSimple/DListItem_direct_and_vice_versa_referencies_test.

[1] <--> [2] <--> [3]

[1] is: 
-------------------
 Item: 0xc00002e340
-------------------
value: 1
 left: 0x0
right: 0x0
-------------------

add [2]: 
-------------------
 Item: 0xc00002e360
-------------------
value: 2
 left: 0xc00002e340
right: 0x0
-------------------

[1] become: 
-------------------
 Item: 0xc00002e340
-------------------
value: 1
 left: 0x0
right: 0xc00002e360
-------------------

add [3]: 
-------------------
 Item: 0xc00002e380
-------------------
value: 3
 left: 0xc00002e360
right: 0x0
-------------------

[2] become: 
-------------------
 Item: 0xc00002e360
-------------------
value: 2
 left: 0xc00002e340
right: 0xc00002e380
-------------------

first.RightNeighbour().RightNeighbour().RightNeighbour() is nil. OK.
first.RightNeighbour().RightNeighbour() is third. OK.
third.LeftNeighbour().LeftNeighbour().LeftNeighbour() is nil. OK.
third.LeftNeighbour().LeftNeighbour() is first. OK.

=== RUN   TestDListSimple/Empty_DList_test.

List is:
 
 (nil:0x0)
    ^|
    L|R
     |v
 (nil:0x0)

=== RUN   TestDListSimple/DList_init_test.

List was:
 
 (nil:0x0)
    ^|
    L|R
     |v
 (nil:0x0)

Item was: 
-------------------
 Item: 0xc00002e3e0
-------------------
value: 1
 left: 0x0
right: 0x0
-------------------

dList.LeftEdge() and dList.RightEdge() is item. OK.

List become:
 
 (nil:0x0)
    ^|
    L|R
     |v
-------------------
 Item: 0xc00002e3e0
-------------------
value: 1
 left: 0x0
right: 0x0
-------------------
    ^|
    L|R
     |v
 (nil:0x0)

=== RUN   TestDListSimple/Little_DList_test.

List was:
 
 (nil:0x0)
    ^|
    L|R
     |v
 (nil:0x0)

Item [1] become:
 
-------------------
 Item: 0xc00002e420
-------------------
value: 1
 left: 0x0
right: 0x0
-------------------

list.LeftEdge() become:
 
-------------------
 Item: 0xc00002e420
-------------------
value: 1
 left: 0x0
right: 0x0
-------------------

list.RightEdge() become:
 
-------------------
 Item: 0xc00002e420
-------------------
value: 1
 left: 0x0
right: 0x0
-------------------

Item [2] become:
 
-------------------
 Item: 0xc00002e440
-------------------
value: 2
 left: 0x0
right: 0xc00002e420
-------------------

list.LeftEdge() become:
 
-------------------
 Item: 0xc00002e440
-------------------
value: 2
 left: 0x0
right: 0xc00002e420
-------------------

list.RightEdge() become:
 
-------------------
 Item: 0xc00002e420
-------------------
value: 1
 left: 0xc00002e440
right: 0x0
-------------------

List become:
 
 (nil:0x0)
    ^|
    L|R
     |v
-------------------
 Item: 0xc00002e440
-------------------
value: 2
 left: 0x0
right: 0xc00002e420
-------------------
    ^|
    L|R
     |v
-------------------
 Item: 0xc00002e420
-------------------
value: 1
 left: 0xc00002e440
right: 0x0
-------------------
    ^|
    L|R
     |v
 (nil:0x0)

Was removed:
 
-------------------
 Item: 0xc00002e440
-------------------
value: 2
 left: 0x0
right: 0xc00002e420
-------------------

List become:
 
 (nil:0x0)
    ^|
    L|R
     |v
-------------------
 Item: 0xc00002e420
-------------------
value: 1
 left: 0x0
right: 0x0
-------------------
    ^|
    L|R
     |v
 (nil:0x0)

Was removed:
 
-------------------
 Item: 0xc00002e420
-------------------
value: 1
 left: 0x0
right: 0x0
-------------------

List become:
 
 (nil:0x0)
    ^|
    L|R
     |v
 (nil:0x0)
--- PASS: TestDListSimple (0.00s)
    --- PASS: TestDListSimple/Zero-value_DListItem_test. (0.00s)
    --- PASS: TestDListSimple/DListItem_direct_and_vice_versa_referencies_test. (0.00s)
    --- PASS: TestDListSimple/Empty_DList_test. (0.00s)
    --- PASS: TestDListSimple/DList_init_test. (0.00s)
    --- PASS: TestDListSimple/Little_DList_test. (0.00s)
PASS
ok  	command-line-arguments	0.007s

```

</details>

* TestDListComplex

```shell
go test -v -run TestDListComplex ./dlist.go ./dlist_stringer.go ./dlist_test.go  > dlist_test.go.complex.txt
```

<details>
<summary>см. лог "TestDListComplex":</summary>

```text
=== RUN   TestDListComplex
=== RUN   TestDListComplex/Сomplex_DList_processing_test.
Forward stroke check for [10, 20, 30]. OK.
middle.Value() is 20. OK.
middle was removed. OK.
Forward stroke check for [10, 30]. OK.
DList [10, 30] mixing values {40, 50, 60, 70, 80} with mod(2, index).
dList.Len() is 7. OK.
dList.LeftEdge().Value() is 80. OK.
dList.RightEdge().Value() is 70. OK.
Forward stroke check for [80, 60, 40, 10, 30, 50, 70]. OK.
Remove right and push back to right - check for [80, 60, 40, 10, 30, 50, 70]. OK.
Remove left and push back to left - check for [80, 60, 40, 10, 30, 50, 70]. OK.
Check for dList.LeftEdge() is 80 and dList.RightEdge() is 70. OK.
Check for dList.LeftEdge().Left() and dList.RightEdge().Right() is nils. OK.
--- PASS: TestDListComplex (0.00s)
    --- PASS: TestDListComplex/Сomplex_DList_processing_test. (0.00s)
PASS
ok  	command-line-arguments	0.007s

```

</details>

* TestDListComplex

```shell
go test -v -run TestDListSwap ./dlist.go ./dlist_stringer.go ./dlist_test.go  > dlist_test.go.swap.txt
```

<details>
<summary>см. лог "TestDListSwap":</summary>

```text
=== RUN   TestDListSwap
=== RUN   TestDListSwap/DList_swap_test.
[10, 20, 30, 40, 50]. OK.
swap different element-pairs
[10, 40, 30, 20, 50]. OK.
[10, 20, 30, 40, 50]. OK.
[50, 20, 30, 40, 10]. OK.
[10, 20, 30, 40, 50]. OK.
[30, 20, 10, 40, 50]. OK.
[10, 20, 30, 40, 50]. OK.
[10, 50, 30, 40, 20]. OK.
[10, 20, 30, 40, 50]. OK.
[10, 30, 20, 40, 50]. OK.
[10, 20, 30, 40, 50]. OK.
[10, 20, 40, 30, 50]. OK.
[10, 20, 40, 50, 30]. OK.
[20, 10, 40, 50, 30]. OK.
[10, 20, 40, 50, 30]. OK.
[10, 20, 40, 30, 50]. OK.
[10, 20, 30, 40, 50]. OK.
--- PASS: TestDListSwap (0.00s)
    --- PASS: TestDListSwap/DList_swap_test. (0.00s)
PASS
ok  	command-line-arguments	0.007s

```

</details>

## Документация

```shell
go doc -all ./ > dlist.doc.txt
```

<details>
<summary>см. документацию:</summary>

```text

```

</details>

## Сортировка

Я, как мне кажется (😉), подобрал хороший пример для наглядной демонстрации интерфейса, требуемого sort.Sort (уже присутствуют в коде выше):

* `func (list *DList) Len()`
* `func (list *DList) Less(i, j int) bool`
* `func (list *DList) Swap(i, j int)`

В данном варианте не подойдет sort.Slice, так как перестановка элементов в двусвязном списке влечет перестановку указателей на соседей и с соседей на переставляемые элементы.

Особенность в том, что заранее зная, что будет реализован интерфейс sort.Sort, пришлось отказаться от именования Swap в самой структуре двусвязного списка, так как сигнатура должна быть `Swap (i, j int)`, а не как положено для двусвязного `Swap (i, j *DListItem)`. Это подводный камень для рефакторинга - стоит заранее избегать именований интерфейсных методов.

* TestListSortInterface

```shell
go test -v -run TestDListSortInterface ./dlist.go ./dlist_stringer.go ./dlist_test.go  > dlist_test.go.sort.txt
```

<details>
<summary>см. лог "TestDListSortInterface" (список упорядочился):</summary>

```text
=== RUN   TestDListSortInterface
=== RUN   TestDListSortInterface/Let's_sort_Dlist.

Test dList before sort: [10 30 20 50 40]. OK.

List before sort with Stringer() formatting:

 (nil:0x0)
    ^|
    L|R
     |v
-------------------
 Item: 0xc00002e300
-------------------
value: 10
 left: 0x0
right: 0xc00002e320
-------------------
    ^|
    L|R
     |v
-------------------
 Item: 0xc00002e320
-------------------
value: 30
 left: 0xc00002e300
right: 0xc00002e340
-------------------
    ^|
    L|R
     |v
-------------------
 Item: 0xc00002e340
-------------------
value: 20
 left: 0xc00002e320
right: 0xc00002e360
-------------------
    ^|
    L|R
     |v
-------------------
 Item: 0xc00002e360
-------------------
value: 50
 left: 0xc00002e340
right: 0xc00002e380
-------------------
    ^|
    L|R
     |v
-------------------
 Item: 0xc00002e380
-------------------
value: 40
 left: 0xc00002e360
right: 0x0
-------------------
    ^|
    L|R
     |v
 (nil:0x0)

Test dList after sort: [10 20 30 40 50]. OK.

List after sort with Stringer() formatting:

 (nil:0x0)
    ^|
    L|R
     |v
-------------------
 Item: 0xc00002e300
-------------------
value: 10
 left: 0x0
right: 0xc00002e340
-------------------
    ^|
    L|R
     |v
-------------------
 Item: 0xc00002e340
-------------------
value: 20
 left: 0xc00002e300
right: 0xc00002e320
-------------------
    ^|
    L|R
     |v
-------------------
 Item: 0xc00002e320
-------------------
value: 30
 left: 0xc00002e340
right: 0xc00002e380
-------------------
    ^|
    L|R
     |v
-------------------
 Item: 0xc00002e380
-------------------
value: 40
 left: 0xc00002e320
right: 0xc00002e360
-------------------
    ^|
    L|R
     |v
-------------------
 Item: 0xc00002e360
-------------------
value: 50
 left: 0xc00002e380
right: 0x0
-------------------
    ^|
    L|R
     |v
 (nil:0x0)
--- PASS: TestDListSortInterface (0.00s)
    --- PASS: TestDListSortInterface/Let's_sort_Dlist. (0.00s)
PASS
ok  	command-line-arguments	0.007s

```

</details>

## Послесловие

>
> ```text
> Данный документ составлен с использованием разработанного шаблонизатора. 
> При его использовании избегайте рекурсивной вложенности.
> ```

см. ["Шаблонизатор"](https://github.com/BorisPlus/golang_notes/tree/master/templator)
