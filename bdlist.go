// Package bdlist implements simple bi-directional list
// and its typical operations
package bdlist

import (
	"errors"
)

// An Element contains value and pointers to the other Elements,
// if any
type Element struct {
	Prev, Next *Element
	Value      any
}

// A BDList contains a collection of elements, linked together
// in bi-directional way
type BDList struct {
	Head   *Element // head of the list
	Tail   *Element // tail of the list
	Length int      // length of the list
}

// NewElement creates new Element with given value
func NewElement(value any) *Element {
	element := &Element{Value: value, Prev: nil, Next: nil}
	return element
}

// New creates new, empty BDList
func New() *BDList {
	list := &BDList{Head: nil, Tail: nil, Length: 0}
	return list
}

// Append adds an element after Tail.
// Returns non-nil error if element or list are nil.
func (list *BDList) Append(element *Element) error {
	if list == nil {
		return errors.New("nil list")
	}
	if element == nil {
		return errors.New("nil element")
	}
	if list.Length == 0 {
		list.Head = element
		list.Tail = element
	} else {
		list.Tail.Next = element
		element.Prev = list.Tail
		list.Tail = list.Tail.Next
	}
	list.Length++
	return nil
}

// Prepend adds an element before Head.
// Returns non-nil error if element or list are nil.
func (list *BDList) Prepend(element *Element) error {
	if list == nil {
		return errors.New("nil list")
	}
	if element == nil {
		return errors.New("nil element")
	}
	if list.Length == 0 {
		list.Head = element
		list.Tail = element
	} else {
		list.Head.Prev = element
		element.Next = list.Head
		list.Head = list.Head.Prev
	}
	list.Length++
	return nil

}

// AppendVal creates an Element with given value
// and adds it after Tail.
// Returns non-nil error if element or list are nil.
func (list *BDList) AppendVal(value any) error {
	if list == nil {
		return errors.New("nil list")
	}
	element := &Element{Value: value}
	return list.Append(element)
}

// PrependVal creates an Element with given value
// and adds it before Head.
// Returns error if element or list are nil.
func (list *BDList) PrependVal(value any) error {
	if list == nil {
		return errors.New("nil list")
	}
	element := &Element{Value: value}
	return list.Prepend(element)
}

// InsertAt adds an element at the given position
// (zero-based index)
//
// If pos is negative, InsertAt prepends the element
// If pos is higher than list lenthg, InsertAt appends the element.
// Returns non-nil error if element or list are nil.
func (list *BDList) InsertAt(pos int, element *Element) error {
	if list == nil {
		return errors.New("nil list")
	}
	if element == nil {
		return errors.New("nil element")
	}
	if pos >= list.Length {
		list.Append(element)
	} else if pos < 0 {
		list.Prepend(element)
	} else {
		cur := 0
		for iter := list.Head; iter != nil; iter = iter.Next {
			if cur == pos {
				element.Next = iter
				element.Prev = iter.Prev
				iter.Prev = element
				element.Prev.Next = element
			}
			cur++
		}
	}
	return nil
}

// InsertValAt creates an Element with given value
// and adds it at given position (zero-based index)
//
// If pos is negative, InsertAt prepends the element.
// If pos is higher than list lenthg, InsertAt appends the element.
// Returns non-nil error if list is nil.
func (list *BDList) InsertValAt(pos int, value any) error {
	if list == nil {
		return errors.New("nil list")
	}
	element := NewElement(value)
	return list.InsertAt(pos, element)
}

// Remove removes an Element and returns its value.
// If the element is not found, nil is returned.
// Also, returns non-nil error if list or element are nil.
func (list *BDList) Remove(element *Element) (any, error) {
	if list == nil {
		return nil, errors.New("nil list")
	}
	if element == nil {
		return nil, errors.New("nil element")
	}
	if list.Head == nil {
		return nil, errors.New("nil list")
	}
	if element == list.Head {
		list.Head = list.Head.Next
	} else if element == list.Tail {
		list.Tail = list.Tail.Prev
		list.Tail.Next = nil
	} else {
		element.Prev.Next = element.Next
		element.Next.Prev = element.Prev
	}
	list.Length--
	return element.Value, nil
}

// GetAt returns the element at given position
// or nil, if not found.
// Also, returns non-nil error if list is nil.
func (list *BDList) GetAt(pos int) (*Element, error) {
	if list == nil {
		return nil, errors.New("nil list")
	}
	var result *Element = nil
	cur := 0
	for iter := list.Head; iter != nil; iter = iter.Next {
		if cur == pos {
			result = iter
		}
		cur++
	}
	return result, nil
}

// GetValueAt returns value of the element at given
// position or nil, if not found.
// Also, returns non-nil error if list is nil.
func (list *BDList) GetValueAt(pos int) (any, error) {
	if list == nil {
		return nil, errors.New("nil list")
	}
	result, err := list.GetAt(pos)
	if result != nil {
		return result.Value, err
	}
	return nil, err
}

// RemoveAt removes element at given position, if found,
// and returns it.
// Returns nil when element not found.
// Also, returns non-nil error if list is nil.
func (list *BDList) RemoveAt(pos int) (any, error) {
	if list == nil {
		return nil, errors.New("nil list")
	}
	cur := 0
	for iter := list.Head; iter != nil; iter = iter.Next {
		if cur == pos {
			return list.Remove(iter)
		}
		cur++
	}
	return nil, nil
}
