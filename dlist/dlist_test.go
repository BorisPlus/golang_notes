package dlist_test

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"

	dlist "github.com/BorisPlus/golang_notes/dlist"
)

func ElementsLeftToRightInt(dList *dlist.DList) []int {
	elements := make([]int, 0, dList.Len())
	for i := dList.LeftEdge(); i != nil; i = i.RightNeighbour() {
		elements = append(elements, i.Value().(int))
	}
	return elements
}

func ElementsRightToLeftInt(dList *dlist.DList) []int {
	elements := make([]int, 0, dList.Len())
	for i := dList.RightEdge(); i != nil; i = i.LeftNeighbour() {
		elements = append([]int{i.Value().(int)}, elements...)
	}
	return elements
}

func ElementsLeftToRightRune(dList *dlist.DList) []rune {
	elements := make([]rune, 0, dList.Len())
	for i := dList.LeftEdge(); i != nil; i = i.RightNeighbour() {
		elements = append(elements, i.Value().(rune))
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
		require.Equal(t, []int{10, 20, 30}, ElementsLeftToRightInt(dList.(*dlist.DList)))
		require.Equal(t, []int{10, 20, 30}, ElementsRightToLeftInt(dList.(*dlist.DList)))
		fmt.Println("Forward stroke check for [10, 20, 30]. OK.")
		require.Equal(t, 3, dList.Len())
		middle := dList.LeftEdge().RightNeighbour()
		require.Equal(t, middle.Value(), 20)
		fmt.Printf("middle.Value() is %v. OK.\n", middle.Value())
		dList.Remove(middle)
		fmt.Printf("middle was removed. OK.\n")
		require.Equal(t, 2, dList.Len())
		require.Equal(t, []int{10, 30}, ElementsLeftToRightInt(dList.(*dlist.DList)))
		require.Equal(t, []int{10, 30}, ElementsRightToLeftInt(dList.(*dlist.DList)))
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

		require.Equal(t, []int{80, 60, 40, 10, 30, 50, 70}, ElementsLeftToRightInt(dList.(*dlist.DList)))
		require.Equal(t, []int{80, 60, 40, 10, 30, 50, 70}, ElementsRightToLeftInt(dList.(*dlist.DList)))
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
		require.Equal(t, []int{10, 20, 30, 40, 50}, ElementsLeftToRightInt(dList.(*dlist.DList)))
		require.Equal(t, []int{10, 20, 30, 40, 50}, ElementsRightToLeftInt(dList.(*dlist.DList)))
		fmt.Println("[10, 20, 30, 40, 50]. OK.")
		dList.SwapItems(two, four)
		fmt.Println("swap different element-pairs")
		require.Equal(t, []int{10, 40, 30, 20, 50}, ElementsLeftToRightInt(dList.(*dlist.DList)))
		require.Equal(t, []int{10, 40, 30, 20, 50}, ElementsRightToLeftInt(dList.(*dlist.DList)))
		fmt.Println("[10, 40, 30, 20, 50]. OK.")
		dList.SwapItems(two, four)
		require.Equal(t, []int{10, 20, 30, 40, 50}, ElementsLeftToRightInt(dList.(*dlist.DList)))
		require.Equal(t, []int{10, 20, 30, 40, 50}, ElementsRightToLeftInt(dList.(*dlist.DList)))
		fmt.Println("[10, 20, 30, 40, 50]. OK.")
		dList.SwapItems(one, five)
		require.Equal(t, []int{50, 20, 30, 40, 10}, ElementsLeftToRightInt(dList.(*dlist.DList)))
		require.Equal(t, []int{50, 20, 30, 40, 10}, ElementsRightToLeftInt(dList.(*dlist.DList)))
		fmt.Println("[50, 20, 30, 40, 10]. OK.")
		dList.SwapItems(one, five)
		require.Equal(t, []int{10, 20, 30, 40, 50}, ElementsLeftToRightInt(dList.(*dlist.DList)))
		require.Equal(t, []int{10, 20, 30, 40, 50}, ElementsRightToLeftInt(dList.(*dlist.DList)))
		fmt.Println("[10, 20, 30, 40, 50]. OK.")
		dList.SwapItems(one, three)
		require.Equal(t, []int{30, 20, 10, 40, 50}, ElementsLeftToRightInt(dList.(*dlist.DList)))
		require.Equal(t, []int{30, 20, 10, 40, 50}, ElementsRightToLeftInt(dList.(*dlist.DList)))
		fmt.Println("[30, 20, 10, 40, 50]. OK.")
		dList.SwapItems(one, three)
		require.Equal(t, []int{10, 20, 30, 40, 50}, ElementsLeftToRightInt(dList.(*dlist.DList)))
		require.Equal(t, []int{10, 20, 30, 40, 50}, ElementsRightToLeftInt(dList.(*dlist.DList)))
		fmt.Println("[10, 20, 30, 40, 50]. OK.")
		dList.SwapItems(five, two)
		require.Equal(t, []int{10, 50, 30, 40, 20}, ElementsLeftToRightInt(dList.(*dlist.DList)))
		require.Equal(t, []int{10, 50, 30, 40, 20}, ElementsRightToLeftInt(dList.(*dlist.DList)))
		fmt.Println("[10, 50, 30, 40, 20]. OK.")
		dList.SwapItems(two, five)
		require.Equal(t, []int{10, 20, 30, 40, 50}, ElementsLeftToRightInt(dList.(*dlist.DList)))
		require.Equal(t, []int{10, 20, 30, 40, 50}, ElementsRightToLeftInt(dList.(*dlist.DList)))
		fmt.Println("[10, 20, 30, 40, 50]. OK.")
		dList.SwapItems(two, three)
		require.Equal(t, []int{10, 30, 20, 40, 50}, ElementsLeftToRightInt(dList.(*dlist.DList)))
		require.Equal(t, []int{10, 30, 20, 40, 50}, ElementsRightToLeftInt(dList.(*dlist.DList)))
		fmt.Println("[10, 30, 20, 40, 50]. OK.")
		dList.SwapItems(two, three)
		require.Equal(t, []int{10, 20, 30, 40, 50}, ElementsLeftToRightInt(dList.(*dlist.DList)))
		require.Equal(t, []int{10, 20, 30, 40, 50}, ElementsRightToLeftInt(dList.(*dlist.DList)))
		fmt.Println("[10, 20, 30, 40, 50]. OK.")
		dList.SwapItems(four, three)
		require.Equal(t, []int{10, 20, 40, 30, 50}, ElementsLeftToRightInt(dList.(*dlist.DList)))
		require.Equal(t, []int{10, 20, 40, 30, 50}, ElementsRightToLeftInt(dList.(*dlist.DList)))
		fmt.Println("[10, 20, 40, 30, 50]. OK.")
		dList.SwapItems(five, three)
		require.Equal(t, []int{10, 20, 40, 50, 30}, ElementsLeftToRightInt(dList.(*dlist.DList)))
		require.Equal(t, []int{10, 20, 40, 50, 30}, ElementsRightToLeftInt(dList.(*dlist.DList)))
		fmt.Println("[10, 20, 40, 50, 30]. OK.")
		dList.SwapItems(one, two)
		require.Equal(t, []int{20, 10, 40, 50, 30}, ElementsLeftToRightInt(dList.(*dlist.DList)))
		require.Equal(t, []int{20, 10, 40, 50, 30}, ElementsRightToLeftInt(dList.(*dlist.DList)))
		fmt.Println("[20, 10, 40, 50, 30]. OK.")
		dList.SwapItems(one, two)
		require.Equal(t, []int{10, 20, 40, 50, 30}, ElementsLeftToRightInt(dList.(*dlist.DList)))
		require.Equal(t, []int{10, 20, 40, 50, 30}, ElementsRightToLeftInt(dList.(*dlist.DList)))
		fmt.Println("[10, 20, 40, 50, 30]. OK.")
		dList.SwapItems(three, five)
		require.Equal(t, []int{10, 20, 40, 30, 50}, ElementsLeftToRightInt(dList.(*dlist.DList)))
		require.Equal(t, []int{10, 20, 40, 30, 50}, ElementsRightToLeftInt(dList.(*dlist.DList)))
		fmt.Println("[10, 20, 40, 30, 50]. OK.")
		dList.SwapItems(three, four)
		require.Equal(t, []int{10, 20, 30, 40, 50}, ElementsLeftToRightInt(dList.(*dlist.DList)))
		require.Equal(t, []int{10, 20, 30, 40, 50}, ElementsRightToLeftInt(dList.(*dlist.DList)))
		fmt.Println("[10, 20, 30, 40, 50]. OK.")
	})
}

