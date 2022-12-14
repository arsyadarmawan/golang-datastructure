package main

import (
	"fmt"
	link "golang-datastructure/datastructure"
	multi "golang-datastructure/doublelinklist"
	graph "golang-datastructure/graph"
	_queue "golang-datastructure/queue"
	tree "golang-datastructure/tree"

	// queue "golang-datastructure/queue"
	stack "golang-datastructure/stack"
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
	// dataMulti.DeleteLast()
	dataMulti.Print()

	fmt.Println("Stack Theory")
	stack := stack.Stack{}
	stack.Push(1)
	stack.Push(2)
	stack.Pop()
	fmt.Println(stack)

	fmt.Println("--- Queue Theory ---")
	queues := _queue.Queue{}
	queues.Enqueu(1)
	queues.Enqueu(2)
	queues.Dequeue()
	fmt.Println(queues)

	fmt.Println("--- Trees ---")
	trees := tree.Tree{Data: 100}
	trees.Insert(10)
	trees.Insert(200)
	fmt.Println(trees.Search(200))
	fmt.Println(trees)

	fmt.Println("--- Graph ---")
	graphs := &graph.Graph{}
	for i := 1; i <= 5; i++ {
		graphs.AddVertex(i)
	}
	graphs.AddEdgge(1, 2)
	graphs.AddEdgge(1, 2)
	graphs.AddEdgge(2, 1)
	graphs.AddEdgge(6, 1)

	graphs.Print()

}
