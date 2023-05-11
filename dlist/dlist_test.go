package dlist_test

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"

	dlist "github.com/BorisPlus/golang_notes/dlist" 
)

// go test -v dlist.go list_string.go list_test.go

func ElementsLeftToRight(dlist *dlist.DList) []int {
	elements := make([]int, 0, dlist.Len())
	for i := dlist.LeftEdge(); i != nil; i = i.RightNeighbour() {
		elements = append(elements, i.Value().(int))
	}
	return elements
}

func ElementsRightToLeft(dlist *dlist.DList) []int {
	elements := make([]int, 0, dlist.Len())
	for i := dlist.RightEdge(); i != nil; i = i.LeftNeighbour() {
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
		dlist := dlist.NewDList()
		require.Equal(t, 0, dlist.Len())
		require.Nil(t, dlist.LeftEdge())
		require.Nil(t, dlist.RightEdge())
		fmt.Println("\nList is:\n", dlist)
		fmt.Println()
	})

	t.Run("DList init test.", func(t *testing.T) {
		dlist := dlist.NewDList()

		fmt.Println("\nList was:\n", dlist)

		item := dlist.PushToLeftEdge(1)
		fmt.Println("\nItem was:", item)
		require.Equal(t, item, dlist.LeftEdge())
		require.Equal(t, item, dlist.RightEdge())
		fmt.Println()
		fmt.Println("dlist.LeftEdge() and dlist.RightEdge() is item. OK.")
		fmt.Println("\nList become:\n", dlist)
		fmt.Println()
	})

	t.Run("Little DList test.", func(t *testing.T) {
		dlist := dlist.NewDList()
		fmt.Println("\nList was:\n", dlist)

		itemFirst := dlist.PushToLeftEdge(1) // [1]
		fmt.Println("\nItem [1] become:\n", itemFirst)

		leftEnd := dlist.LeftEdge()
		fmt.Println("\nlist.LeftEdge() become:\n", leftEnd)
		rightEnd := dlist.RightEdge()
		fmt.Println("\nlist.RightEdge() become:\n", rightEnd)

		itemSecond := dlist.PushToLeftEdge(2) // [2]
		fmt.Println("\nItem [2] become:\n", itemSecond)

		leftEnd = dlist.LeftEdge()
		fmt.Println("\nlist.LeftEdge() become:\n", leftEnd)
		rightEnd = dlist.RightEdge()
		fmt.Println("\nlist.RightEdge() become:\n", rightEnd)
		fmt.Println("\nList become:\n", dlist)
		require.Equal(t, itemSecond, dlist.LeftEdge())
		require.Equal(t, itemFirst, dlist.RightEdge())

		someItem, _ := dlist.Remove(dlist.LeftEdge())
		fmt.Println("\nWas removed:\n", someItem)
		require.Equal(t, itemSecond, someItem)

		fmt.Println("\nList become:\n", dlist)

		require.Equal(t, dlist.LeftEdge(), dlist.RightEdge())
		require.Nil(t, dlist.LeftEdge().LeftNeighbour())
		require.Nil(t, dlist.LeftEdge().RightNeighbour())

		someItem, _ = dlist.Remove(dlist.LeftEdge())
		fmt.Println("\nWas removed:\n", someItem)
		require.Equal(t, itemFirst, someItem)

		fmt.Println("\nList become:\n", dlist)

		require.Equal(t, dlist.LeftEdge(), dlist.RightEdge())
		require.Nil(t, dlist.LeftEdge())
		require.Nil(t, dlist.RightEdge())
	})
}

