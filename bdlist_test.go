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
