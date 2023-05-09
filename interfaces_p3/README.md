# Интерфейсы. Двусвязный список

## Реализация

Интерфейс:

```go
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

```

Наглядность:

```go
package main

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
func (listItem *ListItem) String() string {
    template := `
-------------------
 Item: %p
-------------------
value: %v
 left: %p
right: %p
-------------------`
    return fmt.Sprintf(template, listItem, listItem.Value(), listItem.LeftNeighbour(), listItem.RightNeighbour())
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
    for i := list.LeftEdge(); i != nil; i = i.RightNeighbour() {
        result += i.String()
        result += delimiter
    }
    result += Nill
    return result
}

```

Тестирование:

```go
package main

import (
    "fmt"
    "testing"

    "github.com/stretchr/testify/require"
)

// go test -v list.go list_string.go list_test.go

func ElementsLeftToRight(list *List) []int {
    elements := make([]int, 0, list.Length())
    for i := list.LeftEdge(); i != nil; i = i.RightNeighbour() {
        elements = append(elements, i.Value().(int))
    }
    return elements
}

func ElementsRightToLeft(list *List) []int {
    elements := make([]int, 0, list.Length())
    for i := list.RightEdge(); i != nil; i = i.LeftNeighbour() {
        elements = append([]int{i.Value().(int)}, elements...)
    }
    return elements
}

func TestList(t *testing.T) {
    t.Run("Zero-value ListItem test.", func(t *testing.T) {
        zeroValueItem := ListItem{}
        require.Nil(t, zeroValueItem.Value())
        require.Nil(t, zeroValueItem.LeftNeighbour())
        require.Nil(t, zeroValueItem.RightNeighbour())
        fmt.Println("\n[zero-value] is:", &zeroValueItem)
        fmt.Println()
    })

    t.Run("ListItem direct and vice versa referencies test.", func(t *testing.T) {
        fmt.Println("\n[1] <--> [2] <--> [3]")
        first := &ListItem{
            value:          1,
            leftNeighbour:  nil,
            rightNeighbour: nil,
        }
        fmt.Println("\n[1] is:", first)

        second := &ListItem{
            value:          2,
            leftNeighbour:  first,
            rightNeighbour: nil,
        }

        first.SetRightNeighbour(second)
        fmt.Println("\nadd [2]:", second)
        fmt.Println("\n[1] become:", first)

        third := ListItem{
            value:          3,
            leftNeighbour:  second,
            rightNeighbour: nil,
        }

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

    t.Run("Empty List test.", func(t *testing.T) {
        list := NewList()
        require.Equal(t, 0, list.Length())
        require.Nil(t, list.LeftEdge())
        require.Nil(t, list.RightEdge())
        fmt.Println("\nList is:\n", list)
        fmt.Println()
    })

    t.Run("List init test.", func(t *testing.T) {
        list := NewList()

        fmt.Println("\nList was:\n", list)

        item := list.PushToLeftEdge(1)
        fmt.Println("\nItem was:", item)
        require.Equal(t, item, list.LeftEdge())
        require.Equal(t, item, list.RightEdge())
        fmt.Println()
        fmt.Println("list.LeftEdge() and list.RightEdge() is item. OK.")
        fmt.Println("\nList become:\n", list)
        fmt.Println()
    })

    t.Run("Little List test.", func(t *testing.T) {
        list := NewList()
        fmt.Println("\nList was:\n", list)

        itemFirst := list.PushToLeftEdge(1) // [1]
        fmt.Println("\nItem [1] become:\n", itemFirst)

        leftEnd := list.LeftEdge()
        fmt.Println("\nlist.LeftEdge() become:\n", leftEnd)
        rightEnd := list.RightEdge()
        fmt.Println("\nlist.RightEdge() become:\n", rightEnd)

        itemSecond := list.PushToLeftEdge(2) // [2]
        fmt.Println("\nItem [2] become:\n", itemSecond)

        leftEnd = list.LeftEdge()
        fmt.Println("\nlist.LeftEdge() become:\n", leftEnd)
        rightEnd = list.RightEdge()
        fmt.Println("\nlist.RightEdge() become:\n", rightEnd)
        fmt.Println("\nList become:\n", list)
        require.Equal(t, itemSecond, list.LeftEdge())
        require.Equal(t, itemFirst, list.RightEdge())

        someItem, _ := list.Remove(list.LeftEdge())
        fmt.Println("\nWas removed:\n", someItem)
        require.Equal(t, itemSecond, someItem)

        fmt.Println("\nList become:\n", list)

        require.Equal(t, list.LeftEdge(), list.RightEdge())
        require.Nil(t, list.LeftEdge().LeftNeighbour())
        require.Nil(t, list.LeftEdge().RightNeighbour())

        someItem, _ = list.Remove(list.LeftEdge())
        fmt.Println("\nWas removed:\n", someItem)
        require.Equal(t, itemFirst, someItem)

        fmt.Println("\nList become:\n", list)

        require.Equal(t, list.LeftEdge(), list.RightEdge())
        require.Nil(t, list.LeftEdge())
        require.Nil(t, list.RightEdge())
    })
}

func TestListComplex(t *testing.T) {
    t.Run("Сomplex List processing test.", func(t *testing.T) {
        list := NewList()
        list.PushToRightEdge(10) // [10]
        list.PushToRightEdge(20) // [10, 20]
        list.PushToRightEdge(30) // [10, 20, 30]
        require.Equal(t, []int{10, 20, 30}, ElementsLeftToRight(list.(*List)))
        require.Equal(t, []int{10, 20, 30}, ElementsRightToLeft(list.(*List)))
        fmt.Println("Forward stroke check for [10, 20, 30]. OK.")
        require.Equal(t, 3, list.Length())
        middle := list.LeftEdge().RightNeighbour()
        require.Equal(t, middle.Value(), 20)
        fmt.Printf("middle.Value() is %v. OK.\n", middle.Value())
        list.Remove(middle)
        fmt.Printf("middle was removed. OK.\n")
        require.Equal(t, 2, list.Length())
        require.Equal(t, []int{10, 30}, ElementsLeftToRight(list.(*List)))
        require.Equal(t, []int{10, 30}, ElementsRightToLeft(list.(*List)))
        fmt.Println("Forward stroke check for [10, 30]. OK.")
        for i, v := range [...]int{40, 50, 60, 70, 80} {
            if i%2 == 0 {
                list.PushToLeftEdge(v)
            } else {
                list.PushToRightEdge(v)
            }
        } // [80, 60, 40, 10, 30, 50, 70]
        fmt.Println("List [10, 30] mixing values {40, 50, 60, 70, 80} with mod(2, index).")

        require.Equal(t, 7, list.Length())
        fmt.Printf("list.Length() is %v. OK.\n", list.Length())
        require.Equal(t, 80, list.LeftEdge().Value())
        fmt.Printf("list.LeftEdge().Value() is %v. OK.\n", list.LeftEdge().Value())
        require.Equal(t, 70, list.RightEdge().Value())
        fmt.Printf("list.RightEdge().Value() is %v. OK.\n", list.RightEdge().Value())

        require.Equal(t, []int{80, 60, 40, 10, 30, 50, 70}, ElementsLeftToRight(list.(*List)))
        require.Equal(t, []int{80, 60, 40, 10, 30, 50, 70}, ElementsRightToLeft(list.(*List)))
        fmt.Println("Forward stroke check for [80, 60, 40, 10, 30, 50, 70]. OK.")

        rightEnd := list.RightEdge()
        list.Remove(rightEnd)
        list.PushToRightEdge(rightEnd.Value())
        fmt.Println("Remove right and push back to right - check for [80, 60, 40, 10, 30, 50, 70]. OK.")
        leftEnd, _ := list.Remove(list.LeftEdge())
        list.PushToLeftEdge(leftEnd.Value())
        fmt.Println("Remove left and push back to left - check for [80, 60, 40, 10, 30, 50, 70]. OK.")
        // - check for nil-refs of first and last
        require.Equal(t, list.LeftEdge().Value(), 80)
        require.Equal(t, list.RightEdge().Value(), 70)
        fmt.Println("Check for list.LeftEdge() is 80 and list.RightEdge() is 70. OK.")
        require.Nil(t, list.LeftEdge().LeftNeighbour())
        require.Nil(t, list.RightEdge().RightNeighbour())
        fmt.Println("Check for list.LeftEdge().Left() and list.RightEdge().Right() is nils. OK.")

    })
}

func TestListSwap(t *testing.T) {
    t.Run("List swap test.", func(t *testing.T) {
        list := NewList()
        one := list.PushToRightEdge(10)   // [10]
        two := list.PushToRightEdge(20)   // [10, 20]
        three := list.PushToRightEdge(30) // [10, 20, 30]
        four := list.PushToRightEdge(40)  // [10, 20, 30, 40]
        five := list.PushToRightEdge(50)  // [10, 20, 30, 40, 50]
        require.Equal(t, []int{10, 20, 30, 40, 50}, ElementsLeftToRight(list.(*List)))
        require.Equal(t, []int{10, 20, 30, 40, 50}, ElementsRightToLeft(list.(*List)))
        fmt.Println("[10, 20, 30, 40, 50]. OK.")
        list.SwapItems(two, four)
        fmt.Println("swap different element-pairs")
        require.Equal(t, []int{10, 40, 30, 20, 50}, ElementsLeftToRight(list.(*List)))
        require.Equal(t, []int{10, 40, 30, 20, 50}, ElementsRightToLeft(list.(*List)))
        fmt.Println("[10, 40, 30, 20, 50]. OK.")
        list.SwapItems(two, four)
        require.Equal(t, []int{10, 20, 30, 40, 50}, ElementsLeftToRight(list.(*List)))
        require.Equal(t, []int{10, 20, 30, 40, 50}, ElementsRightToLeft(list.(*List)))
        fmt.Println("[10, 20, 30, 40, 50]. OK.")
        list.SwapItems(one, five)
        require.Equal(t, []int{50, 20, 30, 40, 10}, ElementsLeftToRight(list.(*List)))
        require.Equal(t, []int{50, 20, 30, 40, 10}, ElementsRightToLeft(list.(*List)))
        fmt.Println("[50, 20, 30, 40, 10]. OK.")
        list.SwapItems(one, five)
        require.Equal(t, []int{10, 20, 30, 40, 50}, ElementsLeftToRight(list.(*List)))
        require.Equal(t, []int{10, 20, 30, 40, 50}, ElementsRightToLeft(list.(*List)))
        fmt.Println("[10, 20, 30, 40, 50]. OK.")
        list.SwapItems(one, three)
        require.Equal(t, []int{30, 20, 10, 40, 50}, ElementsLeftToRight(list.(*List)))
        require.Equal(t, []int{30, 20, 10, 40, 50}, ElementsRightToLeft(list.(*List)))
        fmt.Println("[30, 20, 10, 40, 50]. OK.")
        list.SwapItems(one, three)
        require.Equal(t, []int{10, 20, 30, 40, 50}, ElementsLeftToRight(list.(*List)))
        require.Equal(t, []int{10, 20, 30, 40, 50}, ElementsRightToLeft(list.(*List)))
        fmt.Println("[10, 20, 30, 40, 50]. OK.")
        list.SwapItems(five, two)
        require.Equal(t, []int{10, 50, 30, 40, 20}, ElementsLeftToRight(list.(*List)))
        require.Equal(t, []int{10, 50, 30, 40, 20}, ElementsRightToLeft(list.(*List)))
        fmt.Println("[10, 50, 30, 40, 20]. OK.")
        list.SwapItems(two, five)
        require.Equal(t, []int{10, 20, 30, 40, 50}, ElementsLeftToRight(list.(*List)))
        require.Equal(t, []int{10, 20, 30, 40, 50}, ElementsRightToLeft(list.(*List)))
        fmt.Println("[10, 20, 30, 40, 50]. OK.")
        list.SwapItems(two, three)
        require.Equal(t, []int{10, 30, 20, 40, 50}, ElementsLeftToRight(list.(*List)))
        require.Equal(t, []int{10, 30, 20, 40, 50}, ElementsRightToLeft(list.(*List)))
        fmt.Println("[10, 30, 20, 40, 50]. OK.")
        list.SwapItems(two, three)
        require.Equal(t, []int{10, 20, 30, 40, 50}, ElementsLeftToRight(list.(*List)))
        require.Equal(t, []int{10, 20, 30, 40, 50}, ElementsRightToLeft(list.(*List)))
        fmt.Println("[10, 20, 30, 40, 50]. OK.")
        list.SwapItems(four, three)
        require.Equal(t, []int{10, 20, 40, 30, 50}, ElementsLeftToRight(list.(*List)))
        require.Equal(t, []int{10, 20, 40, 30, 50}, ElementsRightToLeft(list.(*List)))
        fmt.Println("[10, 20, 40, 30, 50]. OK.")
        list.SwapItems(five, three)
        require.Equal(t, []int{10, 20, 40, 50, 30}, ElementsLeftToRight(list.(*List)))
        require.Equal(t, []int{10, 20, 40, 50, 30}, ElementsRightToLeft(list.(*List)))
        fmt.Println("[10, 20, 40, 50, 30]. OK.")
        list.SwapItems(one, two)
        require.Equal(t, []int{20, 10, 40, 50, 30}, ElementsLeftToRight(list.(*List)))
        require.Equal(t, []int{20, 10, 40, 50, 30}, ElementsRightToLeft(list.(*List)))
        fmt.Println("[20, 10, 40, 50, 30]. OK.")
        list.SwapItems(one, two)
        require.Equal(t, []int{10, 20, 40, 50, 30}, ElementsLeftToRight(list.(*List)))
        require.Equal(t, []int{10, 20, 40, 50, 30}, ElementsRightToLeft(list.(*List)))
        fmt.Println("[10, 20, 40, 50, 30]. OK.")
        list.SwapItems(three, five)
        require.Equal(t, []int{10, 20, 40, 30, 50}, ElementsLeftToRight(list.(*List)))
        require.Equal(t, []int{10, 20, 40, 30, 50}, ElementsRightToLeft(list.(*List)))
        fmt.Println("[10, 20, 40, 30, 50]. OK.")
        list.SwapItems(three, four)
        require.Equal(t, []int{10, 20, 30, 40, 50}, ElementsLeftToRight(list.(*List)))
        require.Equal(t, []int{10, 20, 30, 40, 50}, ElementsRightToLeft(list.(*List)))
        fmt.Println("[10, 20, 30, 40, 50]. OK.")
    })
}

```

