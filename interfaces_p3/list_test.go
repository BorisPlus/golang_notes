package interfaces_p3_test

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"

	dlist "github.com/BorisPlus/golang_notes/interfaces_p3" 
)

// go test -v list.go list_string.go list_test.go

func ElementsLeftToRight(list *dlist.List) []int {
	elements := make([]int, 0, list.Len())
	for i := list.LeftEdge(); i != nil; i = i.RightNeighbour() {
		elements = append(elements, i.Value().(int))
	}
	return elements
}

func ElementsRightToLeft(list *dlist.List) []int {
	elements := make([]int, 0, list.Len())
	for i := list.RightEdge(); i != nil; i = i.LeftNeighbour() {
		elements = append([]int{i.Value().(int)}, elements...)
	}
	return elements
}

func TestListSimple(t *testing.T) {
	t.Run("Zero-value ListItem test.", func(t *testing.T) {
		zeroValueItem := dlist.ListItem{}
		require.Nil(t, zeroValueItem.Value())
		require.Nil(t, zeroValueItem.LeftNeighbour())
		require.Nil(t, zeroValueItem.RightNeighbour())
		fmt.Println("\n[zero-value] is:", &zeroValueItem)
		fmt.Println()
	})

	t.Run("ListItem direct and vice versa referencies test.", func(t *testing.T) {
		fmt.Println("\n[1] <--> [2] <--> [3]")
		first := &dlist.ListItem{}
		first.SetValue(1)
		fmt.Println("\n[1] is:", first)

		second := &dlist.ListItem{}
		second.SetValue(2)
		second.SetLeftNeighbour(first)
		first.SetRightNeighbour(second)
		fmt.Println("\nadd [2]:", second)
		fmt.Println("\n[1] become:", first)

		third := dlist.ListItem{}
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

	t.Run("Empty List test.", func(t *testing.T) {
		list := dlist.NewList()
		require.Equal(t, 0, list.Len())
		require.Nil(t, list.LeftEdge())
		require.Nil(t, list.RightEdge())
		fmt.Println("\nList is:\n", list)
		fmt.Println()
	})

	t.Run("List init test.", func(t *testing.T) {
		list := dlist.NewList()

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
		list := dlist.NewList()
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
		list := dlist.NewList()
		list.PushToRightEdge(10) // [10]
		list.PushToRightEdge(20) // [10, 20]
		list.PushToRightEdge(30) // [10, 20, 30]
		require.Equal(t, []int{10, 20, 30}, ElementsLeftToRight(list.(*dlist.List)))
		require.Equal(t, []int{10, 20, 30}, ElementsRightToLeft(list.(*dlist.List)))
		fmt.Println("Forward stroke check for [10, 20, 30]. OK.")
		require.Equal(t, 3, list.Len())
		middle := list.LeftEdge().RightNeighbour()
		require.Equal(t, middle.Value(), 20)
		fmt.Printf("middle.Value() is %v. OK.\n", middle.Value())
		list.Remove(middle)
		fmt.Printf("middle was removed. OK.\n")
		require.Equal(t, 2, list.Len())
		require.Equal(t, []int{10, 30}, ElementsLeftToRight(list.(*dlist.List)))
		require.Equal(t, []int{10, 30}, ElementsRightToLeft(list.(*dlist.List)))
		fmt.Println("Forward stroke check for [10, 30]. OK.")
		for i, v := range [...]int{40, 50, 60, 70, 80} {
			if i%2 == 0 {
				list.PushToLeftEdge(v)
			} else {
				list.PushToRightEdge(v)
			}
		} // [80, 60, 40, 10, 30, 50, 70]
		fmt.Println("List [10, 30] mixing values {40, 50, 60, 70, 80} with mod(2, index).")

		require.Equal(t, 7, list.Len())
		fmt.Printf("list.Len() is %v. OK.\n", list.Len())
		require.Equal(t, 80, list.LeftEdge().Value())
		fmt.Printf("list.LeftEdge().Value() is %v. OK.\n", list.LeftEdge().Value())
		require.Equal(t, 70, list.RightEdge().Value())
		fmt.Printf("list.RightEdge().Value() is %v. OK.\n", list.RightEdge().Value())

		require.Equal(t, []int{80, 60, 40, 10, 30, 50, 70}, ElementsLeftToRight(list.(*dlist.List)))
		require.Equal(t, []int{80, 60, 40, 10, 30, 50, 70}, ElementsRightToLeft(list.(*dlist.List)))
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
		list := dlist.NewList()
		one := list.PushToRightEdge(10)   // [10]
		two := list.PushToRightEdge(20)   // [10, 20]
		three := list.PushToRightEdge(30) // [10, 20, 30]
		four := list.PushToRightEdge(40)  // [10, 20, 30, 40]
		five := list.PushToRightEdge(50)  // [10, 20, 30, 40, 50]
		require.Equal(t, []int{10, 20, 30, 40, 50}, ElementsLeftToRight(list.(*dlist.List)))
		require.Equal(t, []int{10, 20, 30, 40, 50}, ElementsRightToLeft(list.(*dlist.List)))
		fmt.Println("[10, 20, 30, 40, 50]. OK.")
		list.SwapItems(two, four)
		fmt.Println("swap different element-pairs")
		require.Equal(t, []int{10, 40, 30, 20, 50}, ElementsLeftToRight(list.(*dlist.List)))
		require.Equal(t, []int{10, 40, 30, 20, 50}, ElementsRightToLeft(list.(*dlist.List)))
		fmt.Println("[10, 40, 30, 20, 50]. OK.")
		list.SwapItems(two, four)
		require.Equal(t, []int{10, 20, 30, 40, 50}, ElementsLeftToRight(list.(*dlist.List)))
		require.Equal(t, []int{10, 20, 30, 40, 50}, ElementsRightToLeft(list.(*dlist.List)))
		fmt.Println("[10, 20, 30, 40, 50]. OK.")
		list.SwapItems(one, five)
		require.Equal(t, []int{50, 20, 30, 40, 10}, ElementsLeftToRight(list.(*dlist.List)))
		require.Equal(t, []int{50, 20, 30, 40, 10}, ElementsRightToLeft(list.(*dlist.List)))
		fmt.Println("[50, 20, 30, 40, 10]. OK.")
		list.SwapItems(one, five)
		require.Equal(t, []int{10, 20, 30, 40, 50}, ElementsLeftToRight(list.(*dlist.List)))
		require.Equal(t, []int{10, 20, 30, 40, 50}, ElementsRightToLeft(list.(*dlist.List)))
		fmt.Println("[10, 20, 30, 40, 50]. OK.")
		list.SwapItems(one, three)
		require.Equal(t, []int{30, 20, 10, 40, 50}, ElementsLeftToRight(list.(*dlist.List)))
		require.Equal(t, []int{30, 20, 10, 40, 50}, ElementsRightToLeft(list.(*dlist.List)))
		fmt.Println("[30, 20, 10, 40, 50]. OK.")
		list.SwapItems(one, three)
		require.Equal(t, []int{10, 20, 30, 40, 50}, ElementsLeftToRight(list.(*dlist.List)))
		require.Equal(t, []int{10, 20, 30, 40, 50}, ElementsRightToLeft(list.(*dlist.List)))
		fmt.Println("[10, 20, 30, 40, 50]. OK.")
		list.SwapItems(five, two)
		require.Equal(t, []int{10, 50, 30, 40, 20}, ElementsLeftToRight(list.(*dlist.List)))
		require.Equal(t, []int{10, 50, 30, 40, 20}, ElementsRightToLeft(list.(*dlist.List)))
		fmt.Println("[10, 50, 30, 40, 20]. OK.")
		list.SwapItems(two, five)
		require.Equal(t, []int{10, 20, 30, 40, 50}, ElementsLeftToRight(list.(*dlist.List)))
		require.Equal(t, []int{10, 20, 30, 40, 50}, ElementsRightToLeft(list.(*dlist.List)))
		fmt.Println("[10, 20, 30, 40, 50]. OK.")
		list.SwapItems(two, three)
		require.Equal(t, []int{10, 30, 20, 40, 50}, ElementsLeftToRight(list.(*dlist.List)))
		require.Equal(t, []int{10, 30, 20, 40, 50}, ElementsRightToLeft(list.(*dlist.List)))
		fmt.Println("[10, 30, 20, 40, 50]. OK.")
		list.SwapItems(two, three)
		require.Equal(t, []int{10, 20, 30, 40, 50}, ElementsLeftToRight(list.(*dlist.List)))
		require.Equal(t, []int{10, 20, 30, 40, 50}, ElementsRightToLeft(list.(*dlist.List)))
		fmt.Println("[10, 20, 30, 40, 50]. OK.")
		list.SwapItems(four, three)
		require.Equal(t, []int{10, 20, 40, 30, 50}, ElementsLeftToRight(list.(*dlist.List)))
		require.Equal(t, []int{10, 20, 40, 30, 50}, ElementsRightToLeft(list.(*dlist.List)))
		fmt.Println("[10, 20, 40, 30, 50]. OK.")
		list.SwapItems(five, three)
		require.Equal(t, []int{10, 20, 40, 50, 30}, ElementsLeftToRight(list.(*dlist.List)))
		require.Equal(t, []int{10, 20, 40, 50, 30}, ElementsRightToLeft(list.(*dlist.List)))
		fmt.Println("[10, 20, 40, 50, 30]. OK.")
		list.SwapItems(one, two)
		require.Equal(t, []int{20, 10, 40, 50, 30}, ElementsLeftToRight(list.(*dlist.List)))
		require.Equal(t, []int{20, 10, 40, 50, 30}, ElementsRightToLeft(list.(*dlist.List)))
		fmt.Println("[20, 10, 40, 50, 30]. OK.")
		list.SwapItems(one, two)
		require.Equal(t, []int{10, 20, 40, 50, 30}, ElementsLeftToRight(list.(*dlist.List)))
		require.Equal(t, []int{10, 20, 40, 50, 30}, ElementsRightToLeft(list.(*dlist.List)))
		fmt.Println("[10, 20, 40, 50, 30]. OK.")
		list.SwapItems(three, five)
		require.Equal(t, []int{10, 20, 40, 30, 50}, ElementsLeftToRight(list.(*dlist.List)))
		require.Equal(t, []int{10, 20, 40, 30, 50}, ElementsRightToLeft(list.(*dlist.List)))
		fmt.Println("[10, 20, 40, 30, 50]. OK.")
		list.SwapItems(three, four)
		require.Equal(t, []int{10, 20, 30, 40, 50}, ElementsLeftToRight(list.(*dlist.List)))
		require.Equal(t, []int{10, 20, 30, 40, 50}, ElementsRightToLeft(list.(*dlist.List)))
		fmt.Println("[10, 20, 30, 40, 50]. OK.")
	})
}


func TestListSortInterface(t *testing.T) {

	t.Run("Let's sort double-linked list.", func(t *testing.T) {
		list := dlist.NewList()
		list.PushToRightEdge(10) // [10]
		list.PushToRightEdge(30) // [10, 30]
		list.PushToRightEdge(20) // [10, 30, 20]
		list.PushToRightEdge(50) // [10, 30, 20, 50]
		list.PushToRightEdge(40) // [10, 30, 20, 50, 40]
		sample := []int{10, 30, 20, 50, 40}
		require.Equal(t, sample, ElementsLeftToRight(list.(*dlist.List)))
		require.Equal(t, sample, ElementsRightToLeft(list.(*dlist.List)))
		fmt.Printf("\nTest list before sort: %v. OK.\n", sample)
		fmt.Printf("\nList before sort with Stringer() formatting:\n%s\n", list)
		sort.Sort(list.(*dlist.List))
		expected := []int{10, 20, 30, 40, 50}
		require.Equal(t, expected, ElementsLeftToRight(list.(*dlist.List)))
		require.Equal(t, expected, ElementsRightToLeft(list.(*dlist.List)))
		fmt.Printf("\nTest list after sort: %v. OK.\n", expected)
		fmt.Printf("\nList after sort with Stringer() formatting:\n%s\n", list)
	})
}
