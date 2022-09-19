package main

import (
	. "golang-datastructure/datastructure"
)

func main() {
	data := LinkedList{}
	node1 := &Node{Data: 3}
	node2 := &Node{Data: 1}
	node3 := &Node{Data: 5}
	node4 := &Node{Data: 2}

	data.InsertFirst(node1)
	data.InsertLast(node2)
	data.InsertFirst(node3)
	data.InsertAfter(node4, 4)
	data.DeletedWithValue(5)
	data.Print()
}