```shell
go test -v ./list.go ./list_stringer.go ./list_test.go  > list_test.go.txt
```

Лог:

```text
=== RUN   TestList
=== RUN   TestList/Zero-value_ListItem_test.

[zero-value] is: 
-------------------
 Item: 0xc00002e2e0
-------------------
value: <nil>
 left: 0x0
right: 0x0
-------------------

=== RUN   TestList/ListItem_direct_and_vice_versa_referencies_test.

[1] <--> [2] <--> [3]

[1] is: 
-------------------
 Item: 0xc00002e300
-------------------
value: 1
 left: 0x0
right: 0x0
-------------------

add [2]: 
-------------------
 Item: 0xc00002e320
-------------------
value: 2
 left: 0xc00002e300
right: 0x0
-------------------

[1] become: 
-------------------
 Item: 0xc00002e300
-------------------
value: 1
 left: 0x0
right: 0xc00002e320
-------------------

add [3]: 
-------------------
 Item: 0xc00002e340
-------------------
value: 3
 left: 0xc00002e320
right: 0x0
-------------------

[2] become: 
-------------------
 Item: 0xc00002e320
-------------------
value: 2
 left: 0xc00002e300
right: 0xc00002e340
-------------------

first.RightNeighbour().RightNeighbour().RightNeighbour() is nil. OK.
first.RightNeighbour().RightNeighbour() is third. OK.
third.LeftNeighbour().LeftNeighbour().LeftNeighbour() is nil. OK.
third.LeftNeighbour().LeftNeighbour() is first. OK.

=== RUN   TestList/Empty_List_test.

List is:
 
 (nil:0x0)
    ^|
    L|R
     |v
 (nil:0x0)

=== RUN   TestList/List_init_test.

List was:
 
 (nil:0x0)
    ^|
    L|R
     |v
 (nil:0x0)

Item was: 
-------------------
 Item: 0xc00002e360
-------------------
value: 1
 left: 0x0
right: 0x0
-------------------

list.LeftEdge() and list.RightEdge() is item. OK.

List become:
 
 (nil:0x0)
    ^|
    L|R
     |v
-------------------
 Item: 0xc00002e360
-------------------
value: 1
 left: 0x0
right: 0x0
-------------------
    ^|
    L|R
     |v
 (nil:0x0)

=== RUN   TestList/Little_List_test.

List was:
 
 (nil:0x0)
    ^|
    L|R
     |v
 (nil:0x0)

Item [1] become:
 
-------------------
 Item: 0xc00002e380
-------------------
value: 1
 left: 0x0
right: 0x0
-------------------

list.LeftEdge() become:
 
-------------------
 Item: 0xc00002e380
-------------------
value: 1
 left: 0x0
right: 0x0
-------------------

list.RightEdge() become:
 
-------------------
 Item: 0xc00002e380
-------------------
value: 1
 left: 0x0
right: 0x0
-------------------

Item [2] become:
 
-------------------
 Item: 0xc00002e3a0
-------------------
value: 2
 left: 0x0
right: 0xc00002e380
-------------------

list.LeftEdge() become:
 
-------------------
 Item: 0xc00002e3a0
-------------------
value: 2
 left: 0x0
right: 0xc00002e380
-------------------

list.RightEdge() become:
 
-------------------
 Item: 0xc00002e380
-------------------
value: 1
 left: 0xc00002e3a0
right: 0x0
-------------------

List become:
 
 (nil:0x0)
    ^|
    L|R
     |v
-------------------
 Item: 0xc00002e3a0
-------------------
value: 2
 left: 0x0
right: 0xc00002e380
-------------------
    ^|
    L|R
     |v
-------------------
 Item: 0xc00002e380
-------------------
value: 1
 left: 0xc00002e3a0
right: 0x0
-------------------
    ^|
    L|R
     |v
 (nil:0x0)

Was removed:
 
-------------------
 Item: 0xc00002e3a0
-------------------
value: 2
 left: 0x0
right: 0xc00002e380
-------------------

List become:
 
 (nil:0x0)
    ^|
    L|R
     |v
-------------------
 Item: 0xc00002e380
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
 Item: 0xc00002e380
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
--- PASS: TestList (0.00s)
    --- PASS: TestList/Zero-value_ListItem_test. (0.00s)
    --- PASS: TestList/ListItem_direct_and_vice_versa_referencies_test. (0.00s)
    --- PASS: TestList/Empty_List_test. (0.00s)
    --- PASS: TestList/List_init_test. (0.00s)
    --- PASS: TestList/Little_List_test. (0.00s)
=== RUN   TestListComplex
=== RUN   TestListComplex/Сomplex_List_processing_test.
Forward stroke check for [10, 20, 30]. OK.
middle.Value() is 20. OK.
middle was removed. OK.
Forward stroke check for [10, 30]. OK.
List [10, 30] mixing values {40, 50, 60, 70, 80} with mod(2, index).
list.Length() is 7. OK.
list.LeftEdge().Value() is 80. OK.
list.RightEdge().Value() is 70. OK.
Forward stroke check for [80, 60, 40, 10, 30, 50, 70]. OK.
Remove right and push back to right - check for [80, 60, 40, 10, 30, 50, 70]. OK.
Remove left and push back to left - check for [80, 60, 40, 10, 30, 50, 70]. OK.
Check for list.LeftEdge() is 80 and list.RightEdge() is 70. OK.
Check for list.LeftEdge().Left() and list.RightEdge().Right() is nils. OK.
--- PASS: TestListComplex (0.00s)
    --- PASS: TestListComplex/Сomplex_List_processing_test. (0.00s)
=== RUN   TestListSwap
=== RUN   TestListSwap/List_swap_test.
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
--- PASS: TestListSwap (0.00s)
    --- PASS: TestListSwap/List_swap_test. (0.00s)
PASS
ok  	command-line-arguments	0.006s

```