func TestDListComplex(t *testing.T) {
	t.Run("Сomplex DList processing test.", func(t *testing.T) {
		dlist := dlist.NewDList()
		dlist.PushToRightEdge(10) // [10]
		dlist.PushToRightEdge(20) // [10, 20]
		dlist.PushToRightEdge(30) // [10, 20, 30]
		require.Equal(t, []int{10, 20, 30}, ElementsLeftToRight(dlist.(*dlist.DList)))
		require.Equal(t, []int{10, 20, 30}, ElementsRightToLeft(dlist.(*dlist.DList)))
		fmt.Println("Forward stroke check for [10, 20, 30]. OK.")
		require.Equal(t, 3, dlist.Len())
		middle := dlist.LeftEdge().RightNeighbour()
		require.Equal(t, middle.Value(), 20)
		fmt.Printf("middle.Value() is %v. OK.\n", middle.Value())
		dlist.Remove(middle)
		fmt.Printf("middle was removed. OK.\n")
		require.Equal(t, 2, dlist.Len())
		require.Equal(t, []int{10, 30}, ElementsLeftToRight(dlist.(*dlist.DList)))
		require.Equal(t, []int{10, 30}, ElementsRightToLeft(dlist.(*dlist.DList)))
		fmt.Println("Forward stroke check for [10, 30]. OK.")
		for i, v := range [...]int{40, 50, 60, 70, 80} {
			if i%2 == 0 {
				dlist.PushToLeftEdge(v)
			} else {
				dlist.PushToRightEdge(v)
			}
		} // [80, 60, 40, 10, 30, 50, 70]
		fmt.Println("DList [10, 30] mixing values {40, 50, 60, 70, 80} with mod(2, index).")

		require.Equal(t, 7, dlist.Len())
		fmt.Printf("dlist.Len() is %v. OK.\n", dlist.Len())
		require.Equal(t, 80, dlist.LeftEdge().Value())
		fmt.Printf("dlist.LeftEdge().Value() is %v. OK.\n", dlist.LeftEdge().Value())
		require.Equal(t, 70, dlist.RightEdge().Value())
		fmt.Printf("dlist.RightEdge().Value() is %v. OK.\n", dlist.RightEdge().Value())

		require.Equal(t, []int{80, 60, 40, 10, 30, 50, 70}, ElementsLeftToRight(dlist.(*dlist.DList)))
		require.Equal(t, []int{80, 60, 40, 10, 30, 50, 70}, ElementsRightToLeft(dlist.(*dlist.DList)))
		fmt.Println("Forward stroke check for [80, 60, 40, 10, 30, 50, 70]. OK.")

		rightEnd := dlist.RightEdge()
		dlist.Remove(rightEnd)
		dlist.PushToRightEdge(rightEnd.Value())
		fmt.Println("Remove right and push back to right - check for [80, 60, 40, 10, 30, 50, 70]. OK.")
		leftEnd, _ := dlist.Remove(dlist.LeftEdge())
		dlist.PushToLeftEdge(leftEnd.Value())
		fmt.Println("Remove left and push back to left - check for [80, 60, 40, 10, 30, 50, 70]. OK.")
		// - check for nil-refs of first and last
		require.Equal(t, dlist.LeftEdge().Value(), 80)
		require.Equal(t, dlist.RightEdge().Value(), 70)
		fmt.Println("Check for dlist.LeftEdge() is 80 and dlist.RightEdge() is 70. OK.")
		require.Nil(t, dlist.LeftEdge().LeftNeighbour())
		require.Nil(t, dlist.RightEdge().RightNeighbour())
		fmt.Println("Check for dlist.LeftEdge().Left() and dlist.RightEdge().Right() is nils. OK.")

	})
}

