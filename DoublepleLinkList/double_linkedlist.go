package doubleplelinklist

import "fmt"

type MNode struct {
	Data int
	Next *MNode
	Prev *MNode
}

type MultipleList struct {
	Head   *MNode
	Length int
}

func (l *MultipleList) Prepend(Node *MNode) {
	if l.Length == 0 {
		l.Head = Node
		l.Length++
		l.Head.Next = nil
		l.Head.Prev = nil
		return
	}
	second := l.Head
	l.Head = Node
	l.Head.Next = second
	second.Prev = l.Head
	l.Length++
}

func (l *MultipleList) Print() {
	data := l.Head
	for l.Length != 0 {
		fmt.Println(data.Data)
		data = data.Next
		l.Length--
	}
}

func (l *MultipleList) InsertBefore(Node *MNode, element int) {
	length := l.Length
	if length == 0 {
		l.Prepend(Node)
	}
	if element > l.Length {
		return
	}
	if element == 1 {
		temp := l.Head
		Node.Next = temp
		temp.Prev = Node
		l.Head = Node
		l.Length++
		return
	}
	data := l.Head
	for i := 0; i <= length; i++ {
		if i == element-1 {
			fmt.Println("nilai data sekarang ", data.Data)
			nodeBefore := data.Prev
			Node.Next = data
			data.Prev = Node
			data.Prev.Prev = nodeBefore
			nodeBefore.Next = Node
			Node.Prev = nodeBefore
			l.Length++
			return
		}
		data = data.Next
	}
}