Документация:

```shell
go doc -all ./ > list.doc.txt
```

```text


FUNCTIONS

func Eq(x, y ListItem) bool
    Eq(x, y ListItem) - элементы равны тогда и только тогда, когда это один и
    тот же элемент по памяти.


TYPES

type List struct {
	// Has unexported fields.
}
    List - структура двусвязного списка.

func (list *List) Contains(item *ListItem) bool
    Contains(item *ListItem) - проверить, есть ли элемент в списке.

func (list *List) LeftEdge() *ListItem
    LeftEdge() - получить элемент из левого края двусвязного списка.

func (list *List) Length() int
    Length() - получить длину двусвязного списка.

func (list *List) MoveToLeftEdge(item *ListItem) error
    MoveToFront(item *ListItem) - переместить элемент в начало двусвязного
    списка.

func (list *List) PushToLeftEdge(value interface{}) *ListItem
    PushToLeftEdge(value interface{}) - добавить значение в левый край
    двусвязного списка.

func (list *List) PushToRightEdge(value interface{}) *ListItem
    PushToRightEdge(value interface{}) - добавить значение в правый край
    двусвязного списка.

func (list *List) Remove(item *ListItem) (*ListItem, error)
    Remove(item *ListItem) - удалить элемент из двусвязного списка.

func (list *List) RightEdge() *ListItem
    RightEdge() - получить элемент из правого края двусвязного списка.

func (list *List) String() string
    String() - представление двусвязного списка.

    Например,

    - пустой список:

        (nil:0x0)
           ^|
           L|R
            |v
        (nil:0x0)

    - список из двух элементов:

         (nil:0x0)
            ^|
            L|R
             |v
        -------------------
         Item: 0xc00002e3a0 <--------┐
        -------------------          |
        Value: 2                     |
         Left: 0x0                   |
        Right: 0xc00002e380  >>>-----|---┐ Right 0xc00002e380
        -------------------          |   | ссылается на
            ^|                       |   | блок 0xc00002e380
            L|R                      |   |
             |v                      |   |
        -------------------          |   |
         Item: 0xc00002e380  <-----------┘
        -------------------          | Left 0xc00002e3a0
        Value: 1                     | ссылается на
         Left: 0xc00002e3a0  >>>-----┘ блок 0xc00002e3a0
        Right: 0x0
        -------------------
            ^|
            L|R
             |v
         (nil:0x0)

func (list *List) SwapItems(x, y *ListItem) error
    SwapItems(x, y *ListItem) - поменять элементы двусвязного списка местами.

type ListItem struct {
	// Has unexported fields.
}
    ListItem - структура элемента двусвязного списка.

func (listItem *ListItem) LeftNeighbour() *ListItem
    LeftNeighbour() - получить стоящий слева элемент.

func (listItem *ListItem) RightNeighbour() *ListItem
    RightNeighbour() - получить стоящий справа элемент.

func (listItem *ListItem) SetLeftNeighbour(item *ListItem)
    SetLeftNeighbour(item *ListItem) - установить стоящий слева элемент.

func (listItem *ListItem) SetRightNeighbour(item *ListItem)
    SetRightNeighbour(item *ListItem) - установить стоящий слева элемент.

func (listItem *ListItem) String() string
    String() - представление элемента двусвязного списка.

    Например,

        -------------------             -------------------
         Item: 0xc00002e400              Item: 0xc00002e400
        -------------------             -------------------
        Value: 30                или    Value: 30
         Left: 0xc00002e3c0              Left: 0x0
        Right: 0xc00002e440             Right: 0x0
        -------------------             -------------------

func (listItem *ListItem) Value() interface{}
    Value() - получить значение из элемента.

type Lister interface {
	Length() int
	LeftEdge() *ListItem
	RightEdge() *ListItem
	PushToLeftEdge(value interface{}) *ListItem
	PushToRightEdge(value interface{}) *ListItem
	Remove(item *ListItem) (*ListItem, error)
	SwapItems(x, y *ListItem) error
}
    Lister - интерфейс двусвязного списка.

func NewList() Lister


```

