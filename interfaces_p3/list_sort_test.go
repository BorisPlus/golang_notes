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