func TestDListSwap(t *testing.T) {
	t.Run("DList swap test.", func(t *testing.T) {
		dlist := dlist.NewDList()
		one := dlist.PushToRightEdge(10)   // [10]
		two := dlist.PushToRightEdge(20)   // [10, 20]
		three := dlist.PushToRightEdge(30) // [10, 20, 30]
		four := dlist.PushToRightEdge(40)  // [10, 20, 30, 40]
		five := dlist.PushToRightEdge(50)  // [10, 20, 30, 40, 50]
		require.Equal(t, []int{10, 20, 30, 40, 50}, ElementsLeftToRight(dlist.(*dlist.DList)))
		require.Equal(t, []int{10, 20, 30, 40, 50}, ElementsRightToLeft(dlist.(*dlist.DList)))
		fmt.Println("[10, 20, 30, 40, 50]. OK.")
		dlist.SwapItems(two, four)
		fmt.Println("swap different element-pairs")
		require.Equal(t, []int{10, 40, 30, 20, 50}, ElementsLeftToRight(dlist.(*dlist.DList)))
		require.Equal(t, []int{10, 40, 30, 20, 50}, ElementsRightToLeft(dlist.(*dlist.DList)))
		fmt.Println("[10, 40, 30, 20, 50]. OK.")
		dlist.SwapItems(two, four)
		require.Equal(t, []int{10, 20, 30, 40, 50}, ElementsLeftToRight(dlist.(*dlist.DList)))
		require.Equal(t, []int{10, 20, 30, 40, 50}, ElementsRightToLeft(dlist.(*dlist.DList)))
		fmt.Println("[10, 20, 30, 40, 50]. OK.")
		dlist.SwapItems(one, five)
		require.Equal(t, []int{50, 20, 30, 40, 10}, ElementsLeftToRight(dlist.(*dlist.DList)))
		require.Equal(t, []int{50, 20, 30, 40, 10}, ElementsRightToLeft(dlist.(*dlist.DList)))
		fmt.Println("[50, 20, 30, 40, 10]. OK.")
		dlist.SwapItems(one, five)
		require.Equal(t, []int{10, 20, 30, 40, 50}, ElementsLeftToRight(dlist.(*dlist.DList)))
		require.Equal(t, []int{10, 20, 30, 40, 50}, ElementsRightToLeft(dlist.(*dlist.DList)))
		fmt.Println("[10, 20, 30, 40, 50]. OK.")
		dlist.SwapItems(one, three)
		require.Equal(t, []int{30, 20, 10, 40, 50}, ElementsLeftToRight(dlist.(*dlist.DList)))
		require.Equal(t, []int{30, 20, 10, 40, 50}, ElementsRightToLeft(dlist.(*dlist.DList)))
		fmt.Println("[30, 20, 10, 40, 50]. OK.")
		dlist.SwapItems(one, three)
		require.Equal(t, []int{10, 20, 30, 40, 50}, ElementsLeftToRight(dlist.(*dlist.DList)))
		require.Equal(t, []int{10, 20, 30, 40, 50}, ElementsRightToLeft(dlist.(*dlist.DList)))
		fmt.Println("[10, 20, 30, 40, 50]. OK.")
		dlist.SwapItems(five, two)
		require.Equal(t, []int{10, 50, 30, 40, 20}, ElementsLeftToRight(dlist.(*dlist.DList)))
		require.Equal(t, []int{10, 50, 30, 40, 20}, ElementsRightToLeft(dlist.(*dlist.DList)))
		fmt.Println("[10, 50, 30, 40, 20]. OK.")
		dlist.SwapItems(two, five)
		require.Equal(t, []int{10, 20, 30, 40, 50}, ElementsLeftToRight(dlist.(*dlist.DList)))
		require.Equal(t, []int{10, 20, 30, 40, 50}, ElementsRightToLeft(dlist.(*dlist.DList)))
		fmt.Println("[10, 20, 30, 40, 50]. OK.")
		dlist.SwapItems(two, three)
		require.Equal(t, []int{10, 30, 20, 40, 50}, ElementsLeftToRight(dlist.(*dlist.DList)))
		require.Equal(t, []int{10, 30, 20, 40, 50}, ElementsRightToLeft(dlist.(*dlist.DList)))
		fmt.Println("[10, 30, 20, 40, 50]. OK.")
		dlist.SwapItems(two, three)
		require.Equal(t, []int{10, 20, 30, 40, 50}, ElementsLeftToRight(dlist.(*dlist.DList)))
		require.Equal(t, []int{10, 20, 30, 40, 50}, ElementsRightToLeft(dlist.(*dlist.DList)))
		fmt.Println("[10, 20, 30, 40, 50]. OK.")
		dlist.SwapItems(four, three)
		require.Equal(t, []int{10, 20, 40, 30, 50}, ElementsLeftToRight(dlist.(*dlist.DList)))
		require.Equal(t, []int{10, 20, 40, 30, 50}, ElementsRightToLeft(dlist.(*dlist.DList)))
		fmt.Println("[10, 20, 40, 30, 50]. OK.")
		dlist.SwapItems(five, three)
		require.Equal(t, []int{10, 20, 40, 50, 30}, ElementsLeftToRight(dlist.(*dlist.DList)))
		require.Equal(t, []int{10, 20, 40, 50, 30}, ElementsRightToLeft(dlist.(*dlist.DList)))
		fmt.Println("[10, 20, 40, 50, 30]. OK.")
		dlist.SwapItems(one, two)
		require.Equal(t, []int{20, 10, 40, 50, 30}, ElementsLeftToRight(dlist.(*dlist.DList)))
		require.Equal(t, []int{20, 10, 40, 50, 30}, ElementsRightToLeft(dlist.(*dlist.DList)))
		fmt.Println("[20, 10, 40, 50, 30]. OK.")
		dlist.SwapItems(one, two)
		require.Equal(t, []int{10, 20, 40, 50, 30}, ElementsLeftToRight(dlist.(*dlist.DList)))
		require.Equal(t, []int{10, 20, 40, 50, 30}, ElementsRightToLeft(dlist.(*dlist.DList)))
		fmt.Println("[10, 20, 40, 50, 30]. OK.")
		dlist.SwapItems(three, five)
		require.Equal(t, []int{10, 20, 40, 30, 50}, ElementsLeftToRight(dlist.(*dlist.DList)))
		require.Equal(t, []int{10, 20, 40, 30, 50}, ElementsRightToLeft(dlist.(*dlist.DList)))
		fmt.Println("[10, 20, 40, 30, 50]. OK.")
		dlist.SwapItems(three, four)
		require.Equal(t, []int{10, 20, 30, 40, 50}, ElementsLeftToRight(dlist.(*dlist.DList)))
		require.Equal(t, []int{10, 20, 30, 40, 50}, ElementsRightToLeft(dlist.(*dlist.DList)))
		fmt.Println("[10, 20, 30, 40, 50]. OK.")
	})
}


