package list

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// TestItemValue checked that the Value method correctly returns an assigned item value.
func TestItemValue(t *testing.T) {
	item := Item{}

	values := []interface{}{
		nil,
		10,
		"str",
	}

	for _, v := range values {
		item.value = v
		if v != item.Value() {
			t.Errorf("expected value: %v, got: %v", v, item.Value())
		}
	}
}

// TestItemNext checked that the Next method correctly returns an assigned next item.
func TestItemNext(t *testing.T) {
	item := Item{}

	nexts := []*Item{
		nil,
		{},
		{value: 2},
	}

	for _, next := range nexts {
		item.next = next
		if next != item.Next() {
			t.Errorf("expected next: %v, got: %v", next, item.Next())
		}
	}
}

// TestItemPrev checked that the Prev method correctly returns an assigned previous item.
func TestItemPrev(t *testing.T) {
	item := Item{}

	prevs := []*Item{
		nil,
		{},
		{value: 2},
	}

	for _, prev := range prevs {
		item.prev = prev
		if prev != item.Prev() {
			t.Errorf("expected prev: %v, got: %v", prev, item.Prev())
		}
	}
}

// TestListPushFront checks that values are added to the list via PushFront method.
func TestListPushFront(t *testing.T) {
	list := NewList()
	values := []int{3, 4, 1, 2, 8}

	for _, value := range values {
		list.PushFront(value)

		if list.Front().Value() != value {
			t.Errorf("expected front value: %v, got: %v", value, list.Front().Value())
		}
	}

	if list.Len() != len(values) {
		t.Errorf("expected list len: %v, got: %v", len(values), list.Len())
	}
}

// TestListPushBack checks that values are added to the list via PushBack method.
func TestListPushBack(t *testing.T) {
	list := NewList()
	values := []int{3, 4, 1, 2, 8}

	for _, value := range values {
		list.PushBack(value)

		if list.Back().Value() != value {
			t.Errorf("expected back value: %v, got: %v", value, list.Back().Value())
		}
	}

	if list.Len() != len(values) {
		t.Errorf("expected list len: %v, got: %v", len(values), list.Len())
	}
}

// RemoveTestData describes input data for testing list.Remove method.
type RemoveTestData struct {
	Source        []int
	IndexToRemove int
	Result        []int
}

// TestRemove checks that a list item is removed from the list.
// nolint: funlen
func TestRemove(t *testing.T) {
	tds := []RemoveTestData{
		{
			Source:        []int{4},
			IndexToRemove: 0,
			Result:        []int{},
		},
		{
			Source:        []int{4, 2, 8, 4, 1},
			IndexToRemove: 0,
			Result:        []int{2, 8, 4, 1},
		},
		{
			Source:        []int{4, 1, 2, 10, 12, 4},
			IndexToRemove: 5,
			Result:        []int{4, 1, 2, 10, 12},
		},
		{
			Source:        []int{4, 2, 8, 1},
			IndexToRemove: 1,
			Result:        []int{4, 8, 1},
		},
		{
			Source:        []int{4, 1, 2, 10},
			IndexToRemove: 2,
			Result:        []int{4, 1, 10},
		},
	}
	for _, td := range tds {
		list := NewList()

		for _, value := range td.Source {
			list.PushBack(value)
		}

		var toRemove *Item

		var current = list.Front()

		for i := 0; i < list.Len(); i++ {
			if i == td.IndexToRemove {
				toRemove = current
				break
			}

			current = current.Next()
		}

		list.Remove(toRemove)

		var length = len(td.Result)
		if list.Len() != length {
			t.Errorf("expected length: %v, got: %v", length, list.Len())
		}

		if length == 0 {
			continue
		}

		var (
			values = make([]int, length)
			i      = 0
		)

		for cur := list.Front(); cur != nil; cur = cur.Next() {
			values[i] = cur.Value().(int)
			i++
		}

		if !cmp.Equal(values, td.Result) {
			t.Errorf("expected values: %v, got: %v", td.Result, values)
		}

		values = make([]int, length)
		i = length - 1

		for curItem := list.Back(); curItem != nil; curItem = curItem.Prev() {
			values[i] = curItem.Value().(int)
			i--
		}

		if !cmp.Equal(values, td.Result) {
			t.Errorf("expected values: %v, got: %v", td.Result, values)
		}
	}
}

// TestRemoveFromAnotherList checks that the list can't remove an item from a different list.
func TestRemoveFromAnotherList(t *testing.T) {
	first := NewList()
	second := NewList()
	values := []int{3, 4, 1, 2, 8}

	for _, value := range values {
		first.PushBack(value)
		second.PushBack(value)
	}

	first.Remove(second.Front())

	if 5 != first.Len() {
		t.Errorf("expected length: %d, got: %d", 5, first.Len())
	}
}

// TestMoveToFront checks that a list item is correctly moved to the front.
func TestMoveToFront(t *testing.T) {
	list := NewList()
	values := []int{3, 4, 1, 2, 8}

	for _, value := range values {
		list.PushBack(value)
	}

	first := list.Front()
	list.MoveToFront(first)

	if first != list.Front() {
		t.Errorf("expected first item: %v, got: %v", first, list.Front())
	}

	if len(values) != list.Len() {
		t.Errorf("expected length: %v, got: %v", len(values), list.Len())
	}

	eight := list.Back()
	list.MoveToFront(eight)

	expected := []int{8, 3, 4, 1, 2}
	cur := list.Front()

	for _, v := range expected {
		if v != cur.Value() {
			t.Errorf("expected value: %v, got: %v", v, cur.Value())
		}

		cur = cur.Next()
	}

	if len(expected) != list.Len() {
		t.Errorf("expected length: %v, got: %v", len(expected), list.Len())
	}

	four := list.Front().Next().Next()
	list.MoveToFront(four)

	expected = []int{4, 8, 3, 1, 2}
	cur = list.Front()

	for _, v := range expected {
		if v != cur.Value() {
			t.Errorf("expected value: %v, got: %v", v, cur.Value())
		}

		cur = cur.Next()
	}

	if len(expected) != list.Len() {
		t.Errorf("expected length: %v, got: %v", len(expected), list.Len())
	}
}
