# Интерфейсы. Двухсвязный список

## Реализация

Интерфейс:

```go
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
    //     y.right = xRight.right.left
    //     y.left = xLeft
    // } else {
    //     y.right = xRight
    //     y.left = xLeft.left.right
    // }

    // if yLeft.left != nil {
    //     x.left = yLeft.left.right
    //     x.right = yRight
    // } else {
    //     x.left = yLeft
    //     x.right = yRight.right.left
    // }

    // if yLeft.left != nil{
    //     x.left = yLeft.left.right
    // } else if yLeft.right != nil {
    //     x.left = yLeft.right.left
    // }

    // xLeft.right = y
    // yRight.left = x
    // if x.left != nil {
    //     x.left.right = y
    // }
    // x.right = y.right

    // if y.right != nil {
    //     y.right.left = x
    // }
    // yLeft := y.left
    // y.left = x.left

    // if xRight != nil {
    //     xRight.left = y
    // }

    // if xLeft != nil {
    //     xLeft.right = y
    // }

    // if yLeft != nil {
    //     yLeft.right = x
    // }
    // y.right = xRight

    // if x.left != nil {
    //     x.left.right = y
    // }
    // x.right = y.right
    // xLeft, xRight, yLeft, yRight := , x.right, y.left, y.right
    // x.left, x.right, y.left, y.right = yLeft, yRight, xLeft, xRight
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

// String - наглядное представление значения элемента двусвязного списка.
//
// Например,
//
//    -------------------             -------------------
//    Item:  0xc00002e400             Item: 0xc00002e400
//    -------------------             -------------------
//    Value: 30                или    Value: 30
//    Left:  0xc00002e3c0             Left:  0x0
//    Right: 0xc00002e440             Right: 0x0
//    -------------------             -------------------
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
//    Item:  0xc00002e3a0 <--------┐
//    -------------------          |
//    Value: 2                     |
//    Left:  0x0                   |
//    Right: 0xc00002e380  >>>-----|---┐ Right 0xc00002e380
//    -------------------          |   | ссылается на
//        ^|                       |   | блок 0xc00002e380
//        L|R                      |   |
//         |v                      |   |
//    -------------------          |   |
//    Item:  0xc00002e380  <-----------┘
//    -------------------          | Left 0xc00002e3a0
//    Value: 1                     | ссылается на
//    Left:  0xc00002e3a0  >>>-----┘ блок 0xc00002e3a0
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
    for i := list.LeftEnd(); i != nil; i = i.Right() {
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

func Elements(list *List) []int {
    elements := make([]int, 0, list.Len())
    for i := list.LeftEnd(); i != nil; i = i.Right() {
        elements = append(elements, i.Value().(int))
    }
    return elements
}

func TestList(t *testing.T) {
    t.Run("Zero-value ListItem test.", func(t *testing.T) {
        zeroValueItem := ListItem{}
        require.Nil(t, zeroValueItem.value)
        require.Nil(t, zeroValueItem.left)
        require.Nil(t, zeroValueItem.right)
        fmt.Println("\n[zero-value] is:", &zeroValueItem)
        fmt.Println()
    })

    t.Run("ListItem direct and vice versa referencies test.", func(t *testing.T) {
        fmt.Println("\n[1] <--> [2] <--> [3]")
        first := &ListItem{
            value: 1,
            left:  nil,
            right: nil,
        }
        fmt.Println("\n[1] is:", first)

        second := &ListItem{
            value: 2,
            left:  first,
            right: nil,
        }
        first.right = second
        fmt.Println("\nadd [2]:", second)
        fmt.Println("\n[1] become:", first)

        third := ListItem{
            value: 3,
            left:  second,
            right: nil,
        }

        second.right = &third
        fmt.Println("\nadd [3]:", &third)
        fmt.Println("\n[2] become:", second)
        fmt.Println()

        fmt.Println("first.Right().Right().Right() is nil. OK.")
        require.Nil(t, first.Right().Right().Right())
        fmt.Println("first.Right().Right() is third. OK.")
        require.Equal(t, &third, first.Right().Right())
        fmt.Println("third.Left().Left().Left() is nil. OK.")
        require.Nil(t, third.Left().Left().Left())
        fmt.Println("third.Left().Left() is first. OK.")
        require.Equal(t, first, third.Left().Left())
        fmt.Println()
    })

    t.Run("Empty List test.", func(t *testing.T) {
        list := NewList()
        require.Equal(t, 0, list.Len())
        require.Nil(t, list.LeftEnd())
        require.Nil(t, list.RightEnd())
        fmt.Println("\nList is:\n", list)
        fmt.Println()
    })

    t.Run("List init test.", func(t *testing.T) {
        list := NewList()

        fmt.Println("\nList was:\n", list)

        item := list.PushToLeftEnd(1)
        fmt.Println("\nItem was:", item)
        require.Equal(t, item, list.LeftEnd())
        require.Equal(t, item, list.RightEnd())
        fmt.Println()
        fmt.Println("list.LeftEnd() and list.RightEnd() is item. OK.")
        fmt.Println("\nList become:\n", list)
        fmt.Println()
    })

    t.Run("Little List test.", func(t *testing.T) {
        list := NewList()
        fmt.Println("\nList was:\n", list)

        itemFirst := list.PushToLeftEnd(1) // [1]
        fmt.Println("\nItem [1] become:\n", itemFirst)

        leftEnd := list.LeftEnd()
        fmt.Println("\nlist.LeftEnd() become:\n", leftEnd)
        rightEnd := list.RightEnd()
        fmt.Println("\nlist.RightEnd() become:\n", rightEnd)

        itemSecond := list.PushToLeftEnd(2) // [2]
        fmt.Println("\nItem [2] become:\n", itemSecond)

        leftEnd = list.LeftEnd()
        fmt.Println("\nlist.LeftEnd() become:\n", leftEnd)
        rightEnd = list.RightEnd()
        fmt.Println("\nlist.RightEnd() become:\n", rightEnd)
        fmt.Println("\nList become:\n", list)
        require.Equal(t, itemSecond, list.LeftEnd())
        require.Equal(t, itemFirst, list.RightEnd())

        someItem, _ := list.Remove(list.LeftEnd())
        fmt.Println("\nWas removed:\n", someItem)
        require.Equal(t, itemSecond, someItem)

        fmt.Println("\nList become:\n", list)

        require.Equal(t, list.LeftEnd(), list.RightEnd())
        require.Nil(t, list.LeftEnd().Left())
        require.Nil(t, list.LeftEnd().Right())

        someItem, _ = list.Remove(list.LeftEnd())
        fmt.Println("\nWas removed:\n", someItem)
        require.Equal(t, itemFirst, someItem)

        fmt.Println("\nList become:\n", list)

        require.Equal(t, list.LeftEnd(), list.RightEnd())
        require.Nil(t, list.LeftEnd())
        require.Nil(t, list.RightEnd())
    })
}

func TestListComplex(t *testing.T) {
    t.Run("Сomplex List processing test.", func(t *testing.T) {
        list := NewList()
        list.PushToRightEnd(10) // [10]
        list.PushToRightEnd(20) // [10, 20]
        list.PushToRightEnd(30) // [10, 20, 30]
        require.Equal(t, []int{10, 20, 30}, Elements(list.(*List)))
        fmt.Println("Forward stroke check for [10, 20, 30]. OK.")
        require.Equal(t, 3, list.Len())
        middle := list.LeftEnd().Right()
        require.Equal(t, middle.Value(), 20)
        fmt.Printf("middle.Value() is %v. OK.\n", middle.Value())
        list.Remove(middle)
        fmt.Printf("middle was removed. OK.\n")
        require.Equal(t, 2, list.Len())
        require.Equal(t, []int{10, 30}, Elements(list.(*List)))
        fmt.Println("Forward stroke check for [10, 30]. OK.")
        for i, v := range [...]int{40, 50, 60, 70, 80} {
            if i%2 == 0 {
                list.PushToLeftEnd(v)
            } else {
                list.PushToRightEnd(v)
            }
        } // [80, 60, 40, 10, 30, 50, 70]
        fmt.Println("List [10, 30] mixing values {40, 50, 60, 70, 80} with mod(2, index).")

        require.Equal(t, 7, list.Len())
        fmt.Printf("list.Len() is %v. OK.\n", list.Len())
        require.Equal(t, 80, list.LeftEnd().Value())
        fmt.Printf("list.LeftEnd().Value() is %v. OK.\n", list.LeftEnd().Value())
        require.Equal(t, 70, list.RightEnd().Value())
        fmt.Printf("list.RightEnd().Value() is %v. OK.\n", list.RightEnd().Value())

        require.Equal(t, []int{80, 60, 40, 10, 30, 50, 70}, Elements(list.(*List)))
        fmt.Println("Forward stroke check for [80, 60, 40, 10, 30, 50, 70]. OK.")

        rightEnd := list.RightEnd()
        list.Remove(rightEnd)
        list.PushToRightEnd(rightEnd.Value())
        fmt.Println("Remove right and push back to right - check for [80, 60, 40, 10, 30, 50, 70]. OK.")
        leftEnd, _ := list.Remove(list.LeftEnd())
        list.PushToLeftEnd(leftEnd.Value())
        fmt.Println("Remove left and push back to left - check for [80, 60, 40, 10, 30, 50, 70]. OK.")
        // - check for nil-refs of first and last
        require.Equal(t, list.LeftEnd().Value(), 80)
        require.Equal(t, list.RightEnd().Value(), 70)
        fmt.Println("Check for list.LeftEnd() is 80 and list.RightEnd() is 70. OK.")
        require.Nil(t, list.LeftEnd().Left())
        require.Nil(t, list.RightEnd().Right())
        fmt.Println("Check for list.LeftEnd().Left() and list.RightEnd().Right() is nils. OK.")

    })
}

func TestListSwap(t *testing.T) {
    t.Run("List swap test.", func(t *testing.T) {
        list := NewList()
        one := list.PushToRightEnd(10)   // [10]
        two := list.PushToRightEnd(20)   // [10, 20]
        three := list.PushToRightEnd(30) // [10, 20, 30]
        four := list.PushToRightEnd(40)  // [10, 20, 30, 40]
        five := list.PushToRightEnd(50)  // [10, 20, 30, 40, 50]
        require.Equal(t, []int{10, 20, 30, 40, 50}, Elements(list.(*List)))
        fmt.Println("[10, 20, 30, 40, 50]. OK.")
        list.Swap(two, four)
        fmt.Println("swap different element-pairs")
        require.Equal(t, []int{10, 40, 30, 20, 50}, Elements(list.(*List)))
        fmt.Println("[10, 40, 30, 20, 50]. OK.")
        list.Swap(two, four)
        require.Equal(t, []int{10, 20, 30, 40, 50}, Elements(list.(*List)))
        fmt.Println("[10, 20, 30, 40, 50]. OK.")
        list.Swap(one, five)
        require.Equal(t, []int{50, 20, 30, 40, 10}, Elements(list.(*List)))
        fmt.Println("[50, 20, 30, 40, 10]. OK.")
        list.Swap(one, five)
        require.Equal(t, []int{10, 20, 30, 40, 50}, Elements(list.(*List)))
        fmt.Println("[10, 20, 30, 40, 50]. OK.")
        list.Swap(one, three)
        require.Equal(t, []int{30, 20, 10, 40, 50}, Elements(list.(*List)))
        fmt.Println("[30, 20, 10, 40, 50]. OK.")
        list.Swap(one, three)
        require.Equal(t, []int{10, 20, 30, 40, 50}, Elements(list.(*List)))
        fmt.Println("[10, 20, 30, 40, 50]. OK.")
        list.Swap(five, two)
        require.Equal(t, []int{10, 50, 30, 40, 20}, Elements(list.(*List)))
        fmt.Println("[10, 50, 30, 40, 20]. OK.")
        list.Swap(two, five)
        require.Equal(t, []int{10, 20, 30, 40, 50}, Elements(list.(*List)))
        fmt.Println("[10, 20, 30, 40, 50]. OK.")
        list.Swap(two, three)
        require.Equal(t, []int{10, 30, 20, 40, 50}, Elements(list.(*List)))
        fmt.Println("[10, 30, 20, 40, 50]. OK.")
        list.Swap(two, three)
        require.Equal(t, []int{10, 20, 30, 40, 50}, Elements(list.(*List)))
        fmt.Println("[10, 20, 30, 40, 50]. OK.")
        list.Swap(four, three)
        require.Equal(t, []int{10, 20, 40, 30, 50}, Elements(list.(*List)))
        fmt.Println("[10, 20, 40, 30, 50]. OK.")
        list.Swap(five, three)
        require.Equal(t, []int{10, 20, 40, 50, 30}, Elements(list.(*List)))
        fmt.Println("[10, 20, 40, 50, 30]. OK.")
        list.Swap(one, two)
        require.Equal(t, []int{20, 10, 40, 50, 30}, Elements(list.(*List)))
        fmt.Println("[20, 10, 40, 50, 30]. OK.")
        list.Swap(one, two)
        require.Equal(t, []int{10, 20, 40, 50, 30}, Elements(list.(*List)))
        fmt.Println("[10, 20, 40, 50, 30]. OK.")
        list.Swap(three, five)
        require.Equal(t, []int{10, 20, 40, 30, 50}, Elements(list.(*List)))
        fmt.Println("[10, 20, 40, 30, 50]. OK.")
        list.Swap(three, four)
        require.Equal(t, []int{10, 20, 30, 40, 50}, Elements(list.(*List)))
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
Item:  0xc00015c2a0
-------------------
Value: <nil>
Left:  0x0
Right: 0x0
-------------------

=== RUN   TestList/ListItem_direct_and_vice_versa_referencies_test.

[1] <--> [2] <--> [3]

[1] is: 
-------------------
Item:  0xc00015c2c0
-------------------
Value: 1
Left:  0x0
Right: 0x0
-------------------

add [2]: 
-------------------
Item:  0xc00015c2e0
-------------------
Value: 2
Left:  0xc00015c2c0
Right: 0x0
-------------------

[1] become: 
-------------------
Item:  0xc00015c2c0
-------------------
Value: 1
Left:  0x0
Right: 0xc00015c2e0
-------------------

add [3]: 
-------------------
Item:  0xc00015c300
-------------------
Value: 3
Left:  0xc00015c2e0
Right: 0x0
-------------------

[2] become: 
-------------------
Item:  0xc00015c2e0
-------------------
Value: 2
Left:  0xc00015c2c0
Right: 0xc00015c300
-------------------

first.Right().Right().Right() is nil. OK.
first.Right().Right() is third. OK.
third.Left().Left().Left() is nil. OK.
third.Left().Left() is first. OK.

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
Item:  0xc00015c320
-------------------
Value: 1
Left:  0x0
Right: 0x0
-------------------

list.LeftEnd() and list.RightEnd() is item. OK.

List become:
 
 (nil:0x0)
    ^|
    L|R
     |v
-------------------
Item:  0xc00015c320
-------------------
Value: 1
Left:  0x0
Right: 0x0
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
Item:  0xc00015c340
-------------------
Value: 1
Left:  0x0
Right: 0x0
-------------------

list.LeftEnd() become:
 
-------------------
Item:  0xc00015c340
-------------------
Value: 1
Left:  0x0
Right: 0x0
-------------------

list.RightEnd() become:
 
-------------------
Item:  0xc00015c340
-------------------
Value: 1
Left:  0x0
Right: 0x0
-------------------

Item [2] become:
 
-------------------
Item:  0xc00015c360
-------------------
Value: 2
Left:  0x0
Right: 0xc00015c340
-------------------

list.LeftEnd() become:
 
-------------------
Item:  0xc00015c360
-------------------
Value: 2
Left:  0x0
Right: 0xc00015c340
-------------------

list.RightEnd() become:
 
-------------------
Item:  0xc00015c340
-------------------
Value: 1
Left:  0xc00015c360
Right: 0x0
-------------------

List become:
 
 (nil:0x0)
    ^|
    L|R
     |v
-------------------
Item:  0xc00015c360
-------------------
Value: 2
Left:  0x0
Right: 0xc00015c340
-------------------
    ^|
    L|R
     |v
-------------------
Item:  0xc00015c340
-------------------
Value: 1
Left:  0xc00015c360
Right: 0x0
-------------------
    ^|
    L|R
     |v
 (nil:0x0)

Was removed:
 
-------------------
Item:  0xc00015c360
-------------------
Value: 2
Left:  0x0
Right: 0xc00015c340
-------------------

List become:
 
 (nil:0x0)
    ^|
    L|R
     |v
-------------------
Item:  0xc00015c340
-------------------
Value: 1
Left:  0x0
Right: 0x0
-------------------
    ^|
    L|R
     |v
 (nil:0x0)

Was removed:
 
-------------------
Item:  0xc00015c340
-------------------
Value: 1
Left:  0x0
Right: 0x0
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
list.Len() is 7. OK.
list.LeftEnd().Value() is 80. OK.
list.RightEnd().Value() is 70. OK.
Forward stroke check for [80, 60, 40, 10, 30, 50, 70]. OK.
Remove right and push back to right - check for [80, 60, 40, 10, 30, 50, 70]. OK.
Remove left and push back to left - check for [80, 60, 40, 10, 30, 50, 70]. OK.
Check for list.LeftEnd() is 80 and list.RightEnd() is 70. OK.
Check for list.LeftEnd().Left() and list.RightEnd().Right() is nils. OK.
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
ok  	command-line-arguments	0.008s

```

