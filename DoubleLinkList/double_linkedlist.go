package doublelinklist

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

func (l *MultipleList) InsertFirst(Node *MNode) {
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
		l.InsertFirst(Node)
	}
	if element > l.Length || element == 0 {
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

// Delete Before

func (l *MultipleList) DeleteBefore(element int) {
	if element == 0 || element == 1 {
		return
	}
	if l.Length == 2 && element == 2 {
		l.Head = l.Head.Next
		l.Head.Prev = nil
		l.Head.Next = nil
		l.Length--
		return
	}
	if l.Length > 2 {
		data := l.Head
		index := 1
		length := l.Length
		for index <= length {
			if element == index {
				newLastNode := data.Prev.Prev
				newLastNode.Next = data
				data.Prev = newLastNode
				l.Length--
			}
			data = data.Next
			index++
		}
	}
}

func (l *MultipleList) InsertLast(Node *MNode) {
	if l.Length == 0 {
		return
	}
	if l.Length == 1 {
		Node.Prev = l.Head
		l.Head.Next = Node
		l.Length++
	}

	if l.Length > 1 {
		index := 1
		data := l.Head
		for index <= l.Length {
			if data.Next == nil {
				Node.Prev = data
				data.Next = Node
				l.Length++
				break
			}

			data = data.Next
			index++
		}
	}
}

func (l *MultipleList) DeleteLast() {
	if l.Length == 1 {
		l.Head = nil
		l.Head.Next = nil
		l.Head.Prev = nil
		l.Length--
	}

	data := l.Head
	index := 1
	for index <= l.Length {
		if data.Next == nil {
			lastNode := data.Prev
			lastNode.Next = nil
			l.Length--
		}
		data = data.Next
		index++
	}
}

// Insert After

// Delete After

// Insert Last

// Delete Last
