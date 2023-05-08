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
