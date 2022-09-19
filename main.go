package main

import (
	"fmt"
	multi "golang-datastructure/DoublepleLinkList"
	link "golang-datastructure/datastructure"
)

func main() {
	// Linkedlist
	linkedlist := link.LinkedList{}
	node1 := &link.SNode{Data: 3}
	node2 := &link.SNode{Data: 1}
	node3 := &link.SNode{Data: 1}
	node4 := &link.SNode{Data: 5}

	linkedlist.InsertFirst(node1)
	linkedlist.InsertLast(node2)
	linkedlist.InsertFirst(node3)
	linkedlist.InsertAfter(node4, 4)
	linkedlist.DeletedWithValue(5)
	linkedlist.Print()

	// Multiple Linklist
	fmt.Println()
	dataMulti := multi.MultipleList{}
	nodeM1 := &multi.MNode{Data: 2}
	nodeM2 := &multi.MNode{Data: 3}
	nodeM3 := &multi.MNode{Data: 5}
	nodeM4 := &multi.MNode{Data: 1}

	dataMulti.Prepend(nodeM1)
	dataMulti.Prepend(nodeM2)
	dataMulti.Prepend(nodeM3)
	dataMulti.InsertBefore(nodeM4, 4)
	dataMulti.Print()

	// 5,3,2
	// element 1
	// 1 5 3 2
	// element 2
	// 5 1 3 2
	// element 3
	// 5 3 1 2

}
