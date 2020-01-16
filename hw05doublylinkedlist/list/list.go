package list

import "github.com/pkg/errors"

// Item is the item of the doubly linked list.
type Item struct {
	value  interface{}
	listId int
	next   *Item
	prev   *Item
}

// Value returns a value of the item.
func (i Item) Value() interface{} {
	return i.value
}

// newItem creates a new item with the value.
func newItem(value interface{}) *Item {
	newItem := new(Item)
	newItem.value = value
	return newItem
}

// Next returns a next item connected to the current item in the list.
func (i Item) Next() *Item {
	return i.next
}

// Prev returns a previous item connected to the current item in the list.
func (i Item) Prev() *Item {
	return i.prev
}

// List describes a doubly linked list.
type List struct {
	id    int
	first *Item
	last  *Item
	len   int
}

// Len returns a count of elements in the list.
func (l List) Len() int {
	return l.len
}

// First returns a first item of the list.
func (l List) First() *Item {
	return l.first
}

// Last returns a last item of the list.
func (l List) Last() *Item {
	return l.last
}

// PushFront adds a value to the beginning of the list.
func (l *List) PushFront(value interface{}) {
	newItem := newItem(value)
	newItem.listId = l.id
	if l.first == nil {
		l.first = newItem
		l.last = newItem
		l.len++
		return
	}
	l.first.prev = newItem
	newItem.next = l.first
	l.first = newItem
	l.len++
}

// PushBack adds a value to the end of the list.
func (l *List) PushBack(value interface{}) {
	newItem := newItem(value)
	newItem.listId = l.id
	if l.first == nil {
		l.first = newItem
		l.last = newItem
		l.len++
		return
	}
	l.last.next = newItem
	newItem.prev = l.last
	l.last = newItem
	l.len++
}

// Remove removes an item from the list.
func (l *List) Remove(item Item) (bool, error) {
	if item.listId != l.id {
		return false, errors.Errorf("The list doesn't contain the passed item")
	}
	if item.Prev() == nil && item.Next() == nil {
		l.first = nil
		l.last = nil
		l.len = 0
		return true, nil
	}
	if *l.First() == item {
		l.first = item.Next()
	}
	if *l.Last() == item {
		l.last = item.Prev()
	}
	if item.Prev() != nil {
		item.Prev().next = item.Next()
	}
	if item.Next() != nil {
		item.Next().prev = item.Prev()
	}
	l.len--
	return true, nil
}

// NewList creates a new empty list with a random id
func NewList() (*List, error) {
	var list = new(List)
	id, err := GenerateInt(0, 10000)
	if err != nil {
		return nil, errors.Wrap(err, "can't generate new list id")
	}
	list.id = id
	return list, nil
}
