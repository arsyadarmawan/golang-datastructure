package datastructure

import "fmt"

type SNode struct {
	Data int
	Next *SNode
}

type LinkedList struct {
	Head   *SNode
	Length int
}

func (l *LinkedList) InsertFirst(n *SNode) {
	second := l.Head
	l.Head = n
	l.Head.Next = second
	l.Length++
}

func (l *LinkedList) InsertLast(N *SNode) int {
	length := l.Length
	data := l.Head
	if length == 0 {
		l.InsertFirst(N)
		return l.Length
	}
	for i := 0; i <= length; i++ {
		if i == length-1 {
			data.Next = N
			l.Length++
		} else {
			data = data.Next
		}
	}
	return l.Length
}

func (l *LinkedList) Print() {
	data := l.Head
	for l.Length != 0 {
		fmt.Println(data.Data)
		data = data.Next
		l.Length--
	}
}

func (l *LinkedList) DeleteFirst() int {
	data := l.Head.Next
	if l.Length > 1 {
		l.Head = data
		l.Length--
	}
	return l.Length
}

func (l *LinkedList) DeleteLast() int {
	lengthTemp := l.Length
	data := l.Head
	for i := 1; i <= lengthTemp; i++ {
		if i == lengthTemp-1 {
			data.Next = nil
			l.Length--
			break
		}
		data = data.Next
	}
	return l.Length
}

func (l *LinkedList) InsertAfter(N *SNode, element int) int {
	if element == 0 {
		l.InsertFirst(N)
		return l.Length
	} else if element-1 == l.Length {
		l.InsertLast(N)
		return l.Length
	}
	length := l.Length
	data := l.Head
	for i := 0; i <= length; i++ {
		if i == element-1 {
			N.Next = data.Next
			data.Next = N
			l.Length++
		}
		data = data.Next
	}
	return l.Length
}

func (l *LinkedList) DeleteAfter(element int) {
	data := l.Head
	length := l.Length
	for i := 0; i <= length; i++ {
		if i == element-1 {
			if data.Next.Next != nil {
				temp := data.Next.Next
				data.Next = temp
				l.Length--
				return

			} else if data.Next.Next == nil {
				data.Next.Next = nil
				l.Length--
				return
			}
		}
		data = data.Next
	}
}

func (l *LinkedList) DeletedWithValue(number int) {
	if l.Length == 0 {
		return
	}
	if l.Head.Data == number {
		l.Head = l.Head.Next
		l.Length--
		return
	}
	previousToDelete := l.Head
	for previousToDelete.Next.Data != number {
		if previousToDelete.Next.Next == nil {
			return
		}
		previousToDelete = previousToDelete.Next
	}
	previousToDelete.Next = previousToDelete.Next.Next
	l.Length--
}
