package tree

type Node struct {
	Left  *Node
	Right *Node
	Data  int
}

var Count int

func (n *Node) Insert(data int) {
	if n.Data < data {
		if n.Right == nil {
			n.Right = &Node{Data: data}
		} else {
			n.Right.Insert(data)
		}
	} else if n.Data > data {
		if n.Left == nil {
			n.Left = &Node{Data: data}
		} else {
			n.Left.Insert(data)
		}
	}
}

func (n *Node) Search(data int) bool {
	Count++
	if n == nil {
		return false
	}
	if n.Data < data {
		return n.Right.Search(data)
	} else if n.Data > data {
		return n.Left.Search(data)
	}
	return true
}
