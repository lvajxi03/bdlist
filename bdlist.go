// Package bdlist implements simple bi-directional list
// and its typical operations
package bdlist

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

// NewBDList creates new, empty BDList
func New() *BDList {
	list := &BDList{Head: nil, Tail: nil, Length: 0}
	return list
}

// Append adds an element after Tail
func (list *BDList) Append(element *Element) {
	if element != nil {
		if list.Length == 0 {
			list.Head = element
			list.Head.Prev = nil
			list.Tail = element
			list.Head.Next = list.Tail
			list.Tail.Prev = list.Head
			list.Tail.Next = nil
			list.Length++
		} else {
			list.Tail.Next = element
			element.Prev = list.Tail
			element.Next = nil
			list.Tail = list.Tail.Next
			list.Length++
		}
	}
}

// Prepend adds an element before Head
func (list *BDList) Prepend(element *Element) {
	if element != nil {
		if list.Length == 0 {
			list.Head = element
			list.Head.Prev = nil
			list.Tail = element
			list.Head.Next = list.Tail
			list.Tail.Prev = list.Head
			list.Tail.Next = nil
			list.Length++
		} else {
			list.Head.Prev = element
			element.Next = list.Head
			element.Prev = nil
			list.Head = list.Head.Prev
			list.Length++
		}
	}
}

// AppendVal creates an Element with given value
// and adds it after Tail
func (list *BDList) AppendVal(value any) {
	element := &Element{Value: value}
	list.Append(element)
}

// PrependVal creates an Element with given value
// and addis it before Head
func (list *BDList) PrependVal(value any) {
	element := &Element{Value: value}
	list.Prepend(element)
}

// InsertAt adds an element at the given position
// (zero-based index)
//
// If pos is negative, InsertAt prepends the element
// If pos is higher than list lenthg, InsertAt appends the element
func (list *BDList) InsertAt(pos int, element *Element) {
	if element != nil {
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
	}
}

// InsertValAt creates an Element with given value
// and adds it at given position (zero-based index)
//
//  If pos is negative, InsertAt prepends the element
// If pos is higher than list lenthg, InsertAt appends the element
func (list *BDList) InsertValAt(pos int, value any) {
	element := NewElement(value)
	list.InsertAt(pos, element)
}

// Remove removes an Element and returns it.
// If the element is not found, `nil` is returned
func (list *BDList) Remove(element *Element) any {
	if element != nil {
		if element == list.Head {
			if list.Head.Next != nil {
				list.Head = list.Head.Next
				list.Head.Prev = nil
				list.Length--
			} else {
				list.Head = nil
				list.Length--
			}
		} else if element == list.Tail {
			list.Tail = list.Tail.Prev
			list.Tail.Next = nil
			list.Length--
		} else {
			element.Prev.Next = element.Next
			element.Next.Prev = element.Prev
			list.Length--
		}
		return element.Value
	}
	return nil
}

// GetAt returns the element at given position
// or `nil`, if not found
func (list *BDList) GetAt(pos int) *Element {
	var result *Element = nil
	cur := 0
	if list != nil {
		for iter := list.Head; iter != nil; iter = iter.Next {
			if cur == pos {
				result = iter
			}
			cur++
		}
	}
	return result
}

// GetValueAt returns value of the element at given
// position or `nil`, if not found
func (list *BDList) GetValueAt(pos int) any {
	result := list.GetAt(pos)
	if result != nil {
		return result.Value
	}
	return nil
}

// RemoveAt removes element at given position, if found,
// and returns it.
// `nil` is returned when element not found.
func (list *BDList) RemoveAt(pos int) any {
	cur := 0
	if list != nil {
		for iter := list.Head; iter != nil; iter = iter.Next {
			if cur == pos {
				return list.Remove(iter)
			}
			cur++
		}
	}
	return nil
}
