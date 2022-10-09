package linkedlist

import "fmt"

type Node struct {
	Data string
	Next *Node
}

type LinkedList struct {
	Head   *Node
	length int
}

func (l *LinkedList) PrintList(n *Node) {
	fmt.Println(n.Data)
	if n.Next != nil {
		l.PrintList(n.Next)
	} else {
		return
	}
}

func (l *LinkedList) Prepend(n *Node) {
	old := l.Head
	l.Head = n
	l.Head.Next = old
	l.length++
}