func TestDListSortInterface(t *testing.T) {

	t.Run("Let's sort int-Dlist.", func(t *testing.T) {
		dList := dlist.NewDList()
		dList.PushToRightEdge(10) // [10]
		dList.PushToRightEdge(30) // [10, 30]
		dList.PushToRightEdge(20) // [10, 30, 20]
		dList.PushToRightEdge(50) // [10, 30, 20, 50]
		dList.PushToRightEdge(40) // [10, 30, 20, 50, 40]
		sample := []int{10, 30, 20, 50, 40}
		require.Equal(t, sample, ElementsLeftToRightInt(dList.(*dlist.DList)))
		require.Equal(t, sample, ElementsRightToLeftInt(dList.(*dlist.DList)))
		fmt.Printf("\nTest dList before sort: %v. OK.\n", sample)
		fmt.Printf("\nList before sort with Stringer() formatting:\n%s\n", dList)
		dList.SetLessItems(func(x, y *dlist.DListItem) bool {
			return x.Value().(int) < y.Value().(int)
		})
		sort.Sort(dList.(*dlist.DList))
		expected := []int{10, 20, 30, 40, 50}
		require.Equal(t, expected, ElementsLeftToRightInt(dList.(*dlist.DList)))
		require.Equal(t, expected, ElementsRightToLeftInt(dList.(*dlist.DList)))
		fmt.Printf("\nTest dList after sort: %v. OK.\n", expected)
		fmt.Printf("\nList after sort with Stringer() formatting:\n%s\n", dList)
	})
	t.Run("Let's sort rune-Dlist.", func(t *testing.T) {
		dList := dlist.NewDList()
		dList.PushToRightEdge('a') // [a]
		dList.PushToRightEdge('c') // [a, c]
		dList.PushToRightEdge('b') // [a, c, b]
		dList.PushToRightEdge('e') // [a, c, b, e]
		dList.PushToRightEdge('d') // [a, c, b, e, d]
		sample := []rune{'a', 'c', 'b', 'e', 'd'}
		require.Equal(t, sample, ElementsLeftToRightRune(dList.(*dlist.DList)))
		fmt.Printf("\nTest dList before sort: %v. OK.\n", sample)
		fmt.Printf("\nList before sort with Stringer() formatting:\n%s\n", dList)
		dList.SetLessItems(func(x, y *dlist.DListItem) bool {
			return x.Value().(rune) < y.Value().(rune)
		})
		sort.Sort(dList.(*dlist.DList))
		expected := []int32{'a', 'b', 'c', 'd', 'e'}
		require.Equal(t, expected, ElementsLeftToRightRune(dList.(*dlist.DList)))
		fmt.Printf("\nTest dList after sort: %v. OK.\n", expected)
		fmt.Printf("\nList after sort with Stringer() formatting:\n%s\n", dList)
	})
	t.Run("Let's sort struct-Dlist.", func(t *testing.T) {
		
		type KeyValue struct {
			key int
			value string
		}
		dList := dlist.NewDList()
		dList.PushToRightEdge(KeyValue{key:1, value: "one"}) // [1:"one"]
		dList.PushToRightEdge(KeyValue{key:2, value: "two"}) // [1:"one" 2:"two"]
		dList.PushToRightEdge(KeyValue{key:3, value: "three"}) // [1:"one" 2:"two" 3:"three"]
		dList.PushToRightEdge(KeyValue{key:4, value: "four"}) // [1:"one" 2:"two" 3:"three" 4:"four"]
		fmt.Printf("\nList before sort with Stringer() formatting:\n%s\n", dList)
		dList.SetLessItems(func(x, y *dlist.DListItem) bool {
			return x.Value().(KeyValue).value < y.Value().(KeyValue).value
		})
		sort.Sort(dList.(*dlist.DList))
		fmt.Printf("\nList after sort with Stringer() formatting:\n%s\n", dList)
	})
}
