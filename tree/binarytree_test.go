package tree

import (
	"testing"
)

func TestInsert(t *testing.T) {
	tree := &Node{Data: 100}
	tree.Insert(90)
	tree.Insert(101)
	t.Run("right node is more than root node", func(t *testing.T) {
		got := tree.Right.Data > 100
		expected := true
		if got != expected {
			t.Errorf("expected %v, got %v", expected, got)
		}
	})
	t.Run("left node is less than root node", func(t *testing.T) {
		got := tree.Left.Data < 100
		expected := true
		if got != expected {
			t.Errorf("expected %v, got %v", expected, got)
		}
	})
	t.Run("right node is correct value", func(t *testing.T) {
		got := tree.Right.Data
		expected := 101
		if got != expected {
			t.Errorf("expected %d, got %d", expected, got)
		}
	})
	t.Run("left node is correct value", func(t *testing.T) {
		got := tree.Left.Data
		expected := 90
		if got != expected {
			t.Errorf("expected %d, got %d", expected, got)
		}
	})
}
func TestSearch(t *testing.T) {
	tree := &Node{Data: 100}
	tree.Insert(20)
	tree.Insert(31)
	t.Run("search returns true", func(t *testing.T) {
		got := tree.Search(20)
		expected := true
		if got != expected {
			t.Errorf("expected %v, got %v", expected, got)
		}
	})
	t.Run("search returns false", func(t *testing.T) {
		got := tree.Search(1)
		expected := false
		if got != expected {
			t.Errorf("expected %v, got %v", expected, got)
		}
	})
}
