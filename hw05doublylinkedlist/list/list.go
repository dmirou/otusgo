package list

import "errors"

// Item is the item of the doubly linked list
type Item struct {
	value interface{}
	next  *Item
	prev  *Item
}

// SetValue sets a value of the item
func (itm *Item) SetValue(value interface{}) {
	itm.value = value
}

// Value returns a value of the item
func (itm Item) Value() interface{} {
	return itm.value
}

// NewItem creates a new item with the value
func NewItem(value interface{}) *Item {
	newItem := new(Item)
	newItem.SetValue(value)
	return newItem
}

// SetNext set a next item connected to the current item in the list
func (itm *Item) SetNext(next *Item) {
	itm.next = next
}

// Next returns a next item connected to the current item in the list
func (itm Item) Next() *Item {
	return itm.next
}

// SetPrev set a previous item connected to the current item in the list
func (itm *Item) SetPrev(prev *Item) {
	itm.prev = prev
}

// Prev returns a previous item connected to the current item in the list
func (itm Item) Prev() *Item {
	return itm.prev
}

// List describes a doubly linked list
type List struct {
	first *Item
	last  *Item
	len   int
}

// Len returns a count of elements in the list
func (lst List) Len() int {
	return lst.len
}

// First returns a first item of the list
func (lst List) First() *Item {
	return lst.first
}

// Last returns a last item of the list
func (lst List) Last() *Item {
	return lst.last
}

// PushFront adds a value to the beginning of the list
func (lst *List) PushFront(v interface{}) {
	newItem := NewItem(v)
	if lst.first == nil {
		lst.first = newItem
		lst.last = newItem
		lst.len++
		return
	}
	lst.first.SetPrev(newItem)
	newItem.SetNext(lst.first)
	lst.first = newItem
	lst.len++
}

// PushBack adds a value to the end of the list
func (lst *List) PushBack(v interface{}) {
	newItem := NewItem(v)
	if lst.first == nil {
		lst.first = newItem
		lst.last = newItem
		lst.len++
		return
	}
	lst.last.SetNext(newItem)
	newItem.SetPrev(lst.last)
	lst.last = newItem
	lst.len++
}

// Remove removes an item from the list
func (lst *List) Remove(i *Item) (bool, error) {
	if i == nil {
		return false, errors.New("can't remove nil from the list")
	}
	if i.Prev() == nil && i.Next() == nil {
		lst.first = nil
		lst.last = nil
		lst.len = 0
		return true, nil
	}
	if lst.Last() == i {
		lst.last = i.Prev()
	}
	if i.Prev() != nil {
		i.Prev().SetNext(i.Next())
		lst.len--
		return true, nil
	}
	if lst.First() == i {
		lst.first = i.Next()
	}
	i.Next().SetPrev(i.Prev())
	lst.len--
	return true, nil
}