## Сортировка

Пример сортировки в соответствии с интерфейсом sort.Sort() - в случае, когда сортируем не слайсы и sort.Slice() не подходит.

Тестирование:

```go
package main

import (
    "fmt"
    "sort"
    "testing"

    "github.com/stretchr/testify/require"
)

// go test -v list.go list_stringer.go list_sort_test.go > list_sort_test.go.txt

func Elements(list *List) []int {
    elements := make([]int, 0, list.Length())
    for i := list.LeftEdge(); i != nil; i = i.RightNeighbour() {
        elements = append(elements, i.Value().(int))
    }
    return elements
}

func (list *List) Len() int {
    return list.Length()
}

func (list *List) GetByIndex(i int) (*ListItem, error) {
    if i >= list.Length() {
        return nil, fmt.Errorf("index is out of range")
    }
    item := list.LeftEdge()
    for j := 0; j < i; j++ {
        item = item.RightNeighbour()
    }
    return item, nil
}

func (list *List) Swap(i, j int) {
    iItem, _ := list.GetByIndex(i)
    jItem, _ := list.GetByIndex(j)
    list.SwapItems(iItem, jItem)
}

func (list *List) Less(i, j int) bool {
    iItem, _ := list.GetByIndex(i)
    jItem, _ := list.GetByIndex(j)
    return iItem.Value().(int) < jItem.Value().(int)
}

func TestSort(t *testing.T) {

    t.Run("Let's sort double-linked list.", func(t *testing.T) {
        list := NewList()
        list.PushToRightEdge(10) // [10]
        list.PushToRightEdge(30) // [10, 30]
        list.PushToRightEdge(20) // [10, 30, 20]
        list.PushToRightEdge(50) // [10, 30, 20, 50]
        list.PushToRightEdge(40) // [10, 30, 20, 50, 40]
        sample := []int{10, 30, 20, 50, 40}
        require.Equal(t, sample, Elements(list.(*List)))
        fmt.Printf("\nTest list before sort: %v. OK.\n", sample)
        fmt.Printf("\nList before sort with Stringer() formatting:\n%s\n", list)
        sort.Sort(list.(*List))
        expected := []int{10, 20, 30, 40, 50}
        require.Equal(t, expected, Elements(list.(*List)))
        fmt.Printf("\nTest list after sort: %v. OK.\n", expected)
        fmt.Printf("\nList after sort with Stringer() formatting:\n%s\n", list)
    })
}

```

