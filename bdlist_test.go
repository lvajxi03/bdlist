package bdlist_test

import (
	"github.com/lvajxi03/bdlist"
	"testing"
)

func TestCreate(t *testing.T) {
	list := bdlist.New()
	if list.Length != 0 {
		t.Fatal("Non-empty list found!")
	}
}

func TestAppend(t *testing.T) {
	list := bdlist.New()
	list.AppendVal(1)
	list.AppendVal(3)
	list.AppendVal(5)
	val := list.GetValueAt(1)
	if val != 3 {
		t.Fatalf("Expected `3`, found `%d`", val)
	}
}

func TestPrepend(t *testing.T) {
	list := bdlist.New()
	list.PrependVal(2)
	list.PrependVal(4)
	list.PrependVal(6)
	val := list.GetValueAt(2)
	if val != 2 {
		t.Fatalf("Expected `2`, found `%d`", val)
	}
}

func TestRemove(t *testing.T) {
	list := bdlist.New()
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 21, 22, 2021, 2022}
	for _, value := range arr {
		list.PrependVal(value)
	}
	if list.Length != 16 {
		t.Fatalf("Expected 16 elements in the list, found %d!", list.Length)
	}
	for iter := list.Head; iter != nil; iter = iter.Next {
		if iter.Value.(int) %2 == 0 {
			list.Remove(iter)
		}
	}
	if list.Length != 8 {
		t.Fatalf("Expected 8 elements in the list, found %d!", list.Length)
	}

	value := list.GetValueAt(2)
	if value != 11 {
		t.Fatalf("Expected value: 11, found %d!", value)
	}
}
