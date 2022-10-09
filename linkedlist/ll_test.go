package linkedlist

import (
	"testing"
)

func TestPrepend(t *testing.T) {
	t.Run("head of list is correct", func(t *testing.T) {
		list := LinkedList{}
		n1 := &Node{Data: "test"}
		list.Prepend(n1)
		expected := "test"
		got := list.Head.Data
		if got != expected {
			t.Errorf("expected %s, got %s", expected, got)
		}
	})
	t.Run("next node in ll is correct", func(t *testing.T) {
		list := LinkedList{}
		n1 := &Node{Data: "test"}
		n2 := &Node{Data: "test2"}
		list.Prepend(n1)
		list.Prepend(n2)
		expected := "test2"
		got := list.Head.Data
		if got != expected {
			t.Errorf("expected %s, got %s", expected, got)
		}
	})
}
