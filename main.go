package main

import (
	"fmt"

	"github.com/MurrayCode/learn-go-data-structures/tree"
)

func main() {
	t := &tree.Node{Data: 99}
	t.Insert(100)
	t.Insert(10)
	t.Insert(20)
	fmt.Println(t)
	if t.Search(19) {
		fmt.Println(tree.Count)
	} else {
		fmt.Println("Key Doesn't exist in Tree")
	}
}
