package main

import (
	"fmt"
	link "golang-datastructure/datastructure"
	multi "golang-datastructure/doublelinklist"
)

func main() {
	fmt.Println("---Singlelinkedlist---")
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
	fmt.Println("----Multilinkedlist----")
	dataMulti := multi.MultipleList{}
	// nodeM5 := &multi.MNode{Data: 5}
	nodeM5 := &multi.MNode{Data: 0}
	nodeM4 := &multi.MNode{Data: 4}
	nodeM3 := &multi.MNode{Data: 3}
	nodeM2 := &multi.MNode{Data: 2}
	nodeM1 := &multi.MNode{Data: 1}

	dataMulti.InsertFirst(nodeM1)
	dataMulti.InsertFirst(nodeM2)
	dataMulti.InsertFirst(nodeM3)
	dataMulti.InsertFirst(nodeM4)

	// dataMulti.InsertBefore(nodeM5, 2)
	// dataMulti.DeleteBefore(3)
	dataMulti.InsertLast(nodeM5)
	dataMulti.DeleteLast()
	dataMulti.Print()

	// 5,3,2
	// element 1
	// 1 5 3 2
	// element 2
	// 5 1 3 2
	// element 3
	// 5 3 1 2

	// Delete Before
	// 8 5

}
