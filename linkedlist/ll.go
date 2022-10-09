package linkedlist

import "fmt"

type Node struct {
	Data string
	next *Node
}

type LinkedList struct {
	Head   *Node
	length int
}

func (l *LinkedList) PrintList(n *Node) {
	fmt.Println(n.Data)
	if n.next != nil {
		l.PrintList(n.next)
	} else {
		return
	}
}

func (l *LinkedList) Prepend(n *Node) {
	second := l.Head
	l.Head = n
	l.Head.next = second
	l.length++
}