func TestDListSortInterface(t *testing.T) {

	t.Run("Let's sort Dlist.", func(t *testing.T) {
		dlist := dlist.NewDList()
		dlist.PushToRightEdge(10) // [10]
		dlist.PushToRightEdge(30) // [10, 30]
		dlist.PushToRightEdge(20) // [10, 30, 20]
		dlist.PushToRightEdge(50) // [10, 30, 20, 50]
		dlist.PushToRightEdge(40) // [10, 30, 20, 50, 40]
		sample := []int{10, 30, 20, 50, 40}
		require.Equal(t, sample, ElementsLeftToRight(dlist.(*dlist.DList)))
		require.Equal(t, sample, ElementsRightToLeft(dlist.(*dlist.DList)))
		fmt.Printf("\nTest dlist before sort: %v. OK.\n", sample)
		fmt.Printf("\nList before sort with Stringer() formatting:\n%s\n", dlist)
		sort.Sort(dlist.(*dlist.DList))
		expected := []int{10, 20, 30, 40, 50}
		require.Equal(t, expected, ElementsLeftToRight(dlist.(*dlist.DList)))
		require.Equal(t, expected, ElementsRightToLeft(dlist.(*dlist.DList)))
		fmt.Printf("\nTest dlist after sort: %v. OK.\n", expected)
		fmt.Printf("\nList after sort with Stringer() formatting:\n%s\n", dlist)
	})
}
