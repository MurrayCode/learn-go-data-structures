package queue

import (
	"reflect"
	"testing"
)

func TestEnqueue(t *testing.T) {
	t.Run("correct items in queue", func(t *testing.T) {
		got := Queue{}
		expected := Queue{[]string{"a", "b", "c"}}
		got.Enqueue("a")
		got.Enqueue("b")
		got.Enqueue("c")
		if !reflect.DeepEqual(got.items, expected.items) {
			t.Errorf("expected %s, got %s", expected.items, got.items)
		}
	})
}

func TestDequeue(t *testing.T) {
	t.Run("removes correct item", func(t *testing.T) {
		got := Queue{[]string{"a", "b", "c"}}
		got.Dequeue()
		expected := Queue{[]string{"b", "c"}}
		if !reflect.DeepEqual(got.items, expected.items) {
			t.Errorf("expected %s, got %s", expected.items, got.items)
		}
	})

}