```shell
go test -v ./list.go ./list_stringer.go ./list_sort_test.go  > list_sort_test.go.txt
```

Лог (список упорядочился):

```text
=== RUN   TestSort
=== RUN   TestSort/Let's_sort_double-linked_list.

Test list before sort: [10 30 20 50 40]. OK.

List before sort with Stringer() formatting:

 (nil:0x0)
    ^|
    L|R
     |v
-------------------
 Item: 0xc00015c2a0
-------------------
value: 10
 left: 0x0
right: 0xc00015c2c0
-------------------
    ^|
    L|R
     |v
-------------------
 Item: 0xc00015c2c0
-------------------
value: 30
 left: 0xc00015c2a0
right: 0xc00015c2e0
-------------------
    ^|
    L|R
     |v
-------------------
 Item: 0xc00015c2e0
-------------------
value: 20
 left: 0xc00015c2c0
right: 0xc00015c300
-------------------
    ^|
    L|R
     |v
-------------------
 Item: 0xc00015c300
-------------------
value: 50
 left: 0xc00015c2e0
right: 0xc00015c320
-------------------
    ^|
    L|R
     |v
-------------------
 Item: 0xc00015c320
-------------------
value: 40
 left: 0xc00015c300
right: 0x0
-------------------
    ^|
    L|R
     |v
 (nil:0x0)

Test list after sort: [10 20 30 40 50]. OK.

List after sort with Stringer() formatting:

 (nil:0x0)
    ^|
    L|R
     |v
-------------------
 Item: 0xc00015c2a0
-------------------
value: 10
 left: 0x0
right: 0xc00015c2e0
-------------------
    ^|
    L|R
     |v
-------------------
 Item: 0xc00015c2e0
-------------------
value: 20
 left: 0xc00015c2a0
right: 0xc00015c2c0
-------------------
    ^|
    L|R
     |v
-------------------
 Item: 0xc00015c2c0
-------------------
value: 30
 left: 0xc00015c2e0
right: 0xc00015c320
-------------------
    ^|
    L|R
     |v
-------------------
 Item: 0xc00015c320
-------------------
value: 40
 left: 0xc00015c2c0
right: 0xc00015c300
-------------------
    ^|
    L|R
     |v
-------------------
 Item: 0xc00015c300
-------------------
value: 50
 left: 0xc00015c320
right: 0x0
-------------------
    ^|
    L|R
     |v
 (nil:0x0)
--- PASS: TestSort (0.00s)
    --- PASS: TestSort/Let's_sort_double-linked_list. (0.00s)
PASS
ok  	command-line-arguments	0.007s

```

## Послесловие

>
> ```text
> Данный документ составлен с использованием разработанного шаблонизатора. 
> При его использовании избегайте рекурсивной вложенности.
> ```

см. ["Шаблонизатор"](https://github.com/BorisPlus/golang_notes/tree/master/templator)