```shell
go doc -all ./ > list.doc.txt
```

Документация:

```text


FUNCTIONS

func Eq(x, y ListItem) bool
    Eq(x, y ListItem) - элементs равны тогда и только тогда, когда это один и
    тот же элемент.


TYPES

type List struct {
	// Has unexported fields.
}
    List - структура двусвязного списка.

func (list *List) Contains(item *ListItem) bool
    Contains(item *ListItem) - проверить, есть ли элемент в списке.

func (list *List) LeftEnd() *ListItem
    LeftEnd() - получить элемент из левого края двусвязного списка.

func (list *List) Len() int
    Len() - получить длину двусвязного списка.

func (list *List) PushToLeftEnd(value interface{}) *ListItem
    PushToLeftEnd() - добавить значение в левый край двусвязного списка.

func (list *List) PushToRightEnd(value interface{}) *ListItem
    PushToRightEnd(any interface{}) - добавить значение в правый край
    двусвязного списка.

func (list *List) Remove(item *ListItem) (*ListItem, error)
    Remove(item *ListItem) - удалить элемент из двусвязного списка.

func (list *List) RightEnd() *ListItem
    RightEnd() - получить элемент из правого края двусвязного списка.

func (list *List) String() string
    String - наглядное представление всего двусвязного списка.

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
        Item:  0xc00002e3a0 <--------┐
        -------------------          |
        Value: 2                     |
        Left:  0x0                   |
        Right: 0xc00002e380  >>>-----|---┐ Right 0xc00002e380
        -------------------          |   | ссылается на
            ^|                       |   | блок 0xc00002e380
            L|R                      |   |
             |v                      |   |
        -------------------          |   |
        Item:  0xc00002e380  <-----------┘
        -------------------          | Left 0xc00002e3a0
        Value: 1                     | ссылается на
        Left:  0xc00002e3a0  >>>-----┘ блок 0xc00002e3a0
        Right: 0x0
        -------------------
            ^|
            L|R
             |v
         (nil:0x0)

func (list *List) Swap(x, y *ListItem) error

type ListItem struct {
	// Has unexported fields.
}
    ListItem - структура элемента двусвязного списка.

func (listItem *ListItem) Left() *ListItem
    Left() - получить стоящий слева элемент.

func (listItem *ListItem) Right() *ListItem
    Right() - получить стоящий справа элемент.

func (listItem *ListItem) String() string
    String - наглядное представление значения элемента двусвязного списка.

    Например,

        -------------------             -------------------
        Item:  0xc00002e400             Item: 0xc00002e400
        -------------------             -------------------
        Value: 30                или    Value: 30
        Left:  0xc00002e3c0             Left:  0x0
        Right: 0xc00002e440             Right: 0x0
        -------------------             -------------------

func (listItem *ListItem) Value() interface{}
    Value() - получить значение из элемента.

type Lister interface {
	Len() int
	LeftEnd() *ListItem
	RightEnd() *ListItem
	PushToLeftEnd(value interface{}) *ListItem
	PushToRightEnd(value interface{}) *ListItem
	Remove(item *ListItem) (*ListItem, error)
	Swap(x, y *ListItem) error
}
    Lister - интерфейс двусвязного списка.

func NewList() Lister

type Template struct {
	// Has unexported fields.
}


```

## Послесловие

>
> ```text
> Данный документ составлен с использованием разработанного шаблонизатора. 
> При его использовании избегайте рекурсивной вложенности.
> ```

см. ["Шаблонизатор"](https://github.com/BorisPlus/golang_notes/tree/master/templator)
