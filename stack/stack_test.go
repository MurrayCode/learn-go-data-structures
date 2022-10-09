package stack

import (
	"reflect"
	"testing"
)

func TestPush(t *testing.T) {
	t.Run("correct items in stack", func(t *testing.T) {
		got := Stack{}
		expected := Stack{[]int{1, 2, 3}}
		got.Push(1)
		got.Push(2)
		got.Push(3)
		if !reflect.DeepEqual(expected.items, got.items) {
			t.Errorf("expected %d, got %d", expected.items, got.items)
		}
	})
}

func TestPop(t *testing.T) {
	t.Run("correct items in stack", func(t *testing.T) {
		got := Stack{[]int{1, 2, 3}}
		expected := Stack{[]int{1, 2}}
		got.Pop()
		if !reflect.DeepEqual(got.items, expected.items) {
			t.Errorf("expected %d, got %d", expected.items, got.items)
		}
	})
	t.Run("returns correct value", func(t *testing.T) {
		stack := Stack{[]int{1, 2, 3}}
		got := stack.Pop()
		expected := 3
		if got != expected {
			t.Errorf("expected %d, got %d", expected, got)
		}
	})
}
