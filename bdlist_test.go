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

func TestAppend1(t *testing.T) {
	list := bdlist.New()
	list.AppendVal(1)
	list.AppendVal(3)
	list.AppendVal(5)
	val, _ := list.GetValueAt(1)
	if val != 3 {
		t.Fatalf("Expected `3`, found `%d`", val)
	}
}

func TestAppend2(t *testing.T) {
	list := bdlist.New()
	err := list.Append(nil)
	if err == nil {
		t.Fatal("Expected non-nil error")
	}
}

func TestAppend3(t *testing.T) {
	var list *bdlist.BDList = nil
	err := list.AppendVal(5)
	if err == nil {
		t.Fatal("Expected non-nil error")
	}
}

func TestPrepend1(t *testing.T) {
	list := bdlist.New()
	list.PrependVal(2)
	list.PrependVal(4)
	list.PrependVal(6)
	val, _ := list.GetValueAt(2)
	if val != 2 {
		t.Fatalf("Expected `2`, found `%d`", val)
	}
}

func TestPrepend2(t *testing.T) {
	list := bdlist.New()
	err := list.Prepend(nil)
	if err == nil {
		t.Fatal("Expected non-nil error")
	}
}

func TestRemove1(t *testing.T) {
	list := bdlist.New()
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 21, 22, 2021, 2022}
	for _, value := range arr {
		list.PrependVal(value)
	}
	if list.Length != 16 {
		t.Fatalf("Expected 16 elements in the list, found %d!", list.Length)
	}
	for iter := list.Head; iter != nil; iter = iter.Next {
		if iter.Value.(int)%2 == 0 {
			list.Remove(iter)
		}
	}
	if list.Length != 8 {
		t.Fatalf("Expected 8 elements in the list, found %d!", list.Length)
	}

	value, _ := list.GetValueAt(2)
	if value != 11 {
		t.Fatalf("Expected value: 11, found %d!", value)
	}
}

func TestRemove2(t *testing.T) {
	list := bdlist.New()
	_, err := list.Remove(nil)
	if err == nil {
		t.Fatal("Expected non-nil error")
	}
}

func TestRemove3(t *testing.T) {
	var list *bdlist.BDList = nil
	_, err := list.RemoveAt(3)
	if err == nil {
		t.Fatal("Expected non-nil error")
	}
}

func TestGetValueAt1(t *testing.T) {
	list := bdlist.New()
	val, _ := list.GetValueAt(0)
	if val != nil {
		t.Fatal("Expected nil value")
	}
}
