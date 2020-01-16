package list

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestItemValue checked that the Value method correctly returns an assigned item value.
func TestItemValue(t *testing.T) {
	item := Item{}
	item.value = rand.Int()
	assert.Equalf(t, item.value, item.Value(), "Value was not received")
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
		assert.Equalf(t, next, item.Next(), "Next item was not received")
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
		assert.Equalf(t, prev, item.Prev(), "Previous item was not received")
	}
}

// TestListPushFront checks that values are added to the list via PushFront method.
func TestListPushFront(t *testing.T) {
	var list = new(List)
	var values, err = GenerateSliceWithLength(5, 10)
	assert.Nilf(t, err, "Slice was not generated correctly")
	for _, value := range values {
		list.PushFront(value)
		assert.Equalf(t, value, list.First().Value(), "Value was not added to the front of the list")
	}
	assert.Equalf(t, len(values), list.Len(), "Not all values were added to the list")
}

// TestListPushBack checks that values are added to the list via PushBack method.
func TestListPushBack(t *testing.T) {
	generateRandomList(t)
}

// generateRandomList return a random list and it's values.
func generateRandomList(t *testing.T) (*List, []int) {
	var list = new(List)
	var values, err = GenerateSliceWithLength(5, 10)
	assert.Nilf(t, err, "Slice was not generated correctly")
	for _, value := range values {
		list.PushBack(value)
		assert.Equalf(t, value, list.Last().Value(), "Value was not added to the end of the list")
	}
	assert.Equalf(t, len(values), list.Len(), "Not all values were added to the list")
	return list, values
}

// RemoveTestCase describes input data for testing list.Remove method.
type RemoveTestCase struct {
	Source        []int
	IndexToRemove int
	Result        []int
}

// TestRemove checks that a list item is removed from the list.
func TestRemove(t *testing.T) {
	testCases := []RemoveTestCase{
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
	for _, testCase := range testCases {
		var list = new(List)
		for _, value := range testCase.Source {
			list.PushBack(value)
		}
		var curItem = list.First()
		var itemToRemove Item
		for i := 0; i < list.Len(); i++ {
			if i == testCase.IndexToRemove {
				itemToRemove = *curItem
				break
			}
			curItem = curItem.Next()
		}
		list.Remove(itemToRemove)
		var expectedLength = len(testCase.Result)
		assert.Equalf(t, expectedLength, list.Len(), "Item was not removed from the list")
		if expectedLength == 0 {
			continue
		}
		var actualValues = make([]int, expectedLength)
		var i = 0
		for curItem := list.First(); curItem != nil; curItem = curItem.Next() {
			actualValues[i] = curItem.Value().(int)
			i++
		}
		assert.Equalf(t, testCase.Result, actualValues, "incorrect values when traversing a list from the beginning")

		actualValues = make([]int, expectedLength)
		i = expectedLength - 1
		for curItem := list.Last(); curItem != nil; curItem = curItem.Prev() {
			actualValues[i] = curItem.Value().(int)
			i--
		}
		assert.Equalf(t, testCase.Result, actualValues, "incorrect values when traversing a list from the end")
	}
}
