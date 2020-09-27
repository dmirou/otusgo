package list

import (
	"fmt"
)

// Item is the item of the doubly linked list.
type Item struct {
	value interface{}
	list  *List
	next  *Item
	prev  *Item
}

// newItem creates a new item with the value.
func newItem(value interface{}) *Item {
	newItem := new(Item)
	newItem.value = value

	return newItem
}

// Value returns a value of the item.
func (i Item) Value() interface{} {
	return i.value
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
	front *Item
	back  *Item
	len   int
}

// NewList creates a new empty list.
func NewList() *List {
	return new(List)
}

// Len returns a count of elements in the list.
func (l List) Len() int {
	return l.len
}

// Front returns a front item of the list.
func (l List) Front() *Item {
	return l.front
}

// Back returns a back item of the list.
func (l List) Back() *Item {
	return l.back
}

// PushFront adds a value to the beginning of the list.
func (l *List) PushFront(value interface{}) *Item {
	item := newItem(value)
	item.list = l

	if l.front == nil {
		l.front = item
		l.back = item
		l.len++

		return item
	}

	l.front.prev = item
	item.next = l.front
	l.front = item
	l.len++

	return item
}

// PushBack adds a value to the end of the list.
func (l *List) PushBack(value interface{}) *Item {
	item := newItem(value)
	item.list = l

	if l.front == nil {
		l.front = item
		l.back = item
		l.len++

		return item
	}

	l.back.next = item
	item.prev = l.back
	l.back = item
	l.len++

	return item
}

// Remove removes an item from the list.
func (l *List) Remove(item *Item) error {
	if item.list != l {
		return fmt.Errorf("the list doesn't contain the passed item")
	}

	if item.Prev() == nil && item.Next() == nil {
		item.list = nil
		l.front = nil
		l.back = nil
		l.len = 0

		return nil
	}

	item.list = nil

	if l.Front() == item {
		l.front = item.Next()
	}

	if l.Back() == item {
		l.back = item.Prev()
	}

	if item.Prev() != nil {
		item.Prev().next = item.Next()
	}

	if item.Next() != nil {
		item.Next().prev = item.Prev()
	}

	l.len--

	return nil
}

// Move a list item to the beginning of the list.
func (l *List) MoveToFront(item *Item) error {
	if item.list != l {
		return fmt.Errorf("the list doesn't contain the passed item")
	}

	if item == l.front {
		return nil
	}

	item.prev.next = item.next

	if item.next != nil {
		item.next.prev = item.prev
	}

	if item == l.back && item.next != nil {
		l.back = item.next
	} else if item == l.back {
		l.back = item.prev
	}

	item.prev = nil
	item.next = l.front
	l.front = item

	return nil
}
